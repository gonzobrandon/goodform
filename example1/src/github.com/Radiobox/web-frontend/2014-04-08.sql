alter table artists add column pic_cover_id uuid references photos(photo_id);
CREATE TABLE email_templates (
    email_template_id uuid PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    name varchar(40) NOT NULL,
    body_plaintext_link varchar(2048) NOT NULL,
    body_html_link varchar(2048) NOT NULL,
    subject varchar(40) NOT NULL,
    language varchar(3) NOT NULL,
    notes json,
    version decimal NOT NULL
);
create table email_verifications (
    email_verification_id uuid primary key default uuid_generate_v4(),
    user_id bigint references users(user_id),
    expires_at timestamptz not null default now() + '1 day'::interval
);
alter table email_verifications rename to tokens;
alter table tokens rename column email_verification_id to token_id;
alter index email_verifications_pkey rename to tokens_pkey;
alter table tokens rename constraint email_verifications_user_id_fkey to tokens_user_id_fkey;
alter table tokens add column type varchar;
create index tokens_type on tokens(type);

ALTER TABLE event_live_provisioned_broadcasts ADD COLUMN region varchar(40);
