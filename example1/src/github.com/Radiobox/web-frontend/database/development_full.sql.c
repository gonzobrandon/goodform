PGDMP     ,                    r           radiobox    9.3.4    9.3.4 �    �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                       false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                       false            �           1262    46179    radiobox    DATABASE     p   CREATE DATABASE radiobox WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'C' LC_CTYPE = 'en_US.UTF-8';
    DROP DATABASE radiobox;
             radiobox    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
             postgres    false            �           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                  postgres    false    6            �           0    0    public    ACL     �   REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;
                  postgres    false    6            �            3079    11754    plpgsql 	   EXTENSION     ?   CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;
    DROP EXTENSION plpgsql;
                  false            �           0    0    EXTENSION plpgsql    COMMENT     @   COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';
                       false    229            �            3079    46180    postgis 	   EXTENSION     ;   CREATE EXTENSION IF NOT EXISTS postgis WITH SCHEMA public;
    DROP EXTENSION postgis;
                  false    6            �           0    0    EXTENSION postgis    COMMENT     g   COMMENT ON EXTENSION postgis IS 'PostGIS geometry, geography, and raster spatial types and functions';
                       false    231            �            3079    47466 	   uuid-ossp 	   EXTENSION     ?   CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;
    DROP EXTENSION "uuid-ossp";
                  false    6            �           0    0    EXTENSION "uuid-ossp"    COMMENT     W   COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';
                       false    230                       1255    47477    artists_username_insert()    FUNCTION     �   CREATE FUNCTION artists_username_insert() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    BEGIN
        INSERT INTO usernames (target, username) VALUES (('{"artist_id":' || NEW.artist_id || '}')::json, NEW.username);
        RETURN NEW;
    END;
$$;
 0   DROP FUNCTION public.artists_username_insert();
       public       radiobox    false    6    229                       1255    47478    artists_username_update()    FUNCTION     �   CREATE FUNCTION artists_username_update() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    BEGIN
        UPDATE usernames SET username = NEW.username WHERE (usernames.target->>'artist_id')::int = NEW.artist_id;
        RETURN NEW;
    END;
$$;
 0   DROP FUNCTION public.artists_username_update();
       public       radiobox    false    6    229                       1255    47479    users_username_insert()    FUNCTION     �   CREATE FUNCTION users_username_insert() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    BEGIN
        INSERT INTO usernames (target, username) VALUES (('{"user_id":' || NEW.user_id || '}')::json, NEW.username);
        RETURN NEW;
    END;
$$;
 .   DROP FUNCTION public.users_username_insert();
       public       radiobox    false    229    6                       1255    47480    users_username_update()    FUNCTION     �   CREATE FUNCTION users_username_update() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    BEGIN
        UPDATE usernames SET username = NEW.username WHERE (usernames.target->>'user_id')::int = NEW.user_id;
        RETURN NEW;
    END;
