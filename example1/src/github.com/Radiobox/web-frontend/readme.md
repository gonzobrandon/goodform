## Radiobox

This is our primary website.  The database logic and API is all written in Go with
oauth2 authentication, and the frontend is written in pure angular.js (using
angular's templating system and everything).

# Cloning the repo

Go has a tool that is generally used to compile Go source, run tests, and so on.
The tool also handles dependencies automatically, so long as said dependencies are
hosted on github, google code, or one of a few other places.  This makes it very
easy to get a working development instance, but you have to clone the project a
little differently than you might be used to.

Follow the following steps to get everything set up:

1. If they're not already installed, install git, mercurial, and bazaar.  These
are the main tools that will be used for cloning dependencies.
1. Install the [Go tool](http://golang.org/doc/install).
2. Create a directory for all of your Go source code (e.g. ~/dev/golang).
3. Set your GOPATH environment variable to point to your go source directory, then
add $GOPATH/bin to your $PATH.
  1. `export GOPATH=~/dev/golang`
  2. `export PATH="$GOPATH/bin:$PATH"`
  3. You probably want to add the above two lines to your .profile, .bashrc, or
  .zshrc, or whatever else you use.
4. Run the go get command: `go get github.com/Radiobox/web-frontend`
5. If you want to use SSH instead of HTTP for pushing and pulling, change your
origin remote to use SSH:

```
$ cd $GOPATH/src/github.com/Radiobox/web-frontend
$ git remote rm origin
$ git remote add origin git@github.com:Radiobox/web-frontend
```

Now that all that is done, you should have the entire project cloned in
`$GOPATH/src/github.com/Radiobox/web-frontend`.

# Staying up to date

Running `git pull` will pull down updates to *our* project, but what
about the projects that our project depends on?  When we run the `go
get` command, initially, it clones many github projects that we're
importing, and if we start using newer features, you'll need to update
those packages.

Luckily, the `go get` command also has a feature to update all
dependencies of a package.  Just run the following:

1. Make sure that you have committed or stashed all of your local changes.
2. Change directory to the root location of this project.
3. Run `go get -u ./...`

Please note that running the above command will automatically check
out the master branch of this project, so don't be surprised when that
happens.  Once the command has finished, you can simply check out the
branch you were working on.

# Running a local version of the project

##### NOTE: All of the following steps should be run from within the `$GOPATH/src/github.com/Radiobox/web-frontend` directory, unless stated otherwise.

If you followed the above steps, you should already be most of the way there.
The next step is setting up a database.

First, install postgresql 9.3 and PostGIS.  Set up a database cluster.  If you
need help with this, use google, or ask one of the other developers directly
for help.

Next, run the database initialization from our Makefile:

```
make devdb-full
```

This will restore the full database structure, as well as creating some initial
data for things like oauth2 clients and such.

Next, you'll need the [heroku toolbelt](https://toolbelt.heroku.com/),
to pull down some environment variables from heroku.

If you want to skip the rest of these steps, simply run the following:

```
PORT=5000 make dev-run
```

The above command (run from inside the project directory) will run the
rest of the setup commands for y
ou and start the project.  The rest of
this guide is for running without make.

You need to set up your heroku remote before running any commands with
the heroku tool:

```
git remote add heroku git@heroku.com:radiobox-api-frontend-go.git
```

You will need to run the following in any shell that you want to run
the project in:


```
cd $GOPATH/src/github.com/Radiobox/web-frontend
export $(heroku config --shell)
```

If you're developing on this project exclusively, you may want to add
that command to your .profile, .bash_profile, and/or .zsh_profile.

At this point, everything should be set, so you'll want to run the project.
Assuming that your radiobox user can connect without a password on localhost (which
is usually the default in postgres), you should be able to simply compile and install
the project, then run it with a PORT variable to specify which port it should be
run on:

```
go install
PORT=5000 web-frontend
```

The project should now be running at localhost:5000.  To test this, create a
test user (you'll probably need one soon, anyway) by POSTing to /api/user:

```
curl -X POST -d "client_id=00000001&username=test&password=test&email=test@test.com" http://localhost:5000/api/user/
```

You should get back a response with the user's ID.

# Database commands from the Makefile

Most of the things that you'll need to do with the database dumps can be done using
make commands.  Here is a list of make targets that are currently supported, and
a description of what each one does:

```
build-dist
        Runs grunt and compiles the web app into the public/dist folder, which is served in production environment. 
        You may need to update node.js on your local environment. The two CSS/JS files are named for most recent git
        commit.

db-requirements
        Prints out the required postgres version and any other dependencies

devdb-requirements
        If the "radiobox" database does not exist, this creates a "radiobox" user
        and creates a "radiobox" database owned by that user.
        
devdb-dump
        Reads the "radiobox" database and writes schema and data to:
        /database/development_full.sql.c, 
        /database/development_schema.sql.c,
        /database/development_data.sql.c. 
        
devdb-full
        Creates the "radiobox" user and database (if necessary), then restores the
        full database dump.  This will restore tables, indexes, functions, data, etc.
        Since this is a full restore, you should have some basic starting data, e.g.
        an oauth2 client (client_id 00000001).

devdb-blank
        Same as devdb-full, but without the data.  You will need to create your
        own oauth2 clients and all that.

devdb-data
        The opposite of devdb-blank - this will restore data only, without any
        table structure.  It's useful if you've been changing around some database
        structure and want to make sure that the data will easily transfer.

devdb-wipe
        This will drop your "radiobox" database.  It will first print a warning message
        and give you three seconds to change your mind, though.

devdb-reset
        This will reset your data back to a fresh devdb-full run.  Essentially, this
        runs devdb-wipe followed by devdb-full.
```

# Making updates to the database dumps

The database restore make commands above are, behind the scenes, just restoring a
custom format dump from postgres.  The custom format is used because it's a compression
format written by the postgres developers specifically to compress SQL text data, which
basically means that it has extremely high compression efficiency when compressing
SQL commands.

If you ever need to update those dumps with new data or structure, make the changes
in your development instance, make sure that all of your data is data you want to
add to the dumps, and then run the following commands:

```
pg_dump --format c --username radiobox --dbname radiobox -f $GOPATH/src/github.com/Radiobox/web-frontend/database/development_full.sql.c
pg_dump --schema-only --format c --username radiobox --dbname radiobox -f $GOPATH/src/github.com/Radiobox/web-frontend/database/development_schema.sql.c
pg_dump --data-only --format c --username radiobox --dbname radiobox -f $GOPATH/src/github.com/Radiobox/web-frontend/database/development_data.sql.c
```

Then, add those files to your next git commit.

# Working within the repo

First, a quick summary of the major libraries we are using:

1. [goweb](http://github.com/stretchr/goweb) is the main web framework we use.
It's very lightweight and is very good at mapping controllers to URLs and HTTP
verbs, and it does pretty decent automatic conversion of Go's data types to various
response types.  I believe that this part of goweb could be improved, but I plan
to do that myself and send them a pull request.
2. [gorp](http://github.com/coopernurse/gorp) is the ORM (kind of) library that
we use.
3. [testify](http://github.com/stretchr/testify) is what we use for unit tests.
4. We also have a fairly heavily modified version of [osin](http://github.com/RangelReale/osin).


# Data Encapsulation and Accept Headers

Whenever you make a request for data from our API, our default is to
respond with raw JSON data.  However, you can request other formats
using either the except header (preferred) or file extensions.

## File Extensions

Even though file extensions are not our preferred method of requesting
a format, I'm mentioning them first, because they're much simpler.
Basically, all you need to do is end the URL in the file extension
for the MIME type that you want the response in, and the response will
be formatted by a matching codec.

For example, if you request /api/users/1.xml, you will receive the
user profile in XML format; if you request /api/users/1.csv, you will
receive the same data in CSV format.

## The Accept Header

We prefer that you use the HTTP Accept header, which allows for more
flexibility and gives you access to our encapsulation formats (see
below).  Simply put the MIME type(s) that you are capable of handling
in the Accept header of your request, and the response will be in your
preferred format.

Some examples of Accept headers for supported codecs:

```
Accept: application/json
Accept: text/csv
Accept: text/xml
Accept: application/xml
Accept: application/bson
Accept: application/x-msgpack
```

## Encapsulated Codecs

We also support data encapsulation if you request it in your Accept
header.  The format is
`application/vnd.radiobox.encapsulated+<base_type>`, where `<base_type>` is
the name of any of the other codecs that we support.  To put it
simply, the `application/vnd.radiobox.encapsulated` part tells our
servers that you want your data encapsulated, and the `+xml` or
`+json` (or whatever you want) part tells our servers how to format
the data once it has been encapsulated.  The default, if you don't
request a specific type after the `+`, is json.

Some examples:

```
application/vnd.radiobox.encapsulated+json
application/vnd.radiobox.encapsulated+xml
application/vnd.radiobox.encapsulated+bson
```

If you use an encapsulated format in your Accept header, the response
will be encapsulated according to the following format:

1. `meta`: This is data about the data. Meta data can return parameters sent with the request,
the numerical response code and the nominal error type. The error type will be returned in an
underscore typed variable name.
  - `code`: This is the numeric value fo the HTTP response code that was returned.
  - `input_params`: A JSON map of the parameters sent with the request (for POST, GET, PUT
type calls).
2. `notifications`: This is a map of user-friendly error, warning, and info messages to display
to the user.
3. `response`: The response in resource/object form whenever possible. For example, if a user
was requested via GET /users/johndoe, the top level object within the response field will be
`user`. (Not Yet Implemented)

#### Example 1: A simple user fetch

*GET /users/1*

```
{
    "meta": {
        "code": 200,
        "input_params": {}
    },
    "notifications": {
        "err": [],
        "warn": [],
        "info": []
    },
    "response": {
        "user_id": 1,
        "first_name": John,
        "last_name": Doe,
        "locale": USA,
        "username": "johndoe",
        "wall_count": 0
    }
}
```

#### Example 2: Fetching a nonexistant user

*GET /users/1?with=pic*

```
{
    "meta": {
        "code": 404,
        "input_params": {"with": "pic"}
    },
    "notifications": {
        "err": ["A user was not found with that username or ID"],
        "warn": [],
        "info": []
    },
    "response": {}
}
```

#### Example 3: A list of users

*GET /users*

```
{
    "meta": {
        "code": 200,
        "input_params": {}
    },
    "notifications": {
        "err": [],
        "warn": [],
        "info": []
    },
    "response": [
        {
            "display_name": "foo",
            "profile_link": "/users/3312",
            "account_link": "/user/3312"
        },
        {
            "display_name": "bar",
            "profile_link": "/users/28971",
            "account_link": "/user/28971"
        },
    ]
}
```

# Documentation

While our documentation is currently a little sparse (sorry!), you can
access it at any time by querying the endpoint you want more
information about using the OPTIONS http verb.  Here is an example
(formatted for easier reading):

```
curl -X OPTIONS http://localhost:5000/api/users
{
"GET /api/users":"Retrieve a list of users.",
"GET /api/users/id":"Retrieve the details about a user.",
"PATCH /api/users/id":"Update a user's details.  This request requires
  an Authorization header with a valid access token.  Allowed
  parameters: first_name, last_name, sex,
  birth_date.",
"description":"The users endpoint is for retrieving or updating public
  information about a user.  User creation is handled at the user
  endpoint, since the users endpoint is only supposed to deal with
  public details, and therefor can't set email address, username, or
  password details."
}
  
```

These responses will eventually also include links to other, related
endpoints, which may be included in responses.

# Deploying to Heroku

### Getting a github token

If you have not previously deployed to heroku on the computer that you're using,
you will need a github application auth token.  The deployment script will ask
for your token during the deployment process, so all you need to do is go to
github and generate it, then copy/paste.  Here's how to generate your token on
github:

1. Go to your account settings on github.com
2. Click on `Applications`
3. Next to `Personal Access Tokens`, click on `Create new token`
4. Choose a name for heroku deployment from your current computer and generate
your token.
5. Once you have pasted your token at the command line during deployment, it is
stored permanently, so you can safely forget that particular token.


### Adding a heroku remote

In order to deploy to heroku, you need to add our heroku git repo URL as a git
remote if you haven't already.  If you don't have a heroku remote, the make command will error.  To add
our heroku project as your heroku remote, run the following command from
anywhere in the project's source directories:

```
git remote add heroku git@heroku.com:radiobox-api-frontend-go.git
```

### Using make to target Heroku

We have a make target which will do some initial checks to make sure you're set
up to deploy to heroku, then run the deployment.  To use it, run the following
command from within `$GOPATH/src/github/Radiobox/web-frontend`:

```
make heroku-deploy
```

It should do a decent job of telling you about any extra steps you need to take,
but here is some more detailed information, in case the makefile is confusing.



### Check it out on Heroku:

Once the app has been deployed to remote, it will show up here:

#### http://radiobox-api-frontend-go.herokuapp.com

