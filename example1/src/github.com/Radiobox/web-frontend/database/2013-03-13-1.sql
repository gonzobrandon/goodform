ALTER TABLE event_live_provisioned_broadcasts RENAME COLUMN boadcast_url_1_port TO broadcast_url_1_port;
ALTER TABLE event_live_provisioned_broadcasts RENAME COLUMN boadcast_url_2_port TO broadcast_url_2_port;
ALTER TABLE event_live_provisioned_broadcasts RENAME COLUMN broascast_method TO broadcast_method;

CREATE TABLE logs (
       log_id bigserial primary key,
       log_type varchar not null,
       log_details json not null,
       created_at timestamp with time zone not null default now(),
       updated_at timestamp with time zone not null default now()
);
