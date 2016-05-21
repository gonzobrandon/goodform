package users

import (
	"bytes"
	"errors"
	htmltemplates "html/template"
	"log"
	"net/smtp"
	"os"
	"path"
	texttemplates "text/template"
	"time"

	"github.com/Radiobox/osin"
	"github.com/Radiobox/web-frontend/datastore"
	"github.com/Radiobox/web-frontend/models/auth"
	"github.com/Radiobox/web-frontend/models/logs"
	"github.com/Radiobox/web-frontend/settings"
	"github.com/coopernurse/gorp"
	"github.com/nelsam/gophermail"
	"github.com/stretchr/objx"
)

const (
	rbFromAddr  = "robot@theradiobox.com"
	evType      = "email_verification"
	pwResetType = "password_reset"
)

var (
	smtpHost = os.Getenv("SMTP_HOST")
	smtpUser = os.Getenv("SMTP_USER")
	smtpPass = os.Getenv("SMTP_PASSWORD")
	smtpAuth = smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	textVerificationPath        = path.Join(settings.ProjectPath, "email", "verification.txt")
	htmlVerificationPath        = path.Join(settings.ProjectPath, "email", "verification.html")
	textVerificationTemplate, _ = texttemplates.ParseFiles(textVerificationPath)
	htmlVerificationTemplate, _ = htmltemplates.ParseFiles(htmlVerificationPath)
)

type Token struct {
	Id   *string  `db:"token_id" request:"-"`
	User *Account `db:"user_id" request:"user"`
	Type *string  `request:"-"`
}

func (token *Token) UserProfile() *Profile {
	result, err := datastore.Get(new(Profile), token.User.Id)
	if err != nil {
		return nil
	}
	return result.(*Profile)
}

type EmailVerification struct {
	Token
}

func (ev *EmailVerification) PreInsert(exec gorp.SqlExecutor) error {
	if ev.Type == nil {
		ev.Type = new(string)
	}
	*ev.Type = evType
	return nil
}

func (ev *EmailVerification) PostGet(exec gorp.SqlExecutor) error {
	if ev.Type == nil {
		return errors.New("No valid token found")
	} else if *ev.Type != evType {
		return errors.New("Trying to load token of type " + *ev.Type + " as type " + evType)
	}
	return nil
}

func (ev *EmailVerification) SendVerification() error {
	if ev.Id == nil {
		return errors.New("No verification to send")
	}

	textBuffer := new(bytes.Buffer)
	if err := textVerificationTemplate.Execute(textBuffer, ev); err != nil {
		return err
	}
	textBody := textBuffer.Bytes()

	htmlBuffer := new(bytes.Buffer)
	if err := htmlVerificationTemplate.Execute(htmlBuffer, ev); err != nil {
		return err
	}
	htmlBody := htmlBuffer.Bytes()

	message := &gophermail.Message{
		Subject:  "Welcome to Radioox",
		Body:     string(textBody),
		HTMLBody: string(htmlBody),
	}
	if err := message.SetFrom(rbFromAddr); err != nil {
		log.Print("Err while setting from: ", err)
		return err
	}
	if err := message.AddTo(string(ev.User.Email)); err != nil {
		log.Print("Err while setting to: ", err)
		return err
	}
	if err := gophermail.SendMail(smtpHost+":587", smtpAuth, message); err != nil {
		log.Print("Err while sending: ", err)
		return err
	}
	go logs.WriteLog(objx.Map{
		"type":    "email_sent_0.1",
		"from":    rbFromAddr,
		"to":      ev.User.Email,
		"status":  "sent",
		"name":    message.Subject,
		"version": "0.1",
		"caller":  "EmailVerification.SendVerification",
	})
	return nil
}

func (ev *EmailVerification) Verify() error {
	user := ev.UserProfile()
	user.Active = new(bool)
	*user.Active = true
	user.Verified = new(bool)
	*user.Verified = true
	if _, err := datastore.Update(user); err != nil {
		return err
	}
	return nil
}

type PasswordReset struct {
	Token
}

func (reset *PasswordReset) PreInsert(exec gorp.SqlExecutor) error {
	if reset.Type == nil {
		reset.Type = new(string)
	}
	*reset.Type = pwResetType
	return nil
}

func (reset *PasswordReset) PostGet(exec gorp.SqlExecutor) error {
	if reset.Type == nil {
		return errors.New("No valid token found")
	} else if *reset.Type != pwResetType {
		return errors.New("Trying to load token of type " + *reset.Type + " as type " + pwResetType)
	}
	return nil
}