$$;
 .   DROP FUNCTION public.users_username_update();
       public       radiobox    false    6    229            �            1259    47481    album_track    TABLE     -  CREATE TABLE album_track (
    album_track_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    album_id uuid NOT NULL,
    track_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    track_number smallint
);
    DROP TABLE public.album_track;
       public         radiobox    false    230    6    6            �            1259    47487    albums    TABLE     �  CREATE TABLE albums (
    album_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    artist_id bigint,
    downloads bigint DEFAULT (0)::bigint,
    title character varying(255),
    pic_square character varying(255),
    pic_square_id uuid,
    price_usd money DEFAULT (0)::numeric NOT NULL,
    price_usd_min money DEFAULT (0)::numeric NOT NULL,
    price_nyo boolean DEFAULT true NOT NULL,
    purchases bigint DEFAULT (0)::bigint,
    time_length bigint DEFAULT (0)::bigint,
    is_active boolean DEFAULT true NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
    DROP TABLE public.albums;
       public         radiobox    false    230    6    6            �            1259    47503    artist_checkins    TABLE     Y  CREATE TABLE artist_checkins (
    artist_checkin_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    checkin_via character varying(255),
    is_location_matched boolean,
    checkin_scan_code character varying(255),
    checkin_location_point geography(Point,4326),
    checkin_location_match_service character varying(255),
    checkin_location_match_name character varying(255),
    checkin_location_venue_match_id bigint,
    checkin_message character varying(255),
    pic_big character varying(255),
    pic_big_id uuid,
    pic_small character varying(255),
    pic_small_id uuid,
    pic_square character varying(255),
    pic_square_id uuid,
    checkin_timestamp timestamp with time zone,
    artist_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 #   DROP TABLE public.artist_checkins;
       public         radiobox    false    230    6    231    231    6    231    6    231    6    231    6    231    6    231    6    6    231    6    6            �            1259    47512    artist_notifications    TABLE     *  CREATE TABLE artist_notifications (
    artist_notification_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    artist_id bigint NOT NULL,
    notification json NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 (   DROP TABLE public.artist_notifications;
       public         radiobox    false    230    6    6            �            1259    47521    artist_user    TABLE     R  CREATE TABLE artist_user (
    artist_user_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    artist_id bigint NOT NULL,
    user_id bigint NOT NULL,
    is_admin boolean DEFAULT false NOT NULL,
    role json,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
    DROP TABLE public.artist_user;
       public         radiobox    false    230    6    6            �            1259    47531    artists    TABLE     T  CREATE TABLE artists (
    artist_id bigint NOT NULL,
    band_user_ids json,
    booking_user_id bigint,
    contact_email character varying(255) NOT NULL,
    contact_phone smallint,
    contact_phone_country_code smallint DEFAULT (1)::smallint,
    contact_phone_extension character varying(10),
    facebook_page_id character varying(255),
    facebook_page_token_encrypted character varying(255),
    hometown_address json,
    hometown_address_point geography(Point,4326),
    keywords json,
    manager_user_id bigint,
    pic_cover character varying(255),
    pic_square character varying(255),
    pic_square_id uuid,
    record_label_id character varying(255),
    subscriber_count bigint,
    timezone character varying(255),
    twitter_account_settings_id uuid,
    username character varying(255) NOT NULL,
    website character varying(255),
    is_active boolean DEFAULT false,
    is_verified boolean DEFAULT false,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    description text,
    pic_cover_id uuid
);
    DROP TABLE public.artists;
       public         radiobox    false    231    231    6    231    6    231    6    231    6    231    6    231    6    6    231    6    6            �            1259    47542    artists_artist_id_seq    SEQUENCE     w   CREATE SEQUENCE artists_artist_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 ,   DROP SEQUENCE public.artists_artist_id_seq;
       public       radiobox    false    6    188            �           0    0    artists_artist_id_seq    SEQUENCE OWNED BY     A   ALTER SEQUENCE artists_artist_id_seq OWNED BY artists.artist_id;
            public       radiobox    false    189            �            1259    47544    beta_signups    TABLE     �   CREATE TABLE beta_signups (
    beta_signup_id bigint NOT NULL,
    artist_name character varying NOT NULL,
    email character varying NOT NULL
);
     DROP TABLE public.beta_signups;
       public         radiobox    false    6            �            1259    47550    beta_signups_beta_signup_id_seq    SEQUENCE     �   CREATE SEQUENCE beta_signups_beta_signup_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 6   DROP SEQUENCE public.beta_signups_beta_signup_id_seq;
       public       radiobox    false    6    190            �           0    0    beta_signups_beta_signup_id_seq    SEQUENCE OWNED BY     U   ALTER SEQUENCE beta_signups_beta_signup_id_seq OWNED BY beta_signups.beta_signup_id;
            public       radiobox    false    191            �            1259    47552    email_templates    TABLE       CREATE TABLE email_templates (
    email_template_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    name character varying(40) NOT NULL,
    body_plaintext_link character varying(2048) NOT NULL,
    body_html_link character varying(2048) NOT NULL,
    subject character varying(40) NOT NULL,
    language character varying(3) NOT NULL,
    notes json,
    version numeric NOT NULL
);
 #   DROP TABLE public.email_templates;
       public         radiobox    false    230    6    6            �            1259    47559    event_admin_log    TABLE       CREATE TABLE event_admin_log (
    event_admin_log_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    event_live_id uuid NOT NULL,
    data json NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 #   DROP TABLE public.event_admin_log;
       public         radiobox    false    230    6    6            �            1259    47568    event_live_notifications    TABLE     S  CREATE TABLE event_live_notifications (
    event_live_notification_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    event_live_id uuid NOT NULL,
    notification json NOT NULL,
    title json NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);
 ,   DROP TABLE public.event_live_notifications;
       public         radiobox    false    230    6    6            �            1259    47577 !   event_live_provisioned_broadcasts    TABLE     �  CREATE TABLE event_live_provisioned_broadcasts (
    event_live_provisioned_broadcast_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    type character varying(255),
    is_video boolean DEFAULT false NOT NULL,
    broadcast_url_1_port smallint,
    broadcast_url_2_port smallint,
    broadcast_stream_name character varying(255),
    provider_stream_id character varying(255),
    broadcast_url_1 character varying(255),
    broadcast_url_2 character varying(255),
    broadcast_method character varying(255),
    broadcast_username character varying(255),
    broadcast_password character varying(255),
    encode_suggested_params json,
    encode_accepted_params json,
    client_hds character varying(255),
    client_hls character varying(255),
    client_hdflash1 character varying(255),
    client_shoutcast_url character varying(255),
    reserved_until timestamp with time zone,
    is_in_progress boolean DEFAULT false NOT NULL,
    is_concluded boolean DEFAULT false NOT NULL,
    is_available boolean DEFAULT true NOT NULL,
    listeners_max bigint DEFAULT (0)::bigint NOT NULL,
    listeners_now bigint DEFAULT (0)::bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 5   DROP TABLE public.event_live_provisioned_broadcasts;
       public         radiobox    false    230    6    6            �            1259    47592    event_live_routes    TABLE       CREATE TABLE event_live_routes (
    event_live_route_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    event_live_id uuid NOT NULL,
    event_live_provisioned_broadcast_id uuid NOT NULL,
    priority smallint DEFAULT (1000)::smallint NOT NULL,
    is_primary boolean DEFAULT false NOT NULL,
    listeners_max bigint DEFAULT (0)::bigint NOT NULL,
    listeners_now bigint DEFAULT (0)::bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 %   DROP TABLE public.event_live_routes;
       public         radiobox    false    230    6    6            �            1259    47602    event_live_segments    TABLE     �  CREATE TABLE event_live_segments (
    event_live_segment_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    end_time bigint NOT NULL,
    event_live_id uuid NOT NULL,
    media_id json NOT NULL,
    segment_name character varying(255) NOT NULL,
    start_time bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 '   DROP TABLE public.event_live_segments;
       public         radiobox    false    230    6    6            �            1259    47611    events_live    TABLE       CREATE TABLE events_live (
    event_live_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    actual_end timestamp with time zone,
    actual_start timestamp with time zone,
    archive_track_id uuid,
    archiving boolean,
    artist_id bigint,
    title character varying(255),
    title_from character varying(255),
    artist_collaborators_id json,
    is_in_progress boolean,
    is_concluded boolean,
    off_schedule_end_seconds bigint,
    off_schedule_start_seconds bigint,
    scheduled_end timestamp with time zone,
    scheduled_start timestamp with time zone NOT NULL,
    location_address json,
    location_address_point geography(Point,4326),
    venue_id bigint,
    listeners_max bigint,
    is_active boolean DEFAULT true NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    caption character varying,
    is_standby boolean,
    messages json,
    standby_message character varying,
    pic_square_id uuid,
    pic_square character varying
);
    DROP TABLE public.events_live;
       public         radiobox    false    230    6    6    231    231    6    231    6    231    6    231    6    231    6    231    6    6    231    6            �            1259    47621    logs    TABLE     �   CREATE TABLE logs (
    log_id bigint NOT NULL,
    log_type character varying NOT NULL,
    log_details json NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
    DROP TABLE public.logs;
       public         radiobox    false    6            �            1259    47629    logs_log_id_seq    SEQUENCE     q   CREATE SEQUENCE logs_log_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.logs_log_id_seq;
       public       radiobox    false    199    6            �           0    0    logs_log_id_seq    SEQUENCE OWNED BY     5   ALTER SEQUENCE logs_log_id_seq OWNED BY logs.log_id;
            public       radiobox    false    200            �            1259    47631    media    TABLE     �  CREATE TABLE media (
    media_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    file_store_id character varying(255),
    file_store_key character varying(255),
    is_preview boolean,
    media_format_id uuid,
    "media_MD5" character varying(255),
    secret_url character varying(255),
    title character varying(255),
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    original_file_name character varying,
    bucket_name character varying,
    content_type character varying,
    content_length bigint,
    status character varying,
    bitrate character varying DEFAULT ''::character varying NOT NULL,
    duration interval
);
    DROP TABLE public.media;
       public         radiobox    false    230    6    6            �            1259    47641    media_formats    TABLE     �  CREATE TABLE media_formats (
    media_format_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    bitrate smallint,
    channels numeric(8,2),
    format character varying(255),
    is_lossless boolean DEFAULT false NOT NULL,
    name character varying(255) NOT NULL,
    pic_id uuid,
    pic_url character varying(255),
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 !   DROP TABLE public.media_formats;
       public         radiobox    false    230    6    6            �            1259    47651 
   migrations    TABLE     g   CREATE TABLE migrations (
    migration character varying(255) NOT NULL,
    batch integer NOT NULL
);
    DROP TABLE public.migrations;
       public         radiobox    false    6            �            1259    47654    oauth_access_data    TABLE     �  CREATE TABLE oauth_access_data (
    client character varying(255),
    authorizedata character varying(255),
    accessdata character varying(255),
    accesstoken character varying(255) NOT NULL,
    refreshtoken character varying(255),
    expiresin integer,
    scope character varying(255),
    redirecturi character varying(255),
    createdat timestamp with time zone,
    userid character varying(255)
);
 %   DROP TABLE public.oauth_access_data;
       public         radiobox    false    6            �            1259    47660    oauth_auth_data    TABLE     <  CREATE TABLE oauth_auth_data (
    client character varying(255),
    code character varying(255) NOT NULL,
    expiresin integer,
    scope character varying(255),
    redirecturi character varying(255),
    state character varying(255),
    createdat timestamp with time zone,
    userid character varying(255)
);
 #   DROP TABLE public.oauth_auth_data;
       public         radiobox    false    6            �            1259    47666    oauth_clients    TABLE     �   CREATE TABLE oauth_clients (
    id character varying(255) NOT NULL,
    secret character varying(255),
    redirecturi character varying(255)
);
 !   DROP TABLE public.oauth_clients;
       public         radiobox    false    6            �            1259    47672    payment_services    TABLE     0  CREATE TABLE payment_services (
    payment_service_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    name character varying(255),
    script_location character varying(255),
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 $   DROP TABLE public.payment_services;
       public         radiobox    false    230    6    6            �            1259    47681    photos    TABLE     �  CREATE TABLE photos (
    photo_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    caption character varying(255),
    file_store_id character varying(255),
    file_store_key character varying(255),
    format character varying(255),
    height smallint,
    "media_MD5" character varying(255),
    media_mimetype character varying(255),
    owner_id character varying(255),
    owner_type character varying(255),
    secret_url character varying(255),
    url character varying(255),
    width smallint,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    options json
);
    DROP TABLE public.photos;
       public         radiobox    false    230    6    6            �            1259    47690    purchase_items    TABLE     P  CREATE TABLE purchase_items (
    purchase_item_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    purchase_id uuid NOT NULL,
    purchase_type character varying(255) NOT NULL,
    unit_price money NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 "   DROP TABLE public.purchase_items;
       public         radiobox    false    230    6    6            �            1259    47696    purchase_purchaseitem    TABLE     1  CREATE TABLE purchase_purchaseitem (
    purcahse_purchaseitem_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    purchase_id uuid NOT NULL,
    purchase_item_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 )   DROP TABLE public.purchase_purchaseitem;
       public         radiobox    false    230    6    6            �            1259    47702 	   purchases    TABLE     �  CREATE TABLE purchases (
    purchase_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    begin_cart timestamp with time zone,
    is_complete boolean,
    is_paid boolean,
    is_paypal boolean,
    is_bitcoin boolean,
    is_creditcard boolean,
    pay_grand_total money,
    pay_shipping money,
    pay_sub_total money,
    pay_discount money,
    pay_tax money,
    user_payment_id uuid,
    receipt_email character varying(255),
    ship_name_first character varying(255),
    ship_name_mi character varying(255),
    ship_name_last character varying(255),
    ship_address_1 character varying(255),
    ship_address_2 character varying(255),
    ship_address_verified boolean,
    ship_city character varying(255),
    ship_country character varying(3),
    ship_state character varying(30),
    ship_notes json,
    user_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
    DROP TABLE public.purchases;
       public         radiobox    false    230    6    6            �            1259    47711    slugs    TABLE     e   CREATE TABLE slugs (
    id bigint NOT NULL,
    target json,
    slug character varying NOT NULL
);
    DROP TABLE public.slugs;
       public         radiobox    false    6            �            1259    47717    slugs_id_seq    SEQUENCE     n   CREATE SEQUENCE slugs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.slugs_id_seq;
       public       radiobox    false    212    6            �           0    0    slugs_id_seq    SEQUENCE OWNED BY     /   ALTER SEQUENCE slugs_id_seq OWNED BY slugs.id;
            public       radiobox    false    213            �            1259    47719    tokens    TABLE     �   CREATE TABLE tokens (
    token_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    user_id bigint,
    expires_at timestamp with time zone DEFAULT (now() + '1 day'::interval) NOT NULL,
    type character varying
);
    DROP TABLE public.tokens;
       public         radiobox    false    230    6    6            �            1259    47727    tracks    TABLE     �  CREATE TABLE tracks (
    track_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    artist_id bigint,
    checksum_md5 character varying(255),
    downloads bigint DEFAULT (0)::bigint,
    format_keyval character varying(255),
    pic_square character varying(255),
    pic_square_id uuid,
    media_id uuid NOT NULL,
    must_bundle_album_id uuid,
    title character varying(255) NOT NULL,
    preview_media_id uuid NOT NULL,
    preview_only boolean,
    price_usd money DEFAULT (0)::numeric NOT NULL,
    price_usd_min money DEFAULT (0)::numeric NOT NULL,
    price_nyo boolean DEFAULT true NOT NULL,
    purchases bigint DEFAULT (0)::bigint,
    time_length bigint DEFAULT (0)::bigint,
    is_active boolean DEFAULT true NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    encoding_media_id character varying,
    encoding_status character varying
);
    DROP TABLE public.tracks;
       public         radiobox    false    230    6    6            �            1259    47743    twitter_account_settings    TABLE     �  CREATE TABLE twitter_account_settings (
    twitter_account_settings_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    handle character varying(255) NOT NULL,
    token1_encrypted character varying(255) NOT NULL,
    token2_encrypted character varying(255) NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 ,   DROP TABLE public.twitter_account_settings;
       public         radiobox    false    230    6    6            �            1259    47752    user_checkins    TABLE     S  CREATE TABLE user_checkins (
    user_checkin_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    checkin_via character varying(255),
    is_location_matched boolean,
    checkin_scan_code character varying(255),
    checkin_location_point geography(Point,4326),
    checkin_location_match_service character varying(255),
    checkin_location_match_name character varying(255),
    checkin_location_venue_match_id bigint,
    checkin_message character varying(255),
    pic_big character varying(255),
    pic_big_id uuid,
    pic_small character varying(255),
    pic_small_id uuid,
    pic_square character varying(255),
    pic_square_id uuid,
    checkin_timestamp timestamp with time zone,
    user_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 !   DROP TABLE public.user_checkins;
       public         radiobox    false    230    6    231    231    6    231    6    231    6    231    6    231    6    231    6    6    231    6    6            �            1259    47761    user_listens    TABLE     �  CREATE TABLE user_listens (
    user_listen_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    is_concluded boolean,
    is_live boolean DEFAULT false,
    play_max_position integer,
    play_cursor_position integer,
    user_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    asset json
);
     DROP TABLE public.user_listens;
       public         radiobox    false    230    6    6            �            1259    47771    user_notifications    TABLE     U  CREATE TABLE user_notifications (
    user_notification_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    link json,
    notification json,
    title character varying(255) NOT NULL,
    user_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 &   DROP TABLE public.user_notifications;
       public         radiobox    false    230    6    6            �            1259    47780    user_payments    TABLE     �  CREATE TABLE user_payments (
    user_payment_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    payment_service_id uuid NOT NULL,
    token character varying(255),
    token_salt character varying(255),
    payment_identifier_safe_string character varying(255),
    payment_identifier_name character varying(255),
    payment_expiration timestamp with time zone,
    bill_name_first character varying(255),
    bill_name_mi character varying(255),
    bill_name_last character varying(255),
    bill_address_1 character varying(255),
    bill_address_2 character varying(255),
    bill_address_verified boolean,
    bill_city character varying(255),
    bill_country character varying(3),
    bill_state character varying(30),
    bill_notes json,
    currency character varying(3) NOT NULL,
    user_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 !   DROP TABLE public.user_payments;
       public         radiobox    false    230    6    6            �            1259    47789 	   usernames    TABLE     m   CREATE TABLE usernames (
    id bigint NOT NULL,
    target json,
    username character varying NOT NULL
);
    DROP TABLE public.usernames;
       public         radiobox    false    6            �            1259    47795    usernames_id_seq    SEQUENCE     r   CREATE SEQUENCE usernames_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.usernames_id_seq;
       public       radiobox    false    221    6            �           0    0    usernames_id_seq    SEQUENCE OWNED BY     7   ALTER SEQUENCE usernames_id_seq OWNED BY usernames.id;
            public       radiobox    false    222            �            1259    47797    users    TABLE     �  CREATE TABLE users (
    user_id bigint NOT NULL,
    birthday_date date,
    current_address json,
    current_address_point geography(Point,4326),
    current_listen_media_id uuid,
    current_location json,
    current_location_point geography(Point,4326),
    contact_phone bigint,
    contact_phone_country_code smallint DEFAULT (1)::smallint,
    contact_phone_extension character varying(10),
    email_proxy character varying(255),
    email character varying(255) NOT NULL,
    email_pending_update character varying(255),
    facebook_token_encrypted character varying(1000),
    facebook_user character varying(100),
    first_name character varying(255),
    last_name character varying(255),
    interests character varying(255),
    keywords json,
    locale character varying(255),
    pass_hash character varying(64) NOT NULL,
    pic_cover character varying(255),
    pic_cover_id uuid,
    pic_square character varying(255),
    pic_square_id uuid,
    political character varying(255),
    profile_blurb character varying(255),
    user_payment_id uuid,
    relationship_status character varying(255),
    sex character varying(255),
    rb_subscriber_count bigint,
    facebook_subscriber_count bigint,
    facebook_last_data_pull timestamp with time zone,
    twitter_subscriber_count bigint,
    twitter_last_data_pull timestamp with time zone,
    twitter_account_settings_id uuid,
    linkedin_token_encrypted character varying(255),
    linkedin_user character varying(255),
    linkedin_last_data_pull timestamp with time zone,
    timezone character varying(255),
    username character varying(255) NOT NULL,
    wall_count character varying(255),
    work character varying(255),
    is_active boolean DEFAULT true NOT NULL,
    is_verified boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    current_location_updated date
);
    DROP TABLE public.users;
       public         radiobox    false    6    231    231    6    231    6    231    6    231    6    231    6    231    6    6    231    6    231    231    6    231    6    231    6    231    6    231    6    231    6    6    231    6            �            1259    47808    users_user_id_seq    SEQUENCE     s   CREATE SEQUENCE users_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.users_user_id_seq;
       public       radiobox    false    223    6            �           0    0    users_user_id_seq    SEQUENCE OWNED BY     9   ALTER SEQUENCE users_user_id_seq OWNED BY users.user_id;
            public       radiobox    false    224            �            1259    47810    venues    TABLE     P  CREATE TABLE venues (
    venue_id bigint NOT NULL,
    address json,
    email character varying(255) NOT NULL,
    facebook_token character varying(1000),
    facebook_user character varying(100),
    address_point geography(Point,4326),
    name character varying(255) NOT NULL,
    contact_phone bigint,
    pic_cover character varying(255),
    pic_cover_id uuid,
    pic_square character varying(255),
    pic_square_id uuid,
    url character varying(255),
    venue_blurb character varying(5000),
    is_active boolean DEFAULT true NOT NULL,
    is_verified boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    timezone character varying,
    contact_phone_country_code smallint,
    contact_phone_extension character varying
);
    DROP TABLE public.venues;
       public         radiobox    false    231    231    6    231    6    231    6    231    6    231    6    231    6    6    231    6    6            �            1259    47820    venues_venue_id_seq    SEQUENCE     u   CREATE SEQUENCE venues_venue_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 *   DROP SEQUENCE public.venues_venue_id_seq;
       public       radiobox    false    225    6            �           0    0    venues_venue_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE venues_venue_id_seq OWNED BY venues.venue_id;
            public       radiobox    false    226            �            1259    47822    voucher_purchase    TABLE     !  CREATE TABLE voucher_purchase (
    voucher_purchase_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    voucher_id uuid NOT NULL,
    purchase_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 $   DROP TABLE public.voucher_purchase;
       public         radiobox    false    230    6    6            �            1259    47828    vouchers    TABLE     �  CREATE TABLE vouchers (
    voucher_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    amount json,
    issuer_artist_id bigint,
    is_closed_out boolean DEFAULT false NOT NULL,
    expiration timestamp with time zone NOT NULL,
    bill_artist boolean DEFAULT false NOT NULL,
    bill_radiobox boolean DEFAULT true NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
    DROP TABLE public.vouchers;
       public         radiobox    false    230    6    6            �           2604    47840 	   artist_id    DEFAULT     h   ALTER TABLE ONLY artists ALTER COLUMN artist_id SET DEFAULT nextval('artists_artist_id_seq'::regclass);
 @   ALTER TABLE public.artists ALTER COLUMN artist_id DROP DEFAULT;
       public       radiobox    false    189    188            �           2604    47841    beta_signup_id    DEFAULT     |   ALTER TABLE ONLY beta_signups ALTER COLUMN beta_signup_id SET DEFAULT nextval('beta_signups_beta_signup_id_seq'::regclass);
 J   ALTER TABLE public.beta_signups ALTER COLUMN beta_signup_id DROP DEFAULT;
       public       radiobox    false    191    190                       2604    47842    log_id    DEFAULT     \   ALTER TABLE ONLY logs ALTER COLUMN log_id SET DEFAULT nextval('logs_log_id_seq'::regclass);
 :   ALTER TABLE public.logs ALTER COLUMN log_id DROP DEFAULT;
       public       radiobox    false    200    199            &           2604    47843    id    DEFAULT     V   ALTER TABLE ONLY slugs ALTER COLUMN id SET DEFAULT nextval('slugs_id_seq'::regclass);
 7   ALTER TABLE public.slugs ALTER COLUMN id DROP DEFAULT;
       public       radiobox    false    213    212            C           2604    47844    id    DEFAULT     ^   ALTER TABLE ONLY usernames ALTER COLUMN id SET DEFAULT nextval('usernames_id_seq'::regclass);
 ;   ALTER TABLE public.usernames ALTER COLUMN id DROP DEFAULT;
       public       radiobox    false    222    221            I           2604    47845    user_id    DEFAULT     `   ALTER TABLE ONLY users ALTER COLUMN user_id SET DEFAULT nextval('users_user_id_seq'::regclass);
 <   ALTER TABLE public.users ALTER COLUMN user_id DROP DEFAULT;
       public       radiobox    false    224    223            N           2604    47846    venue_id    DEFAULT     d   ALTER TABLE ONLY venues ALTER COLUMN venue_id SET DEFAULT nextval('venues_venue_id_seq'::regclass);
 >   ALTER TABLE public.venues ALTER COLUMN venue_id DROP DEFAULT;
       public       radiobox    false    226    225            e          0    47481    album_track 
   TABLE DATA               h   COPY album_track (album_track_id, album_id, track_id, created_at, updated_at, track_number) FROM stdin;
    public       radiobox    false    183   <t      f          0    47487    albums 
   TABLE DATA               �   COPY albums (album_id, artist_id, downloads, title, pic_square, pic_square_id, price_usd, price_usd_min, price_nyo, purchases, time_length, is_active, created_at, updated_at) FROM stdin;
    public       radiobox    false    184   Yt      g          0    47503    artist_checkins 
   TABLE DATA               p  COPY artist_checkins (artist_checkin_id, checkin_via, is_location_matched, checkin_scan_code, checkin_location_point, checkin_location_match_service, checkin_location_match_name, checkin_location_venue_match_id, checkin_message, pic_big, pic_big_id, pic_small, pic_small_id, pic_square, pic_square_id, checkin_timestamp, artist_id, created_at, updated_at) FROM stdin;
    public       radiobox    false    185   vt      h          0    47512    artist_notifications 
   TABLE DATA               p   COPY artist_notifications (artist_notification_id, artist_id, notification, created_at, updated_at) FROM stdin;
    public       radiobox    false    186   �t      i          0    47521    artist_user 
   TABLE DATA               j   COPY artist_user (artist_user_id, artist_id, user_id, is_admin, role, created_at, updated_at) FROM stdin;
    public       radiobox    false    187   �t      j          0    47531    artists 
   TABLE DATA               �  COPY artists (artist_id, band_user_ids, booking_user_id, contact_email, contact_phone, contact_phone_country_code, contact_phone_extension, facebook_page_id, facebook_page_token_encrypted, hometown_address, hometown_address_point, keywords, manager_user_id, pic_cover, pic_square, pic_square_id, record_label_id, subscriber_count, timezone, twitter_account_settings_id, username, website, is_active, is_verified, created_at, updated_at, description, pic_cover_id) FROM stdin;
    public       radiobox    false    188   �t      �           0    0    artists_artist_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('artists_artist_id_seq', 1, false);
            public       radiobox    false    189            l          0    47544    beta_signups 
   TABLE DATA               C   COPY beta_signups (beta_signup_id, artist_name, email) FROM stdin;
    public       radiobox    false    190   �t      �           0    0    beta_signups_beta_signup_id_seq    SEQUENCE SET     G   SELECT pg_catalog.setval('beta_signups_beta_signup_id_seq', 1, false);
            public       radiobox    false    191            n          0    47552    email_templates 
   TABLE DATA               �   COPY email_templates (email_template_id, name, body_plaintext_link, body_html_link, subject, language, notes, version) FROM stdin;
    public       radiobox    false    192   u      o          0    47559    event_admin_log 
   TABLE DATA               c   COPY event_admin_log (event_admin_log_id, event_live_id, data, created_at, updated_at) FROM stdin;
    public       radiobox    false    193   $u      p          0    47568    event_live_notifications 
   TABLE DATA               �   COPY event_live_notifications (event_live_notification_id, event_live_id, notification, title, created_at, updated_at) FROM stdin;
    public       radiobox    false    194   Au      q          0    47577 !   event_live_provisioned_broadcasts 
   TABLE DATA               �  COPY event_live_provisioned_broadcasts (event_live_provisioned_broadcast_id, type, is_video, broadcast_url_1_port, broadcast_url_2_port, broadcast_stream_name, provider_stream_id, broadcast_url_1, broadcast_url_2, broadcast_method, broadcast_username, broadcast_password, encode_suggested_params, encode_accepted_params, client_hds, client_hls, client_hdflash1, client_shoutcast_url, reserved_until, is_in_progress, is_concluded, is_available, listeners_max, listeners_now, created_at, updated_at) FROM stdin;
    public       radiobox    false    195   ^u      r          0    47592    event_live_routes 
   TABLE DATA               �   COPY event_live_routes (event_live_route_id, event_live_id, event_live_provisioned_broadcast_id, priority, is_primary, listeners_max, listeners_now, created_at, updated_at) FROM stdin;
    public       radiobox    false    196   {u      s          0    47602    event_live_segments 
   TABLE DATA               �   COPY event_live_segments (event_live_segment_id, end_time, event_live_id, media_id, segment_name, start_time, created_at, updated_at) FROM stdin;
    public       radiobox    false    197   �u      t          0    47611    events_live 
   TABLE DATA               �  COPY events_live (event_live_id, actual_end, actual_start, archive_track_id, archiving, artist_id, title, title_from, artist_collaborators_id, is_in_progress, is_concluded, off_schedule_end_seconds, off_schedule_start_seconds, scheduled_end, scheduled_start, location_address, location_address_point, venue_id, listeners_max, is_active, created_at, updated_at, caption, is_standby, messages, standby_message, pic_square_id, pic_square) FROM stdin;
    public       radiobox    false    198   �u      u          0    47621    logs 
   TABLE DATA               N   COPY logs (log_id, log_type, log_details, created_at, updated_at) FROM stdin;
    public       radiobox    false    199   �u      �           0    0    logs_log_id_seq    SEQUENCE SET     6   SELECT pg_catalog.setval('logs_log_id_seq', 1, true);
            public       radiobox    false    200            w          0    47631    media 
   TABLE DATA               �   COPY media (media_id, file_store_id, file_store_key, is_preview, media_format_id, "media_MD5", secret_url, title, created_at, updated_at, original_file_name, bucket_name, content_type, content_length, status, bitrate, duration) FROM stdin;
    public       radiobox    false    201   �v      x          0    47641    media_formats 
   TABLE DATA               �   COPY media_formats (media_format_id, bitrate, channels, format, is_lossless, name, pic_id, pic_url, created_at, updated_at) FROM stdin;
    public       radiobox    false    202   �v      y          0    47651 
   migrations 
   TABLE DATA               /   COPY migrations (migration, batch) FROM stdin;
    public       radiobox    false    203   �v      z          0    47654    oauth_access_data 
   TABLE DATA               �   COPY oauth_access_data (client, authorizedata, accessdata, accesstoken, refreshtoken, expiresin, scope, redirecturi, createdat, userid) FROM stdin;
    public       radiobox    false    204   �v      {          0    47660    oauth_auth_data 
   TABLE DATA               i   COPY oauth_auth_data (client, code, expiresin, scope, redirecturi, state, createdat, userid) FROM stdin;
    public       radiobox    false    205   w      |          0    47666    oauth_clients 
   TABLE DATA               9   COPY oauth_clients (id, secret, redirecturi) FROM stdin;
    public       radiobox    false    206   !w      }          0    47672    payment_services 
   TABLE DATA               f   COPY payment_services (payment_service_id, name, script_location, created_at, updated_at) FROM stdin;
    public       radiobox    false    207   Jw      ~          0    47681    photos 
   TABLE DATA               �   COPY photos (photo_id, caption, file_store_id, file_store_key, format, height, "media_MD5", media_mimetype, owner_id, owner_type, secret_url, url, width, created_at, updated_at, options) FROM stdin;
    public       radiobox    false    208   gw                0    47690    purchase_items 
   TABLE DATA               s   COPY purchase_items (purchase_item_id, purchase_id, purchase_type, unit_price, created_at, updated_at) FROM stdin;
    public       radiobox    false    209   �w      �          0    47696    purchase_purchaseitem 
   TABLE DATA               y   COPY purchase_purchaseitem (purcahse_purchaseitem_id, purchase_id, purchase_item_id, created_at, updated_at) FROM stdin;
    public       radiobox    false    210   �w      �          0    47702 	   purchases 
   TABLE DATA               �  COPY purchases (purchase_id, begin_cart, is_complete, is_paid, is_paypal, is_bitcoin, is_creditcard, pay_grand_total, pay_shipping, pay_sub_total, pay_discount, pay_tax, user_payment_id, receipt_email, ship_name_first, ship_name_mi, ship_name_last, ship_address_1, ship_address_2, ship_address_verified, ship_city, ship_country, ship_state, ship_notes, user_id, created_at, updated_at) FROM stdin;
    public       radiobox    false    211   �w      �          0    47711    slugs 
   TABLE DATA               *   COPY slugs (id, target, slug) FROM stdin;
    public       radiobox    false    212   �w      �           0    0    slugs_id_seq    SEQUENCE SET     3   SELECT pg_catalog.setval('slugs_id_seq', 1, true);
            public       radiobox    false    213            �          0    46448    spatial_ref_sys 
   TABLE DATA               Q   COPY spatial_ref_sys (srid, auth_name, auth_srid, srtext, proj4text) FROM stdin;
    public       postgres    false    171   x      �          0    47719    tokens 
   TABLE DATA               >   COPY tokens (token_id, user_id, expires_at, type) FROM stdin;
    public       radiobox    false    214   )x      �          0    47727    tracks 
   TABLE DATA               <  COPY tracks (track_id, artist_id, checksum_md5, downloads, format_keyval, pic_square, pic_square_id, media_id, must_bundle_album_id, title, preview_media_id, preview_only, price_usd, price_usd_min, price_nyo, purchases, time_length, is_active, created_at, updated_at, encoding_media_id, encoding_status) FROM stdin;
    public       radiobox    false    215   �x      �          0    47743    twitter_account_settings 
   TABLE DATA               �   COPY twitter_account_settings (twitter_account_settings_id, handle, token1_encrypted, token2_encrypted, created_at, updated_at) FROM stdin;
    public       radiobox    false    216   �x      �          0    47752    user_checkins 
   TABLE DATA               j  COPY user_checkins (user_checkin_id, checkin_via, is_location_matched, checkin_scan_code, checkin_location_point, checkin_location_match_service, checkin_location_match_name, checkin_location_venue_match_id, checkin_message, pic_big, pic_big_id, pic_small, pic_small_id, pic_square, pic_square_id, checkin_timestamp, user_id, created_at, updated_at) FROM stdin;
    public       radiobox    false    217   �x      �          0    47761    user_listens 
   TABLE DATA               �   COPY user_listens (user_listen_id, is_concluded, is_live, play_max_position, play_cursor_position, user_id, created_at, updated_at, asset) FROM stdin;
    public       radiobox    false    218   �x      �          0    47771    user_notifications 
   TABLE DATA               w   COPY user_notifications (user_notification_id, link, notification, title, user_id, created_at, updated_at) FROM stdin;
    public       radiobox    false    219   	y      �          0    47780    user_payments 
   TABLE DATA               g  COPY user_payments (user_payment_id, payment_service_id, token, token_salt, payment_identifier_safe_string, payment_identifier_name, payment_expiration, bill_name_first, bill_name_mi, bill_name_last, bill_address_1, bill_address_2, bill_address_verified, bill_city, bill_country, bill_state, bill_notes, currency, user_id, created_at, updated_at) FROM stdin;
    public       radiobox    false    220   &y      �          0    47789 	   usernames 
   TABLE DATA               2   COPY usernames (id, target, username) FROM stdin;
    public       radiobox    false    221   Cy      �           0    0    usernames_id_seq    SEQUENCE SET     7   SELECT pg_catalog.setval('usernames_id_seq', 1, true);
            public       radiobox    false    222            �          0    47797    users 
   TABLE DATA               4  COPY users (user_id, birthday_date, current_address, current_address_point, current_listen_media_id, current_location, current_location_point, contact_phone, contact_phone_country_code, contact_phone_extension, email_proxy, email, email_pending_update, facebook_token_encrypted, facebook_user, first_name, last_name, interests, keywords, locale, pass_hash, pic_cover, pic_cover_id, pic_square, pic_square_id, political, profile_blurb, user_payment_id, relationship_status, sex, rb_subscriber_count, facebook_subscriber_count, facebook_last_data_pull, twitter_subscriber_count, twitter_last_data_pull, twitter_account_settings_id, linkedin_token_encrypted, linkedin_user, linkedin_last_data_pull, timezone, username, wall_count, work, is_active, is_verified, created_at, updated_at, current_location_updated) FROM stdin;
    public       radiobox    false    223   ty      �           0    0    users_user_id_seq    SEQUENCE SET     8   SELECT pg_catalog.setval('users_user_id_seq', 1, true);
            public       radiobox    false    224            �          0    47810    venues 
   TABLE DATA               +  COPY venues (venue_id, address, email, facebook_token, facebook_user, address_point, name, contact_phone, pic_cover, pic_cover_id, pic_square, pic_square_id, url, venue_blurb, is_active, is_verified, created_at, updated_at, timezone, contact_phone_country_code, contact_phone_extension) FROM stdin;
    public       radiobox    false    225   z      �           0    0    venues_venue_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('venues_venue_id_seq', 1, false);
            public       radiobox    false    226            �          0    47822    voucher_purchase 
   TABLE DATA               i   COPY voucher_purchase (voucher_purchase_id, voucher_id, purchase_id, created_at, updated_at) FROM stdin;
    public       radiobox    false    227   ,z      �          0    47828    vouchers 
   TABLE DATA               �   COPY vouchers (voucher_id, amount, issuer_artist_id, is_closed_out, expiration, bill_artist, bill_radiobox, created_at, updated_at) FROM stdin;
    public       radiobox    false    228   Iz      Z           2606    47848    album_track_pkey 
   CONSTRAINT     _   ALTER TABLE ONLY album_track
    ADD CONSTRAINT album_track_pkey PRIMARY KEY (album_track_id);
 F   ALTER TABLE ONLY public.album_track DROP CONSTRAINT album_track_pkey;
       public         radiobox    false    183    183            \           2606    47850    albums_pkey 
   CONSTRAINT     O   ALTER TABLE ONLY albums
    ADD CONSTRAINT albums_pkey PRIMARY KEY (album_id);
 <   ALTER TABLE ONLY public.albums DROP CONSTRAINT albums_pkey;
       public         radiobox    false    184    184            ^           2606    47852    artist_checkins_pkey 
   CONSTRAINT     j   ALTER TABLE ONLY artist_checkins
    ADD CONSTRAINT artist_checkins_pkey PRIMARY KEY (artist_checkin_id);
 N   ALTER TABLE ONLY public.artist_checkins DROP CONSTRAINT artist_checkins_pkey;
       public         radiobox    false    185    185            `           2606    47854    artist_notifications_pkey 
   CONSTRAINT     y   ALTER TABLE ONLY artist_notifications
    ADD CONSTRAINT artist_notifications_pkey PRIMARY KEY (artist_notification_id);
 X   ALTER TABLE ONLY public.artist_notifications DROP CONSTRAINT artist_notifications_pkey;
       public         radiobox    false    186    186            b           2606    47856    artist_user_pkey 
   CONSTRAINT     _   ALTER TABLE ONLY artist_user
    ADD CONSTRAINT artist_user_pkey PRIMARY KEY (artist_user_id);
 F   ALTER TABLE ONLY public.artist_user DROP CONSTRAINT artist_user_pkey;
       public         radiobox    false    187    187            d           2606    47858    artists_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY artists
    ADD CONSTRAINT artists_pkey PRIMARY KEY (artist_id);
 >   ALTER TABLE ONLY public.artists DROP CONSTRAINT artists_pkey;
       public         radiobox    false    188    188            f           2606    47860    beta_signups_pkey 
   CONSTRAINT     a   ALTER TABLE ONLY beta_signups
    ADD CONSTRAINT beta_signups_pkey PRIMARY KEY (beta_signup_id);
 H   ALTER TABLE ONLY public.beta_signups DROP CONSTRAINT beta_signups_pkey;
       public         radiobox    false    190    190            h           2606    47862    email_templates_pkey 
   CONSTRAINT     j   ALTER TABLE ONLY email_templates
    ADD CONSTRAINT email_templates_pkey PRIMARY KEY (email_template_id);
 N   ALTER TABLE ONLY public.email_templates DROP CONSTRAINT email_templates_pkey;
       public         radiobox    false    192    192            j           2606    47864    event_admin_log_pkey 
   CONSTRAINT     k   ALTER TABLE ONLY event_admin_log
    ADD CONSTRAINT event_admin_log_pkey PRIMARY KEY (event_admin_log_id);
 N   ALTER TABLE ONLY public.event_admin_log DROP CONSTRAINT event_admin_log_pkey;
       public         radiobox    false    193    193            l           2606    47866    event_live_notifications_pkey 
   CONSTRAINT     �   ALTER TABLE ONLY event_live_notifications
    ADD CONSTRAINT event_live_notifications_pkey PRIMARY KEY (event_live_notification_id);
 `   ALTER TABLE ONLY public.event_live_notifications DROP CONSTRAINT event_live_notifications_pkey;
       public         radiobox    false    194    194            n           2606    47868 >   event_live_provisioned_broadcasts_broadcast_stream_name_unique 
   CONSTRAINT     �   ALTER TABLE ONLY event_live_provisioned_broadcasts
    ADD CONSTRAINT event_live_provisioned_broadcasts_broadcast_stream_name_unique UNIQUE (broadcast_stream_name);
 �   ALTER TABLE ONLY public.event_live_provisioned_broadcasts DROP CONSTRAINT event_live_provisioned_broadcasts_broadcast_stream_name_unique;
       public         radiobox    false    195    195            p           2606    47870 &   event_live_provisioned_broadcasts_pkey 
   CONSTRAINT     �   ALTER TABLE ONLY event_live_provisioned_broadcasts
    ADD CONSTRAINT event_live_provisioned_broadcasts_pkey PRIMARY KEY (event_live_provisioned_broadcast_id);
 r   ALTER TABLE ONLY public.event_live_provisioned_broadcasts DROP CONSTRAINT event_live_provisioned_broadcasts_pkey;
       public         radiobox    false    195    195            r           2606    47872 ;   event_live_provisioned_broadcasts_provider_stream_id_unique 
   CONSTRAINT     �   ALTER TABLE ONLY event_live_provisioned_broadcasts
    ADD CONSTRAINT event_live_provisioned_broadcasts_provider_stream_id_unique UNIQUE (provider_stream_id);
 �   ALTER TABLE ONLY public.event_live_provisioned_broadcasts DROP CONSTRAINT event_live_provisioned_broadcasts_provider_stream_id_unique;
       public         radiobox    false    195    195            t           2606    47874    event_live_routes_pkey 
   CONSTRAINT     p   ALTER TABLE ONLY event_live_routes
    ADD CONSTRAINT event_live_routes_pkey PRIMARY KEY (event_live_route_id);
 R   ALTER TABLE ONLY public.event_live_routes DROP CONSTRAINT event_live_routes_pkey;
       public         radiobox    false    196    196            v           2606    47876    event_live_segments_pkey 
   CONSTRAINT     v   ALTER TABLE ONLY event_live_segments
    ADD CONSTRAINT event_live_segments_pkey PRIMARY KEY (event_live_segment_id);
 V   ALTER TABLE ONLY public.event_live_segments DROP CONSTRAINT event_live_segments_pkey;
       public         radiobox    false    197    197            x           2606    47878    events_live_pkey 
   CONSTRAINT     ^   ALTER TABLE ONLY events_live
    ADD CONSTRAINT events_live_pkey PRIMARY KEY (event_live_id);
 F   ALTER TABLE ONLY public.events_live DROP CONSTRAINT events_live_pkey;
       public         radiobox    false    198    198            z           2606    47880 	   logs_pkey 
   CONSTRAINT     I   ALTER TABLE ONLY logs
    ADD CONSTRAINT logs_pkey PRIMARY KEY (log_id);
 8   ALTER TABLE ONLY public.logs DROP CONSTRAINT logs_pkey;
       public         radiobox    false    199    199            ~           2606    47882    media_formats_pkey 
   CONSTRAINT     d   ALTER TABLE ONLY media_formats
    ADD CONSTRAINT media_formats_pkey PRIMARY KEY (media_format_id);
 J   ALTER TABLE ONLY public.media_formats DROP CONSTRAINT media_formats_pkey;
       public         radiobox    false    202    202            |           2606    47884    medias_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY media
    ADD CONSTRAINT medias_pkey PRIMARY KEY (media_id);
 ;   ALTER TABLE ONLY public.media DROP CONSTRAINT medias_pkey;
       public         radiobox    false    201    201            �           2606    47886    oauth_access_data_pkey 
   CONSTRAINT     h   ALTER TABLE ONLY oauth_access_data
    ADD CONSTRAINT oauth_access_data_pkey PRIMARY KEY (accesstoken);
 R   ALTER TABLE ONLY public.oauth_access_data DROP CONSTRAINT oauth_access_data_pkey;
       public         radiobox    false    204    204            �           2606    47888    oauth_auth_data_pkey 
   CONSTRAINT     ]   ALTER TABLE ONLY oauth_auth_data
    ADD CONSTRAINT oauth_auth_data_pkey PRIMARY KEY (code);
 N   ALTER TABLE ONLY public.oauth_auth_data DROP CONSTRAINT oauth_auth_data_pkey;
       public         radiobox    false    205    205            �           2606    47890    oauth_clients_pkey 
   CONSTRAINT     W   ALTER TABLE ONLY oauth_clients
    ADD CONSTRAINT oauth_clients_pkey PRIMARY KEY (id);
 J   ALTER TABLE ONLY public.oauth_clients DROP CONSTRAINT oauth_clients_pkey;
       public         radiobox    false    206    206            �           2606    47892    payment_services_pkey 
   CONSTRAINT     m   ALTER TABLE ONLY payment_services
    ADD CONSTRAINT payment_services_pkey PRIMARY KEY (payment_service_id);
 P   ALTER TABLE ONLY public.payment_services DROP CONSTRAINT payment_services_pkey;
       public         radiobox    false    207    207            �           2606    47894    photos_pkey 
   CONSTRAINT     O   ALTER TABLE ONLY photos
    ADD CONSTRAINT photos_pkey PRIMARY KEY (photo_id);
 <   ALTER TABLE ONLY public.photos DROP CONSTRAINT photos_pkey;
       public         radiobox    false    208    208            �           2606    47896    purchase_items_pkey 
   CONSTRAINT     g   ALTER TABLE ONLY purchase_items
    ADD CONSTRAINT purchase_items_pkey PRIMARY KEY (purchase_item_id);
 L   ALTER TABLE ONLY public.purchase_items DROP CONSTRAINT purchase_items_pkey;
       public         radiobox    false    209    209            �           2606    47898    purchase_purchaseitem_pkey 
   CONSTRAINT     }   ALTER TABLE ONLY purchase_purchaseitem
    ADD CONSTRAINT purchase_purchaseitem_pkey PRIMARY KEY (purcahse_purchaseitem_id);
 Z   ALTER TABLE ONLY public.purchase_purchaseitem DROP CONSTRAINT purchase_purchaseitem_pkey;
       public         radiobox    false    210    210            �           2606    47900    purchases_pkey 
   CONSTRAINT     X   ALTER TABLE ONLY purchases
    ADD CONSTRAINT purchases_pkey PRIMARY KEY (purchase_id);
 B   ALTER TABLE ONLY public.purchases DROP CONSTRAINT purchases_pkey;
       public         radiobox    false    211    211            �           2606    47902 
   slugs_pkey 
   CONSTRAINT     G   ALTER TABLE ONLY slugs
    ADD CONSTRAINT slugs_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.slugs DROP CONSTRAINT slugs_pkey;
       public         radiobox    false    212    212            �           2606    47904    tokens_pkey 
   CONSTRAINT     O   ALTER TABLE ONLY tokens
    ADD CONSTRAINT tokens_pkey PRIMARY KEY (token_id);
 <   ALTER TABLE ONLY public.tokens DROP CONSTRAINT tokens_pkey;
       public         radiobox    false    214    214            �           2606    47906    tracks_pkey 
   CONSTRAINT     O   ALTER TABLE ONLY tracks
    ADD CONSTRAINT tracks_pkey PRIMARY KEY (track_id);
 <   ALTER TABLE ONLY public.tracks DROP CONSTRAINT tracks_pkey;
       public         radiobox    false    215    215            �           2606    47908    twitter_account_settings_pkey 
   CONSTRAINT     �   ALTER TABLE ONLY twitter_account_settings
    ADD CONSTRAINT twitter_account_settings_pkey PRIMARY KEY (twitter_account_settings_id);
 `   ALTER TABLE ONLY public.twitter_account_settings DROP CONSTRAINT twitter_account_settings_pkey;
       public         radiobox    false    216    216            �           2606    47910    user_checkins_pkey 
   CONSTRAINT     d   ALTER TABLE ONLY user_checkins
    ADD CONSTRAINT user_checkins_pkey PRIMARY KEY (user_checkin_id);
 J   ALTER TABLE ONLY public.user_checkins DROP CONSTRAINT user_checkins_pkey;
       public         radiobox    false    217    217            �           2606    47912    user_media_listens_pkey 
   CONSTRAINT     g   ALTER TABLE ONLY user_listens
    ADD CONSTRAINT user_media_listens_pkey PRIMARY KEY (user_listen_id);
 N   ALTER TABLE ONLY public.user_listens DROP CONSTRAINT user_media_listens_pkey;
       public         radiobox    false    218    218            �           2606    47914    user_notifications_pkey 
   CONSTRAINT     s   ALTER TABLE ONLY user_notifications
    ADD CONSTRAINT user_notifications_pkey PRIMARY KEY (user_notification_id);
 T   ALTER TABLE ONLY public.user_notifications DROP CONSTRAINT user_notifications_pkey;
       public         radiobox    false    219    219            �           2606    47916    user_payments_pkey 
   CONSTRAINT     d   ALTER TABLE ONLY user_payments
    ADD CONSTRAINT user_payments_pkey PRIMARY KEY (user_payment_id);
 J   ALTER TABLE ONLY public.user_payments DROP CONSTRAINT user_payments_pkey;
       public         radiobox    false    220    220            �           2606    47918    usernames_pkey 
   CONSTRAINT     O   ALTER TABLE ONLY usernames
    ADD CONSTRAINT usernames_pkey PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.usernames DROP CONSTRAINT usernames_pkey;
       public         radiobox    false    221    221            �           2606    47920    usernames_username_key 
   CONSTRAINT     X   ALTER TABLE ONLY usernames
    ADD CONSTRAINT usernames_username_key UNIQUE (username);
 J   ALTER TABLE ONLY public.usernames DROP CONSTRAINT usernames_username_key;
       public         radiobox    false    221    221            �           2606    47922    users_email_unique 
   CONSTRAINT     M   ALTER TABLE ONLY users
    ADD CONSTRAINT users_email_unique UNIQUE (email);
 B   ALTER TABLE ONLY public.users DROP CONSTRAINT users_email_unique;
       public         radiobox    false    223    223            �           2606    47924 
   users_pkey 
   CONSTRAINT     L   ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public         radiobox    false    223    223            �           2606    47926    users_username_unique 
   CONSTRAINT     S   ALTER TABLE ONLY users
    ADD CONSTRAINT users_username_unique UNIQUE (username);
 E   ALTER TABLE ONLY public.users DROP CONSTRAINT users_username_unique;
       public         radiobox    false    223    223            �           2606    47928    venues_pkey 
   CONSTRAINT     O   ALTER TABLE ONLY venues
    ADD CONSTRAINT venues_pkey PRIMARY KEY (venue_id);
 <   ALTER TABLE ONLY public.venues DROP CONSTRAINT venues_pkey;
       public         radiobox    false    225    225            �           2606    47930    voucher_purchase_pkey 
   CONSTRAINT     n   ALTER TABLE ONLY voucher_purchase
    ADD CONSTRAINT voucher_purchase_pkey PRIMARY KEY (voucher_purchase_id);
 P   ALTER TABLE ONLY public.voucher_purchase DROP CONSTRAINT voucher_purchase_pkey;
       public         radiobox    false    227    227            �           2606    47932    vouchers_pkey 
   CONSTRAINT     U   ALTER TABLE ONLY vouchers
    ADD CONSTRAINT vouchers_pkey PRIMARY KEY (voucher_id);
 @   ALTER TABLE ONLY public.vouchers DROP CONSTRAINT vouchers_pkey;
       public         radiobox    false    228    228            X           1259    47933 (   album_track_album_id_track_number_unique    INDEX     r   CREATE UNIQUE INDEX album_track_album_id_track_number_unique ON album_track USING btree (album_id, track_number);
 <   DROP INDEX public.album_track_album_id_track_number_unique;
       public         radiobox    false    183    183            �           1259    47934    slugs_case_insensitive_slug_key    INDEX     `   CREATE UNIQUE INDEX slugs_case_insensitive_slug_key ON slugs USING btree (lower((slug)::text));
 3   DROP INDEX public.slugs_case_insensitive_slug_key;
       public         radiobox    false    212    212            �           1259    47935    slugs_target_artist_id    INDEX     h   CREATE INDEX slugs_target_artist_id ON slugs USING btree ((((target ->> 'artist_id'::text))::integer));
 *   DROP INDEX public.slugs_target_artist_id;
       public         radiobox    false    212    212            �           1259    47936    slugs_target_user_id    INDEX     d   CREATE INDEX slugs_target_user_id ON slugs USING btree ((((target ->> 'user_id'::text))::integer));
 (   DROP INDEX public.slugs_target_user_id;
       public         radiobox    false    212    212            �           1259    47937    tokens_type    INDEX     7   CREATE INDEX tokens_type ON tokens USING btree (type);
    DROP INDEX public.tokens_type;
       public         radiobox    false    214            �           1259    47938    tracks_encoding_media_id_idx    INDEX     U   CREATE INDEX tracks_encoding_media_id_idx ON tracks USING btree (encoding_media_id);
 0   DROP INDEX public.tracks_encoding_media_id_idx;
       public         radiobox    false    215            �           1259    47939    usernames_target_artist_id    INDEX     p   CREATE INDEX usernames_target_artist_id ON usernames USING btree ((((target ->> 'artist_id'::text))::integer));
 .   DROP INDEX public.usernames_target_artist_id;
       public         radiobox    false    221    221            �           1259    47940    usernames_target_user_id    INDEX     l   CREATE INDEX usernames_target_user_id ON usernames USING btree ((((target ->> 'user_id'::text))::integer));
 ,   DROP INDEX public.usernames_target_user_id;
       public         radiobox    false    221    221            �           2620    47941    artists_username_insert    TRIGGER     z   CREATE TRIGGER artists_username_insert BEFORE INSERT ON artists FOR EACH ROW EXECUTE PROCEDURE artists_username_insert();
 8   DROP TRIGGER artists_username_insert ON public.artists;
       public       radiobox    false    1305    188            �           2620    47942    artists_username_update    TRIGGER     z   CREATE TRIGGER artists_username_update BEFORE UPDATE ON artists FOR EACH ROW EXECUTE PROCEDURE artists_username_update();
 8   DROP TRIGGER artists_username_update ON public.artists;
       public       radiobox    false    1306    188            �           2620    47943    users_username_insert    TRIGGER     t   CREATE TRIGGER users_username_insert BEFORE INSERT ON users FOR EACH ROW EXECUTE PROCEDURE users_username_insert();
 4   DROP TRIGGER users_username_insert ON public.users;
       public       radiobox    false    1307    223            �           2620    47944    users_username_update    TRIGGER     t   CREATE TRIGGER users_username_update BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE users_username_update();
 4   DROP TRIGGER users_username_update ON public.users;
       public       radiobox    false    223    1308            �           2606    47945    album_track_album_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY album_track
    ADD CONSTRAINT album_track_album_id_foreign FOREIGN KEY (album_id) REFERENCES albums(album_id);
 R   ALTER TABLE ONLY public.album_track DROP CONSTRAINT album_track_album_id_foreign;
       public       radiobox    false    183    3420    184            �           2606    47950    album_track_track_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY album_track
    ADD CONSTRAINT album_track_track_id_foreign FOREIGN KEY (track_id) REFERENCES tracks(track_id);
 R   ALTER TABLE ONLY public.album_track DROP CONSTRAINT album_track_track_id_foreign;
       public       radiobox    false    3481    183    215            �           2606    47955    albums_artist_id_foreign    FK CONSTRAINT     {   ALTER TABLE ONLY albums
    ADD CONSTRAINT albums_artist_id_foreign FOREIGN KEY (artist_id) REFERENCES artists(artist_id);
 I   ALTER TABLE ONLY public.albums DROP CONSTRAINT albums_artist_id_foreign;
       public       radiobox    false    184    3428    188            �           2606    47960    albums_pic_square_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY albums
    ADD CONSTRAINT albums_pic_square_id_foreign FOREIGN KEY (pic_square_id) REFERENCES photos(photo_id);
 M   ALTER TABLE ONLY public.albums DROP CONSTRAINT albums_pic_square_id_foreign;
       public       radiobox    false    208    184    3464            �           2606    47965 !   artist_checkins_artist_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artist_checkins
    ADD CONSTRAINT artist_checkins_artist_id_foreign FOREIGN KEY (artist_id) REFERENCES artists(artist_id);
 [   ALTER TABLE ONLY public.artist_checkins DROP CONSTRAINT artist_checkins_artist_id_foreign;
       public       radiobox    false    3428    188    185            �           2606    47970 7   artist_checkins_checkin_location_venue_match_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artist_checkins
    ADD CONSTRAINT artist_checkins_checkin_location_venue_match_id_foreign FOREIGN KEY (checkin_location_venue_match_id) REFERENCES venues(venue_id);
 q   ALTER TABLE ONLY public.artist_checkins DROP CONSTRAINT artist_checkins_checkin_location_venue_match_id_foreign;
       public       radiobox    false    3505    185    225            �           2606    47975 "   artist_checkins_pic_big_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artist_checkins
    ADD CONSTRAINT artist_checkins_pic_big_id_foreign FOREIGN KEY (pic_big_id) REFERENCES photos(photo_id);
 \   ALTER TABLE ONLY public.artist_checkins DROP CONSTRAINT artist_checkins_pic_big_id_foreign;
       public       radiobox    false    3464    208    185            �           2606    47980 $   artist_checkins_pic_small_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artist_checkins
    ADD CONSTRAINT artist_checkins_pic_small_id_foreign FOREIGN KEY (pic_small_id) REFERENCES photos(photo_id);
 ^   ALTER TABLE ONLY public.artist_checkins DROP CONSTRAINT artist_checkins_pic_small_id_foreign;
       public       radiobox    false    185    3464    208            �           2606    47985 %   artist_checkins_pic_square_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artist_checkins
    ADD CONSTRAINT artist_checkins_pic_square_id_foreign FOREIGN KEY (pic_square_id) REFERENCES photos(photo_id);
 _   ALTER TABLE ONLY public.artist_checkins DROP CONSTRAINT artist_checkins_pic_square_id_foreign;
       public       radiobox    false    185    208    3464            �           2606    47990 &   artist_notifications_artist_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artist_notifications
    ADD CONSTRAINT artist_notifications_artist_id_foreign FOREIGN KEY (artist_id) REFERENCES artists(artist_id);
 e   ALTER TABLE ONLY public.artist_notifications DROP CONSTRAINT artist_notifications_artist_id_foreign;
       public       radiobox    false    3428    188    186            �           2606    47995    artist_user_artist_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artist_user
    ADD CONSTRAINT artist_user_artist_id_foreign FOREIGN KEY (artist_id) REFERENCES artists(artist_id);
 S   ALTER TABLE ONLY public.artist_user DROP CONSTRAINT artist_user_artist_id_foreign;
       public       radiobox    false    3428    188    187            �           2606    48000    artist_user_user_id_foreign    FK CONSTRAINT     }   ALTER TABLE ONLY artist_user
    ADD CONSTRAINT artist_user_user_id_foreign FOREIGN KEY (user_id) REFERENCES users(user_id);
 Q   ALTER TABLE ONLY public.artist_user DROP CONSTRAINT artist_user_user_id_foreign;
       public       radiobox    false    3501    223    187            �           2606    48005    artists_booking_user_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artists
    ADD CONSTRAINT artists_booking_user_id_foreign FOREIGN KEY (booking_user_id) REFERENCES users(user_id);
 Q   ALTER TABLE ONLY public.artists DROP CONSTRAINT artists_booking_user_id_foreign;
       public       radiobox    false    3501    188    223            �           2606    48010    artists_manager_user_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artists
    ADD CONSTRAINT artists_manager_user_id_foreign FOREIGN KEY (manager_user_id) REFERENCES users(user_id);
 Q   ALTER TABLE ONLY public.artists DROP CONSTRAINT artists_manager_user_id_foreign;
       public       radiobox    false    188    223    3501            �           2606    48015    artists_pic_cover_id_fkey    FK CONSTRAINT     ~   ALTER TABLE ONLY artists
    ADD CONSTRAINT artists_pic_cover_id_fkey FOREIGN KEY (pic_cover_id) REFERENCES photos(photo_id);
 K   ALTER TABLE ONLY public.artists DROP CONSTRAINT artists_pic_cover_id_fkey;
       public       radiobox    false    3464    188    208            �           2606    48020    artists_pic_square_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artists
    ADD CONSTRAINT artists_pic_square_id_foreign FOREIGN KEY (pic_square_id) REFERENCES photos(photo_id);
 O   ALTER TABLE ONLY public.artists DROP CONSTRAINT artists_pic_square_id_foreign;
       public       radiobox    false    3464    188    208            �           2606    48025 +   artists_twitter_account_settings_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artists
    ADD CONSTRAINT artists_twitter_account_settings_id_foreign FOREIGN KEY (twitter_account_settings_id) REFERENCES twitter_account_settings(twitter_account_settings_id);
 ]   ALTER TABLE ONLY public.artists DROP CONSTRAINT artists_twitter_account_settings_id_foreign;
       public       radiobox    false    188    216    3483            �           2606    48030 %   event_admin_log_event_live_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY event_admin_log
    ADD CONSTRAINT event_admin_log_event_live_id_foreign FOREIGN KEY (event_live_id) REFERENCES events_live(event_live_id);
 _   ALTER TABLE ONLY public.event_admin_log DROP CONSTRAINT event_admin_log_event_live_id_foreign;
       public       radiobox    false    193    198    3448            �           2606    48035 .   event_live_notifications_event_live_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY event_live_notifications
    ADD CONSTRAINT event_live_notifications_event_live_id_foreign FOREIGN KEY (event_live_id) REFERENCES events_live(event_live_id);
 q   ALTER TABLE ONLY public.event_live_notifications DROP CONSTRAINT event_live_notifications_event_live_id_foreign;
       public       radiobox    false    194    198    3448            �           2606    48040 '   event_live_routes_event_live_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY event_live_routes
    ADD CONSTRAINT event_live_routes_event_live_id_foreign FOREIGN KEY (event_live_id) REFERENCES events_live(event_live_id);
 c   ALTER TABLE ONLY public.event_live_routes DROP CONSTRAINT event_live_routes_event_live_id_foreign;
       public       radiobox    false    196    198    3448            �           2606    48045 =   event_live_routes_event_live_provisioned_broadcast_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY event_live_routes
    ADD CONSTRAINT event_live_routes_event_live_provisioned_broadcast_id_foreign FOREIGN KEY (event_live_provisioned_broadcast_id) REFERENCES event_live_provisioned_broadcasts(event_live_provisioned_broadcast_id);
 y   ALTER TABLE ONLY public.event_live_routes DROP CONSTRAINT event_live_routes_event_live_provisioned_broadcast_id_foreign;
       public       radiobox    false    3440    195    196            �           2606    48050 )   event_live_segments_event_live_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY event_live_segments
    ADD CONSTRAINT event_live_segments_event_live_id_foreign FOREIGN KEY (event_live_id) REFERENCES events_live(event_live_id);
 g   ALTER TABLE ONLY public.event_live_segments DROP CONSTRAINT event_live_segments_event_live_id_foreign;
       public       radiobox    false    197    198    3448            �           2606    48055 $   events_live_archive_track_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY events_live
    ADD CONSTRAINT events_live_archive_track_id_foreign FOREIGN KEY (archive_track_id) REFERENCES tracks(track_id);
 Z   ALTER TABLE ONLY public.events_live DROP CONSTRAINT events_live_archive_track_id_foreign;
       public       radiobox    false    198    215    3481            �           2606    48060    events_live_artist_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY events_live
    ADD CONSTRAINT events_live_artist_id_foreign FOREIGN KEY (artist_id) REFERENCES artists(artist_id);
 S   ALTER TABLE ONLY public.events_live DROP CONSTRAINT events_live_artist_id_foreign;
       public       radiobox    false    198    188    3428            �           2606    48268    events_live_pic_square_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY events_live
    ADD CONSTRAINT events_live_pic_square_id_fkey FOREIGN KEY (pic_square_id) REFERENCES photos(photo_id);
 T   ALTER TABLE ONLY public.events_live DROP CONSTRAINT events_live_pic_square_id_fkey;
       public       radiobox    false    198    3464    208            �           2606    48075    events_live_venue_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY events_live
    ADD CONSTRAINT events_live_venue_id_foreign FOREIGN KEY (venue_id) REFERENCES venues(venue_id);
 R   ALTER TABLE ONLY public.events_live DROP CONSTRAINT events_live_venue_id_foreign;
       public       radiobox    false    3505    198    225            �           2606    48080    media_formats_pic_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY media_formats
    ADD CONSTRAINT media_formats_pic_id_foreign FOREIGN KEY (pic_id) REFERENCES photos(photo_id);
 T   ALTER TABLE ONLY public.media_formats DROP CONSTRAINT media_formats_pic_id_foreign;
       public       radiobox    false    202    208    3464            �           2606    48085    medias_media_format_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY media
    ADD CONSTRAINT medias_media_format_id_foreign FOREIGN KEY (media_format_id) REFERENCES media_formats(media_format_id);
 N   ALTER TABLE ONLY public.media DROP CONSTRAINT medias_media_format_id_foreign;
       public       radiobox    false    202    3454    201            �           2606    48090 "   purchase_items_purchase_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY purchase_items
    ADD CONSTRAINT purchase_items_purchase_id_foreign FOREIGN KEY (purchase_id) REFERENCES purchases(purchase_id);
 [   ALTER TABLE ONLY public.purchase_items DROP CONSTRAINT purchase_items_purchase_id_foreign;
       public       radiobox    false    3470    209    211            �           2606    48095 )   purchase_purchaseitem_purchase_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY purchase_purchaseitem
    ADD CONSTRAINT purchase_purchaseitem_purchase_id_foreign FOREIGN KEY (purchase_id) REFERENCES purchases(purchase_id);
 i   ALTER TABLE ONLY public.purchase_purchaseitem DROP CONSTRAINT purchase_purchaseitem_purchase_id_foreign;
       public       radiobox    false    211    3470    210            �           2606    48100 .   purchase_purchaseitem_purchase_item_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY purchase_purchaseitem
    ADD CONSTRAINT purchase_purchaseitem_purchase_item_id_foreign FOREIGN KEY (purchase_item_id) REFERENCES purchase_items(purchase_item_id);
 n   ALTER TABLE ONLY public.purchase_purchaseitem DROP CONSTRAINT purchase_purchaseitem_purchase_item_id_foreign;
       public       radiobox    false    3466    210    209            �           2606    48105    purchases_user_id_foreign    FK CONSTRAINT     y   ALTER TABLE ONLY purchases
    ADD CONSTRAINT purchases_user_id_foreign FOREIGN KEY (user_id) REFERENCES users(user_id);
 M   ALTER TABLE ONLY public.purchases DROP CONSTRAINT purchases_user_id_foreign;
       public       radiobox    false    211    223    3501            �           2606    48110 !   purchases_user_payment_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY purchases
    ADD CONSTRAINT purchases_user_payment_id_foreign FOREIGN KEY (user_payment_id) REFERENCES user_payments(user_payment_id);
 U   ALTER TABLE ONLY public.purchases DROP CONSTRAINT purchases_user_payment_id_foreign;
       public       radiobox    false    220    211    3491            �           2606    48115    tokens_user_id_fkey    FK CONSTRAINT     p   ALTER TABLE ONLY tokens
    ADD CONSTRAINT tokens_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(user_id);
 D   ALTER TABLE ONLY public.tokens DROP CONSTRAINT tokens_user_id_fkey;
       public       radiobox    false    214    223    3501            �           2606    48120    tracks_artist_id_foreign    FK CONSTRAINT     {   ALTER TABLE ONLY tracks
    ADD CONSTRAINT tracks_artist_id_foreign FOREIGN KEY (artist_id) REFERENCES artists(artist_id);
 I   ALTER TABLE ONLY public.tracks DROP CONSTRAINT tracks_artist_id_foreign;
       public       radiobox    false    215    188    3428            �           2606    48125    tracks_media_id_foreign    FK CONSTRAINT     v   ALTER TABLE ONLY tracks
    ADD CONSTRAINT tracks_media_id_foreign FOREIGN KEY (media_id) REFERENCES media(media_id);
 H   ALTER TABLE ONLY public.tracks DROP CONSTRAINT tracks_media_id_foreign;
       public       radiobox    false    215    201    3452            �           2606    48130 #   tracks_must_bundle_album_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY tracks
    ADD CONSTRAINT tracks_must_bundle_album_id_foreign FOREIGN KEY (must_bundle_album_id) REFERENCES albums(album_id);
 T   ALTER TABLE ONLY public.tracks DROP CONSTRAINT tracks_must_bundle_album_id_foreign;
       public       radiobox    false    215    184    3420            �           2606    48135    tracks_pic_square_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY tracks
    ADD CONSTRAINT tracks_pic_square_id_foreign FOREIGN KEY (pic_square_id) REFERENCES photos(photo_id);
 M   ALTER TABLE ONLY public.tracks DROP CONSTRAINT tracks_pic_square_id_foreign;
       public       radiobox    false    215    208    3464            �           2606    48140    tracks_preview_media_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY tracks
    ADD CONSTRAINT tracks_preview_media_id_foreign FOREIGN KEY (preview_media_id) REFERENCES media(media_id);
 P   ALTER TABLE ONLY public.tracks DROP CONSTRAINT tracks_preview_media_id_foreign;
       public       radiobox    false    215    201    3452            �           2606    48145 5   user_checkins_checkin_location_venue_match_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY user_checkins
    ADD CONSTRAINT user_checkins_checkin_location_venue_match_id_foreign FOREIGN KEY (checkin_location_venue_match_id) REFERENCES venues(venue_id);
 m   ALTER TABLE ONLY public.user_checkins DROP CONSTRAINT user_checkins_checkin_location_venue_match_id_foreign;
       public       radiobox    false    217    225    3505            �           2606    48150     user_checkins_pic_big_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY user_checkins
    ADD CONSTRAINT user_checkins_pic_big_id_foreign FOREIGN KEY (pic_big_id) REFERENCES photos(photo_id);
 X   ALTER TABLE ONLY public.user_checkins DROP CONSTRAINT user_checkins_pic_big_id_foreign;
       public       radiobox    false    217    208    3464            �           2606    48155 "   user_checkins_pic_small_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY user_checkins
    ADD CONSTRAINT user_checkins_pic_small_id_foreign FOREIGN KEY (pic_small_id) REFERENCES photos(photo_id);
 Z   ALTER TABLE ONLY public.user_checkins DROP CONSTRAINT user_checkins_pic_small_id_foreign;
       public       radiobox    false    208    217    3464            �           2606    48160 #   user_checkins_pic_square_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY user_checkins
    ADD CONSTRAINT user_checkins_pic_square_id_foreign FOREIGN KEY (pic_square_id) REFERENCES photos(photo_id);
 [   ALTER TABLE ONLY public.user_checkins DROP CONSTRAINT user_checkins_pic_square_id_foreign;
       public       radiobox    false    3464    208    217            �           2606    48165    user_checkins_user_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY user_checkins
    ADD CONSTRAINT user_checkins_user_id_foreign FOREIGN KEY (user_id) REFERENCES users(user_id);
 U   ALTER TABLE ONLY public.user_checkins DROP CONSTRAINT user_checkins_user_id_foreign;
       public       radiobox    false    217    3501    223            �           2606    48170 "   user_notifications_user_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY user_notifications
    ADD CONSTRAINT user_notifications_user_id_foreign FOREIGN KEY (user_id) REFERENCES users(user_id);
 _   ALTER TABLE ONLY public.user_notifications DROP CONSTRAINT user_notifications_user_id_foreign;
       public       radiobox    false    223    219    3501            �           2606    48175 (   user_payments_payment_service_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY user_payments
    ADD CONSTRAINT user_payments_payment_service_id_foreign FOREIGN KEY (payment_service_id) REFERENCES payment_services(payment_service_id);
 `   ALTER TABLE ONLY public.user_payments DROP CONSTRAINT user_payments_payment_service_id_foreign;
       public       radiobox    false    220    207    3462            �           2606    48180    user_payments_user_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY user_payments
    ADD CONSTRAINT user_payments_user_id_foreign FOREIGN KEY (user_id) REFERENCES users(user_id);
 U   ALTER TABLE ONLY public.user_payments DROP CONSTRAINT user_payments_user_id_foreign;
       public       radiobox    false    220    223    3501            �           2606    48185 %   users_current_listen_media_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY users
    ADD CONSTRAINT users_current_listen_media_id_foreign FOREIGN KEY (current_listen_media_id) REFERENCES media(media_id);
 U   ALTER TABLE ONLY public.users DROP CONSTRAINT users_current_listen_media_id_foreign;
       public       radiobox    false    223    201    3452            �           2606    48190    users_pic_cover_id_foreign    FK CONSTRAINT     }   ALTER TABLE ONLY users
    ADD CONSTRAINT users_pic_cover_id_foreign FOREIGN KEY (pic_cover_id) REFERENCES photos(photo_id);
 J   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pic_cover_id_foreign;
       public       radiobox    false    223    208    3464            �           2606    48195    users_pic_square_id_foreign    FK CONSTRAINT        ALTER TABLE ONLY users
    ADD CONSTRAINT users_pic_square_id_foreign FOREIGN KEY (pic_square_id) REFERENCES photos(photo_id);
 K   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pic_square_id_foreign;
       public       radiobox    false    223    208    3464            �           2606    48200 )   users_twitter_account_settings_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY users
    ADD CONSTRAINT users_twitter_account_settings_id_foreign FOREIGN KEY (twitter_account_settings_id) REFERENCES twitter_account_settings(twitter_account_settings_id);
 Y   ALTER TABLE ONLY public.users DROP CONSTRAINT users_twitter_account_settings_id_foreign;
       public       radiobox    false    223    216    3483            �           2606    48205    users_user_payment_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY users
    ADD CONSTRAINT users_user_payment_id_foreign FOREIGN KEY (user_payment_id) REFERENCES user_payments(user_payment_id);
 M   ALTER TABLE ONLY public.users DROP CONSTRAINT users_user_payment_id_foreign;
       public       radiobox    false    223    220    3491            �           2606    48210    venues_pic_cover_id_foreign    FK CONSTRAINT        ALTER TABLE ONLY venues
    ADD CONSTRAINT venues_pic_cover_id_foreign FOREIGN KEY (pic_cover_id) REFERENCES photos(photo_id);
 L   ALTER TABLE ONLY public.venues DROP CONSTRAINT venues_pic_cover_id_foreign;
       public       radiobox    false    225    208    3464            �           2606    48215    venues_pic_square_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY venues
    ADD CONSTRAINT venues_pic_square_id_foreign FOREIGN KEY (pic_square_id) REFERENCES photos(photo_id);
 M   ALTER TABLE ONLY public.venues DROP CONSTRAINT venues_pic_square_id_foreign;
       public       radiobox    false    3464    225    208            �           2606    48220 $   voucher_purchase_purchase_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY voucher_purchase
    ADD CONSTRAINT voucher_purchase_purchase_id_foreign FOREIGN KEY (purchase_id) REFERENCES purchases(purchase_id);
 _   ALTER TABLE ONLY public.voucher_purchase DROP CONSTRAINT voucher_purchase_purchase_id_foreign;
       public       radiobox    false    227    211    3470            �           2606    48225 #   voucher_purchase_voucher_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY voucher_purchase
    ADD CONSTRAINT voucher_purchase_voucher_id_foreign FOREIGN KEY (voucher_id) REFERENCES vouchers(voucher_id);
 ^   ALTER TABLE ONLY public.voucher_purchase DROP CONSTRAINT voucher_purchase_voucher_id_foreign;
       public       radiobox    false    228    227    3509            �           2606    48230 !   vouchers_issuer_artist_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY vouchers
    ADD CONSTRAINT vouchers_issuer_artist_id_foreign FOREIGN KEY (issuer_artist_id) REFERENCES artists(artist_id);
 T   ALTER TABLE ONLY public.vouchers DROP CONSTRAINT vouchers_issuer_artist_id_foreign;
       public       radiobox    false    188    3428    228            e      x������ � �      f      x������ � �      g      x������ � �      h      x������ � �      i      x������ � �      j      x������ � �      l      x������ � �      n      x������ � �      o      x������ � �      p      x������ � �      q      x������ � �      r      x������ � �      s      x������ � �      t      x������ � �      u   �   x�}��
�0D��W��چ�F�9��(�E(�]1�&�D)���������Yfe���nL��(���0�0��C!g��n'�m���~�5��{ec��<�F3[�+��o��	_�юu���J,e���h)�=�&����J�z>U'�j�j���;��Z���J��F���]9c�"7F�      w      x������ � �      x      x������ � �      y      x������ � �      z      x������ � �      {      x������ � �      |      x�30 CΒ��0����� IJ�      }      x������ � �      ~      x������ � �            x������ � �      �      x������ � �      �      x������ � �      �   !   x�3�V*-N-��LQ�2��L,N����� b�      �      x������ � �      �   \   x���
�  �g���c�֦ٷak��"���6Ե� �Ԁ��j�2�&N��QD��@�C��L��H�>Fggm�������m�����O'�      �      x������ � �      �      x������ � �      �      x������ � �      �      x������ � �      �      x������ � �      �      x������ � �      �   !   x�3�V*-N-��LQ�2��L,N����� b�      �   �   x�3��Â�tbq��%�cWD*F�*�*��)��aUa9���a�f���F��E%�ő��)��I�!����YIɑ��!���8%�����J8�8�MtLt�,M�L,�����L�-��K��q��qqq ��@�      �      x������ � �      �      x������ � �      �      x������ � �     