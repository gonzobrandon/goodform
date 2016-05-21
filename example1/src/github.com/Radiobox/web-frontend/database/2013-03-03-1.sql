ALTER TABLE photos ADD COLUMN options json;

ALTER TABLE artists ADD COLUMN pic_cover_id uuid references photos(photo_id);

ALTER TABLE artists ADD COLUMN description text;