func (reset *PasswordReset) GetUser() *Account {
	result, err := datastore.Get(new(Account), reset.User.Id)
	if err != nil || result == nil {
		return nil
	}
	return result.(*Account)
}

func (reset *PasswordReset) Access() (*auth.AccessData, error) {
	result, err := datastore.Get(new(auth.AccessData), reset.Id)
	return result.(*auth.AccessData), err
}

func (reset *PasswordReset) CreateAccess(client *auth.Client) *auth.AccessData {
	return &auth.AccessData{
		Client: client,
		BasicAccessData: osin.BasicAccessData{
			AuthorizeData: nil,
			AccessData:    nil,
			AccessToken:   *reset.Id,
			RefreshToken:  "",
			ExpiresIn:     3600,
			Scope:         "password-reset",
			RedirectUri:   "",
			CreatedAt:     time.Now(),
		},
		UserId: reset.User.Id,
	}
}

func (reset *PasswordReset) SendResetEmail() error {
	if reset.Id == nil {
		return errors.New("No reset to send")
	}
	message := &gophermail.Message{
		Subject: "Radiobox Password Reset",
		Body:    "This is your password reset email.  Visit http://www.theradiobox.com/password-reset/#" + *reset.Id,
		HTMLBody: `<HTML><BODY><H1>This is your password reset email</H1><P><A HREF="http://www.theradiobox.com/password-reset/#` + *reset.Id +
			`">Click here</A> to reset your password.</P></BODY></HTML>`,
	}
	if err := message.SetFrom(rbFromAddr); err != nil {
		log.Print("Error while setting from: ", err)
		return err
	}
	if err := message.AddTo(string(reset.User.Email)); err != nil {
		log.Print("Error while setting to: ", err)
		return err
	}
	if err := gophermail.SendMail(smtpHost+":587", smtpAuth, message); err != nil {
		log.Print("Err while sending: ", err)
		return err
	}
	go logs.WriteLog(objx.Map{
		"type":    "email_sent_0.1",
		"from":    rbFromAddr,
		"to":      reset.User.Email,
		"status":  "sent",
		"name":    message.Subject,
		"version": "I'm not sure what to put here",
		"caller":  "PasswordReset.SendResetEmail",
	})
	go logs.WriteLog(objx.Map{
		"type":     "user_password_reset_0.1",
		"user_id":  reset.User.Id,
		"ip":       "We have no way of retrieving this at this time",
		"username": reset.User.Username,
	})
	return nil
}

type betaEmail struct {
	userEmail
}

func (email betaEmail) ValidateInput(input interface{}) error {
	emailPtr := &email
	if err := emailPtr.Receive(input); err != nil {
		return err
	}
	query := "SELECT count(*) FROM beta_signups WHERE email = $1"
	if count, err := datastore.SelectInt(query, input); err != nil {
		return errors.New("Internal error: could not check for duplicate: " + err.Error())
	} else if count > 0 {
		return errors.New("Duplicate email found")
	}
	return nil
}

func (email betaEmail) String() string {
	return string(email.userEmail)
}

type BetaSignup struct {
	Id         int64  `db:"beta_signup_id" request:"-"`
	ArtistName string `db:"artist_name"`
	Email      betaEmail
}

func (signup *BetaSignup) PostInsert(exec gorp.SqlExecutor) error {
	message := &gophermail.Message{
		Subject:  "Beta Signup Confirmation",
		Body:     "Thank you!\n\nWe have received your beta signup and will contact you when it's time to start setting up your beta account.",
		HTMLBody: `<HTML><BODY><H3>Thank you!</H3><P>We have received your beta signup and will contact you when it's time to start setting up your beta account.`,
	}
	if err := message.SetFrom(rbFromAddr); err != nil {
		log.Print("Error while setting from: ", err)
		return err
	}
	if err := message.AddTo(signup.Email.String()); err != nil {
		log.Print("Error while setting to: ", err)
		return err
	}
	if err := gophermail.SendMail(smtpHost+":587", smtpAuth, message); err != nil {
		log.Print("Err while sending: ", err)
		return err
	}
	go logs.WriteLog(objx.Map{
		"type":    "email_sent_0.1",
		"from":    rbFromAddr,
		"to":      signup.Email,
		"status":  "sent",
		"name":    message.Subject,
		"version": "I'm not sure what to put here",
		"caller":  "BetaSignup.PostInsert",
	})
	return nil
}
