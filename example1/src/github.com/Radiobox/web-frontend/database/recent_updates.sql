CREATE TABLE beta_signups (
  beta_signup_id bigserial primary key,
  artist_name varchar not null,
  email varchar not null
);

alter table media add column original_file_name varchar;
alter table media add column bucket_name varchar;
alter table media add column content_type varchar;
alter table media add column content_length bigint;
alter table media add column status varchar;
alter table media add column bitrate varchar not null default '';
alter table tracks add column encoding_media_id varchar;
create index tracks_encoding_media_id_idx on tracks(encoding_media_id);
alter table tracks add column encoding_status varchar;
alter table venues add column timezone varchar;
alter table media add column duration interval;

ALTER TABLE artists ALTER COLUMN contact_phone_country_code DROP NOT NULL;
ALTER TABLE artists ALTER COLUMN is_verified DROP NOT NULL;
ALTER TABLE artists ALTER COLUMN is_verified SET DEFAULT false;
ALTER TABLE artists ALTER COLUMN is_active DROP NOT NULL;
ALTER TABLE artists ALTER COLUMN is_active SET DEFAULT false;
ALTER TABLE artists ALTER COLUMN contact_phone SET DATA TYPE smallint USING (contact_phone::integer);
ALTER TABLE artists DROP COLUMN slug;
ALTER TABLE artists DROP COLUMN pic;
ALTER TABLE artists DROP COLUMN pic_id;
ALTER TABLE artists DROP COLUMN pic_big;
ALTER TABLE artists DROP COLUMN pic_big_id;
ALTER TABLE artists DROP COLUMN pic_small;
ALTER TABLE artists DROP COLUMN pic_small_id;

ALTER TABLE users DROP COLUMN pic;
ALTER TABLE users DROP COLUMN pic_id;
ALTER TABLE users DROP COLUMN pic_big;
ALTER TABLE users DROP COLUMN pic_big_id;
ALTER TABLE users DROP COLUMN pic_small;
ALTER TABLE users DROP COLUMN pic_small_id;
ALTER TABLE users ALTER COLUMN contact_phone_country_code DROP NOT NULL;
ALTER TABLE users ALTER COLUMN contact_phone SET DATA TYPE bigint USING (contact_phone::integer);

ALTER TABLE venues DROP COLUMN pic;
ALTER TABLE venues DROP COLUMN pic_id;
ALTER TABLE venues DROP COLUMN pic_big;
ALTER TABLE venues DROP COLUMN pic_big_id;
ALTER TABLE venues DROP COLUMN pic_small;
ALTER TABLE venues DROP COLUMN pic_small_id;
ALTER TABLE venues RENAME phone TO contact_phone;
ALTER TABLE venues ALTER COLUMN contact_phone SET DATA TYPE bigint USING (contact_phone::integer);
ALTER TABLE venues ADD COLUMN contact_phone_country_code smallint;
ALTER TABLE venues ADD COLUMN contact_phone_extension varchar;
ALTER TABLE venues DROP COLUMN postal_code;
ALTER TABLE venues DROP COLUMN state;

ALTER TABLE albums DROP COLUMN pic_big;
ALTER TABLE albums DROP COLUMN pic_big_id;
ALTER TABLE albums DROP COLUMN pic_small;
ALTER TABLE albums DROP COLUMN pic_small_id;

ALTER TABLE tracks DROP COLUMN pic_big;
ALTER TABLE tracks DROP COLUMN pic_big_id;
ALTER TABLE tracks DROP COLUMN pic_small;
ALTER TABLE tracks DROP COLUMN pic_small_id;

ALTER TABLE users ADD COLUMN current_location_updated date;

alter table album_track add column track_number smallint;
create unique index album_track_album_id_track_number_unique on album_track(album_id, track_number);

alter table events_live add column pic_square_id uuid references photos(photo_id);
alter table events_live add column pic_square varchar;
alter table events_live drop column pic_big;
alter table events_live drop column pic_big_id;
alter table events_live drop column pic_small;
alter table events_live drop column pic_small_id;
