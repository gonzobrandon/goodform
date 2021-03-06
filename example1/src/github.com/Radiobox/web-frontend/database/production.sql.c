PGDMP     6                    r            d7fe4r5qfb3f6e    9.3.2    9.3.3 �    �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                       false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                       false            �           1262    16385    d7fe4r5qfb3f6e    DATABASE     �   CREATE DATABASE d7fe4r5qfb3f6e WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';
    DROP DATABASE d7fe4r5qfb3f6e;
             ueoddrluibf30n    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
             ueoddrluibf30n    false            �           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                  ueoddrluibf30n    false    6            �            3079    16388    plpgsql 	   EXTENSION     ?   CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;
    DROP EXTENSION plpgsql;
                  false            �           0    0    EXTENSION plpgsql    COMMENT     @   COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';
                       false    225            �            3079    24576    pg_stat_statements 	   EXTENSION     F   CREATE EXTENSION IF NOT EXISTS pg_stat_statements WITH SCHEMA public;
 #   DROP EXTENSION pg_stat_statements;
                  false    6            �           0    0    EXTENSION pg_stat_statements    COMMENT     h   COMMENT ON EXTENSION pg_stat_statements IS 'track execution statistics of all SQL statements executed';
                       false    228            �            3079    16424    postgis 	   EXTENSION     ;   CREATE EXTENSION IF NOT EXISTS postgis WITH SCHEMA public;
    DROP EXTENSION postgis;
                  false    6            �           0    0    EXTENSION postgis    COMMENT     g   COMMENT ON EXTENSION postgis IS 'PostGIS geometry, geography, and raster spatial types and functions';
                       false    227            �            3079    17708 	   uuid-ossp 	   EXTENSION     ?   CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;
    DROP EXTENSION "uuid-ossp";
                  false    6            �           0    0    EXTENSION "uuid-ossp"    COMMENT     W   COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';
                       false    226                       1255    24617    artists_username_insert()    FUNCTION     �   CREATE FUNCTION artists_username_insert() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    BEGIN
        INSERT INTO usernames (target, username) VALUES (('{"artist_id":' || NEW.artist_id || '}')::json, NEW.username);
        RETURN NEW;
    END;
$$;
 0   DROP FUNCTION public.artists_username_insert();
       public       ueoddrluibf30n    false    6    225                       1255    24619    artists_username_update()    FUNCTION     �   CREATE FUNCTION artists_username_update() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    BEGIN
        UPDATE usernames SET username = NEW.username WHERE (usernames.target->>'artist_id')::int = NEW.artist_id;
        RETURN NEW;
    END;
$$;
 0   DROP FUNCTION public.artists_username_update();
       public       ueoddrluibf30n    false    6    225                       1255    24613    users_username_insert()    FUNCTION     �   CREATE FUNCTION users_username_insert() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    BEGIN
        INSERT INTO usernames (target, username) VALUES (('{"user_id":' || NEW.user_id || '}')::json, NEW.username);
        RETURN NEW;
    END;
$$;
 .   DROP FUNCTION public.users_username_insert();
       public       ueoddrluibf30n    false    225    6            �           1255    24615    users_username_update()    FUNCTION     �   CREATE FUNCTION users_username_update() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    BEGIN
        UPDATE usernames SET username = NEW.username WHERE (usernames.target->>'user_id')::int = NEW.user_id;
        RETURN NEW;
    END;
$$;
 .   DROP FUNCTION public.users_username_update();
       public       ueoddrluibf30n    false    6    225            �            1259    17719    album_track    TABLE       CREATE TABLE album_track (
    album_track_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    album_id uuid NOT NULL,
    track_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
    DROP TABLE public.album_track;
       public         ueoddrluibf30n    false    226    6    6            �            1259    17725    albums    TABLE        CREATE TABLE albums (
    album_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    artist_id bigint,
    downloads bigint DEFAULT (0)::bigint,
    title character varying(255),
    pic_big character varying(255),
    pic_big_id uuid,
    pic_small character varying(255),
    pic_small_id uuid,
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
       public         ueoddrluibf30n    false    226    6    6            �            1259    17741    artist_checkins    TABLE     Y  CREATE TABLE artist_checkins (
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
       public         ueoddrluibf30n    false    226    6    227    227    6    227    6    227    6    227    6    227    6    227    6    227    6    6    6            �            1259    17750    artist_notifications    TABLE     *  CREATE TABLE artist_notifications (
    artist_notification_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    artist_id bigint NOT NULL,
    notification json NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 (   DROP TABLE public.artist_notifications;
       public         ueoddrluibf30n    false    226    6    6            �            1259    17759    artist_user    TABLE     R  CREATE TABLE artist_user (
    artist_user_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    artist_id bigint NOT NULL,
    user_id bigint NOT NULL,
    is_admin boolean DEFAULT false NOT NULL,
    role json,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
    DROP TABLE public.artist_user;
       public         ueoddrluibf30n    false    226    6    6            �            1259    17769    artists    TABLE     E  CREATE TABLE artists (
    artist_id bigint NOT NULL,
    band_user_ids json,
    booking_user_id bigint,
    contact_email character varying(255) NOT NULL,
    contact_phone character varying(255),
    contact_phone_country_code smallint DEFAULT (1)::smallint NOT NULL,
    contact_phone_extension character varying(10),
    facebook_page_id character varying(255),
    facebook_page_token_encrypted character varying(255),
    hometown_address json,
    hometown_address_point geography(Point,4326),
    keywords json,
    manager_user_id bigint,
    pic character varying(255),
    pic_id uuid,
    pic_big character varying(255),
    pic_big_id uuid,
    pic_cover character varying(255),
    pic_small character varying(255),
    pic_small_id uuid,
    pic_square character varying(255),
    pic_square_id uuid,
    record_label_id character varying(255),
    slug character varying(255),
    subscriber_count bigint,
    timezone character varying(255),
    twitter_account_settings_id uuid,
    username character varying(255) NOT NULL,
    website character varying(255),
    is_active boolean DEFAULT false NOT NULL,
    is_verified boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    pic_cover_id uuid,
    description text
);
    DROP TABLE public.artists;
       public         ueoddrluibf30n    false    6    227    227    6    227    6    227    6    227    6    227    6    227    6    227    6    6            �            1259    17780    artists_artist_id_seq    SEQUENCE     w   CREATE SEQUENCE artists_artist_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 ,   DROP SEQUENCE public.artists_artist_id_seq;
       public       ueoddrluibf30n    false    188    6            �           0    0    artists_artist_id_seq    SEQUENCE OWNED BY     A   ALTER SEQUENCE artists_artist_id_seq OWNED BY artists.artist_id;
            public       ueoddrluibf30n    false    189            �            1259    17782    event_admin_log    TABLE       CREATE TABLE event_admin_log (
    event_admin_log_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    event_live_id uuid NOT NULL,
    data json NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 #   DROP TABLE public.event_admin_log;
       public         ueoddrluibf30n    false    226    6    6            �            1259    17791    event_live_notifications    TABLE     S  CREATE TABLE event_live_notifications (
    event_live_notification_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    event_live_id uuid NOT NULL,
    notification json NOT NULL,
    title json NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);
 ,   DROP TABLE public.event_live_notifications;
       public         ueoddrluibf30n    false    226    6    6            �            1259    17800 !   event_live_provisioned_broadcasts    TABLE     �  CREATE TABLE event_live_provisioned_broadcasts (
    event_live_provisioned_broadcast_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    type character varying(255),
    is_video boolean DEFAULT false NOT NULL,
    boadcast_url_1_port smallint,
    boadcast_url_2_port smallint,
    broadcast_stream_name character varying(255),
    provider_stream_id character varying(255),
    broadcast_url_1 character varying(255),
    broadcast_url_2 character varying(255),
    broascast_method character varying(255),
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
       public         ueoddrluibf30n    false    226    6    6            �            1259    17815    event_live_routes    TABLE       CREATE TABLE event_live_routes (
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
       public         ueoddrluibf30n    false    226    6    6            �            1259    17825    event_live_segments    TABLE     �  CREATE TABLE event_live_segments (
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
       public         ueoddrluibf30n    false    226    6    6            �            1259    17834    events_live    TABLE     �  CREATE TABLE events_live (
    event_live_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    actual_end timestamp with time zone,
    actual_start timestamp with time zone,
    archive_track_id uuid,
    archiving boolean,
    artist_id bigint,
    name character varying(255),
    name_live_from character varying(255),
    artist_collaborators_id json,
    is_in_progress boolean,
    is_concluded boolean,
    off_schedule_end_seconds bigint,
    off_schedule_start_seconds bigint,
    pic_big character varying(255),
    pic_big_id uuid,
    pic_small character varying(255),
    pic_small_id uuid,
    scheduled_end timestamp with time zone,
    scheduled_start timestamp with time zone NOT NULL,
    location_address json,
    location_address_point geography(Point,4326),
    venue_id bigint,
    listeners_max bigint,
    is_active boolean DEFAULT true NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
    DROP TABLE public.events_live;
       public         ueoddrluibf30n    false    226    6    6    227    227    6    227    6    227    6    227    6    227    6    227    6    227    6    6            �            1259    17854    media    TABLE     �  CREATE TABLE media (
    media_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    file_store_id character varying(255),
    file_store_key character varying(255),
    is_preview boolean,
    media_format_id uuid,
    "media_MD5" character varying(255),
    secret_url character varying(255),
    title character varying(255),
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
    DROP TABLE public.media;
       public         ueoddrluibf30n    false    226    6    6            �            1259    17844    media_formats    TABLE     �  CREATE TABLE media_formats (
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
       public         ueoddrluibf30n    false    226    6    6            �            1259    17863 
   migrations    TABLE     g   CREATE TABLE migrations (
    migration character varying(255) NOT NULL,
    batch integer NOT NULL
);
    DROP TABLE public.migrations;
       public         ueoddrluibf30n    false    6            �            1259    18515    oauth_access_data    TABLE     �  CREATE TABLE oauth_access_data (
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
       public         ueoddrluibf30n    false    6            �            1259    18507    oauth_auth_data    TABLE     <  CREATE TABLE oauth_auth_data (
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
       public         ueoddrluibf30n    false    6            �            1259    18499    oauth_clients    TABLE     �   CREATE TABLE oauth_clients (
    id character varying(255) NOT NULL,
    secret character varying(255),
    redirecturi character varying(255)
);
 !   DROP TABLE public.oauth_clients;
       public         ueoddrluibf30n    false    6            �            1259    17906    payment_services    TABLE     0  CREATE TABLE payment_services (
    payment_service_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    name character varying(255),
    script_location character varying(255),
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 $   DROP TABLE public.payment_services;
       public         ueoddrluibf30n    false    226    6    6            �            1259    17915    photos    TABLE     �  CREATE TABLE photos (
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
       public         ueoddrluibf30n    false    226    6    6            �            1259    17924    purchase_items    TABLE     P  CREATE TABLE purchase_items (
    purchase_item_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    purchase_id uuid NOT NULL,
    purchase_type character varying(255) NOT NULL,
    unit_price money NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 "   DROP TABLE public.purchase_items;
       public         ueoddrluibf30n    false    226    6    6            �            1259    17930    purchase_purchaseitem    TABLE     1  CREATE TABLE purchase_purchaseitem (
    purcahse_purchaseitem_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    purchase_id uuid NOT NULL,
    purchase_item_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 )   DROP TABLE public.purchase_purchaseitem;
       public         ueoddrluibf30n    false    226    6    6            �            1259    17936 	   purchases    TABLE     �  CREATE TABLE purchases (
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
       public         ueoddrluibf30n    false    226    6    6            �            1259    17945    sessions    TABLE     �   CREATE TABLE sessions (
    id character varying(255) NOT NULL,
    payload text,
    last_activity integer,
    token character varying(255),
    user_id bigint
);
    DROP TABLE public.sessions;
       public         ueoddrluibf30n    false    6            �            1259    24585    slugs    TABLE     e   CREATE TABLE slugs (
    id bigint NOT NULL,
    target json,
    slug character varying NOT NULL
);
    DROP TABLE public.slugs;
       public         ueoddrluibf30n    false    6            �            1259    24583    slugs_id_seq    SEQUENCE     n   CREATE SEQUENCE slugs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.slugs_id_seq;
       public       ueoddrluibf30n    false    222    6            �           0    0    slugs_id_seq    SEQUENCE OWNED BY     /   ALTER SEQUENCE slugs_id_seq OWNED BY slugs.id;
            public       ueoddrluibf30n    false    221            �            1259    17951    tracks    TABLE     �  CREATE TABLE tracks (
    track_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    artist_id bigint,
    checksum_md5 character varying(255),
    downloads bigint DEFAULT (0)::bigint,
    format_keyval character varying(255),
    pic_big character varying(255),
    pic_big_id uuid,
    pic_small character varying(255),
    pic_small_id uuid,
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
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
    DROP TABLE public.tracks;
       public         ueoddrluibf30n    false    226    6    6            �            1259    17967    twitter_account_settings    TABLE     �  CREATE TABLE twitter_account_settings (
    twitter_account_settings_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    handle character varying(255) NOT NULL,
    token1_encrypted character varying(255) NOT NULL,
    token2_encrypted character varying(255) NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 ,   DROP TABLE public.twitter_account_settings;
       public         ueoddrluibf30n    false    226    6    6            �            1259    17976    user_checkins    TABLE     S  CREATE TABLE user_checkins (
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
       public         ueoddrluibf30n    false    226    6    227    227    6    227    6    227    6    227    6    227    6    227    6    227    6    6    6            �            1259    17985    user_media_listens    TABLE     �  CREATE TABLE user_media_listens (
    user_media_listen_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    is_concluded boolean,
    is_live boolean DEFAULT false,
    play_max_position integer,
    play_cursor_position integer,
    media_id uuid NOT NULL,
    user_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 &   DROP TABLE public.user_media_listens;
       public         ueoddrluibf30n    false    226    6    6            �            1259    17992    user_notifications    TABLE     U  CREATE TABLE user_notifications (
    user_notification_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    link json,
    notification json,
    title character varying(255) NOT NULL,
    user_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 &   DROP TABLE public.user_notifications;
       public         ueoddrluibf30n    false    226    6    6            �            1259    18001    user_payments    TABLE     �  CREATE TABLE user_payments (
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
       public         ueoddrluibf30n    false    226    6    6            �            1259    24600 	   usernames    TABLE     m   CREATE TABLE usernames (
    id bigint NOT NULL,
    target json,
    username character varying NOT NULL
);
    DROP TABLE public.usernames;
       public         ueoddrluibf30n    false    6            �            1259    24598    usernames_id_seq    SEQUENCE     r   CREATE SEQUENCE usernames_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.usernames_id_seq;
       public       ueoddrluibf30n    false    6    224            �           0    0    usernames_id_seq    SEQUENCE OWNED BY     7   ALTER SEQUENCE usernames_id_seq OWNED BY usernames.id;
            public       ueoddrluibf30n    false    223            �            1259    18010    users    TABLE     9  CREATE TABLE users (
    user_id bigint NOT NULL,
    birthday_date date,
    current_address json,
    current_address_point geography(Point,4326),
    current_listen_media_id uuid,
    current_location json,
    current_location_point geography(Point,4326),
    contact_phone character varying(30),
    contact_phone_country_code smallint DEFAULT (1)::smallint NOT NULL,
    contact_phone_extension character varying(10),
    email_proxy character varying(255),
    email character varying(255),
    email_pending_update character varying(255),
    facebook_token_encrypted character varying(1000),
    facebook_user character varying(100),
    first_name character varying(255),
    last_name character varying(255),
    interests character varying(255),
    keywords json,
    locale character varying(255),
    pass_hash character varying(64),
    pic character varying(255),
    pic_id uuid,
    pic_big character varying(255),
    pic_big_id uuid,
    pic_cover character varying(255),
    pic_cover_id uuid,
    pic_small character varying(255),
    pic_small_id uuid,
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
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
    DROP TABLE public.users;
       public         ueoddrluibf30n    false    227    227    6    227    6    227    6    227    6    227    6    227    6    227    6    6    6    227    227    6    227    6    227    6    227    6    227    6    227    6    227    6    6            �            1259    18021    users_user_id_seq    SEQUENCE     s   CREATE SEQUENCE users_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.users_user_id_seq;
       public       ueoddrluibf30n    false    6    211            �           0    0    users_user_id_seq    SEQUENCE OWNED BY     9   ALTER SEQUENCE users_user_id_seq OWNED BY users.user_id;
            public       ueoddrluibf30n    false    212            �            1259    18023    venues    TABLE     �  CREATE TABLE venues (
    venue_id bigint NOT NULL,
    address json,
    email character varying(255) NOT NULL,
    facebook_token character varying(1000),
    facebook_user character varying(100),
    address_point geography(Point,4326),
    name character varying(255) NOT NULL,
    phone character varying(30),
    pic character varying(255),
    pic_id uuid,
    pic_big character varying(255),
    pic_big_id uuid,
    pic_cover character varying(255),
    pic_cover_id uuid,
    pic_small character varying(255),
    pic_small_id uuid,
    pic_square character varying(255),
    pic_square_id uuid,
    postal_code character varying(15),
    state character varying(20),
    url character varying(255),
    venue_blurb character varying(5000),
    is_active boolean DEFAULT true NOT NULL,
    is_verified boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
    DROP TABLE public.venues;
       public         ueoddrluibf30n    false    227    227    6    227    6    227    6    227    6    227    6    227    6    227    6    6    6            �            1259    18033    venues_venue_id_seq    SEQUENCE     u   CREATE SEQUENCE venues_venue_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 *   DROP SEQUENCE public.venues_venue_id_seq;
       public       ueoddrluibf30n    false    6    213            �           0    0    venues_venue_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE venues_venue_id_seq OWNED BY venues.venue_id;
            public       ueoddrluibf30n    false    214            �            1259    18035    voucher_purchase    TABLE     !  CREATE TABLE voucher_purchase (
    voucher_purchase_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    voucher_id uuid NOT NULL,
    purchase_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
 $   DROP TABLE public.voucher_purchase;
       public         ueoddrluibf30n    false    226    6    6            �            1259    18041    vouchers    TABLE     �  CREATE TABLE vouchers (
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
       public         ueoddrluibf30n    false    226    6    6            V           2604    18053 	   artist_id    DEFAULT     h   ALTER TABLE ONLY artists ALTER COLUMN artist_id SET DEFAULT nextval('artists_artist_id_seq'::regclass);
 @   ALTER TABLE public.artists ALTER COLUMN artist_id DROP DEFAULT;
       public       ueoddrluibf30n    false    189    188            �           2604    24588    id    DEFAULT     V   ALTER TABLE ONLY slugs ALTER COLUMN id SET DEFAULT nextval('slugs_id_seq'::regclass);
 7   ALTER TABLE public.slugs ALTER COLUMN id DROP DEFAULT;
       public       ueoddrluibf30n    false    222    221    222            �           2604    24603    id    DEFAULT     ^   ALTER TABLE ONLY usernames ALTER COLUMN id SET DEFAULT nextval('usernames_id_seq'::regclass);
 ;   ALTER TABLE public.usernames ALTER COLUMN id DROP DEFAULT;
       public       ueoddrluibf30n    false    223    224    224            �           2604    18054    user_id    DEFAULT     `   ALTER TABLE ONLY users ALTER COLUMN user_id SET DEFAULT nextval('users_user_id_seq'::regclass);
 <   ALTER TABLE public.users ALTER COLUMN user_id DROP DEFAULT;
       public       ueoddrluibf30n    false    212    211            �           2604    18055    venue_id    DEFAULT     d   ALTER TABLE ONLY venues ALTER COLUMN venue_id SET DEFAULT nextval('venues_venue_id_seq'::regclass);
 >   ALTER TABLE public.venues ALTER COLUMN venue_id DROP DEFAULT;
       public       ueoddrluibf30n    false    214    213            �           2606    18059    album_track_pkey 
   CONSTRAINT     _   ALTER TABLE ONLY album_track
    ADD CONSTRAINT album_track_pkey PRIMARY KEY (album_track_id);
 F   ALTER TABLE ONLY public.album_track DROP CONSTRAINT album_track_pkey;
       public         ueoddrluibf30n    false    183    183            �           2606    18061    albums_pkey 
   CONSTRAINT     O   ALTER TABLE ONLY albums
    ADD CONSTRAINT albums_pkey PRIMARY KEY (album_id);
 <   ALTER TABLE ONLY public.albums DROP CONSTRAINT albums_pkey;
       public         ueoddrluibf30n    false    184    184            �           2606    18063    artist_checkins_pkey 
   CONSTRAINT     j   ALTER TABLE ONLY artist_checkins
    ADD CONSTRAINT artist_checkins_pkey PRIMARY KEY (artist_checkin_id);
 N   ALTER TABLE ONLY public.artist_checkins DROP CONSTRAINT artist_checkins_pkey;
       public         ueoddrluibf30n    false    185    185            �           2606    18065    artist_notifications_pkey 
   CONSTRAINT     y   ALTER TABLE ONLY artist_notifications
    ADD CONSTRAINT artist_notifications_pkey PRIMARY KEY (artist_notification_id);
 X   ALTER TABLE ONLY public.artist_notifications DROP CONSTRAINT artist_notifications_pkey;
       public         ueoddrluibf30n    false    186    186            �           2606    18067    artist_user_pkey 
   CONSTRAINT     _   ALTER TABLE ONLY artist_user
    ADD CONSTRAINT artist_user_pkey PRIMARY KEY (artist_user_id);
 F   ALTER TABLE ONLY public.artist_user DROP CONSTRAINT artist_user_pkey;
       public         ueoddrluibf30n    false    187    187            �           2606    18069    artists_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY artists
    ADD CONSTRAINT artists_pkey PRIMARY KEY (artist_id);
 >   ALTER TABLE ONLY public.artists DROP CONSTRAINT artists_pkey;
       public         ueoddrluibf30n    false    188    188            �           2606    18075    event_admin_log_pkey 
   CONSTRAINT     k   ALTER TABLE ONLY event_admin_log
    ADD CONSTRAINT event_admin_log_pkey PRIMARY KEY (event_admin_log_id);
 N   ALTER TABLE ONLY public.event_admin_log DROP CONSTRAINT event_admin_log_pkey;
       public         ueoddrluibf30n    false    190    190            �           2606    18077    event_live_notifications_pkey 
   CONSTRAINT     �   ALTER TABLE ONLY event_live_notifications
    ADD CONSTRAINT event_live_notifications_pkey PRIMARY KEY (event_live_notification_id);
 `   ALTER TABLE ONLY public.event_live_notifications DROP CONSTRAINT event_live_notifications_pkey;
       public         ueoddrluibf30n    false    191    191            �           2606    18079 >   event_live_provisioned_broadcasts_broadcast_stream_name_unique 
   CONSTRAINT     �   ALTER TABLE ONLY event_live_provisioned_broadcasts
    ADD CONSTRAINT event_live_provisioned_broadcasts_broadcast_stream_name_unique UNIQUE (broadcast_stream_name);
 �   ALTER TABLE ONLY public.event_live_provisioned_broadcasts DROP CONSTRAINT event_live_provisioned_broadcasts_broadcast_stream_name_unique;
       public         ueoddrluibf30n    false    192    192            �           2606    18081 &   event_live_provisioned_broadcasts_pkey 
   CONSTRAINT     �   ALTER TABLE ONLY event_live_provisioned_broadcasts
    ADD CONSTRAINT event_live_provisioned_broadcasts_pkey PRIMARY KEY (event_live_provisioned_broadcast_id);
 r   ALTER TABLE ONLY public.event_live_provisioned_broadcasts DROP CONSTRAINT event_live_provisioned_broadcasts_pkey;
       public         ueoddrluibf30n    false    192    192            �           2606    18083 ;   event_live_provisioned_broadcasts_provider_stream_id_unique 
   CONSTRAINT     �   ALTER TABLE ONLY event_live_provisioned_broadcasts
    ADD CONSTRAINT event_live_provisioned_broadcasts_provider_stream_id_unique UNIQUE (provider_stream_id);
 �   ALTER TABLE ONLY public.event_live_provisioned_broadcasts DROP CONSTRAINT event_live_provisioned_broadcasts_provider_stream_id_unique;
       public         ueoddrluibf30n    false    192    192            �           2606    18085    event_live_routes_pkey 
   CONSTRAINT     p   ALTER TABLE ONLY event_live_routes
    ADD CONSTRAINT event_live_routes_pkey PRIMARY KEY (event_live_route_id);
 R   ALTER TABLE ONLY public.event_live_routes DROP CONSTRAINT event_live_routes_pkey;
       public         ueoddrluibf30n    false    193    193            �           2606    18087    event_live_segments_pkey 
   CONSTRAINT     v   ALTER TABLE ONLY event_live_segments
    ADD CONSTRAINT event_live_segments_pkey PRIMARY KEY (event_live_segment_id);
 V   ALTER TABLE ONLY public.event_live_segments DROP CONSTRAINT event_live_segments_pkey;
       public         ueoddrluibf30n    false    194    194            �           2606    18089    events_live_pkey 
   CONSTRAINT     ^   ALTER TABLE ONLY events_live
    ADD CONSTRAINT events_live_pkey PRIMARY KEY (event_live_id);
 F   ALTER TABLE ONLY public.events_live DROP CONSTRAINT events_live_pkey;
       public         ueoddrluibf30n    false    195    195            �           2606    18091    media_formats_pkey 
   CONSTRAINT     d   ALTER TABLE ONLY media_formats
    ADD CONSTRAINT media_formats_pkey PRIMARY KEY (media_format_id);
 J   ALTER TABLE ONLY public.media_formats DROP CONSTRAINT media_formats_pkey;
       public         ueoddrluibf30n    false    196    196            �           2606    18093    medias_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY media
    ADD CONSTRAINT medias_pkey PRIMARY KEY (media_id);
 ;   ALTER TABLE ONLY public.media DROP CONSTRAINT medias_pkey;
       public         ueoddrluibf30n    false    197    197                       2606    18522    oauth_access_data_pkey 
   CONSTRAINT     h   ALTER TABLE ONLY oauth_access_data
    ADD CONSTRAINT oauth_access_data_pkey PRIMARY KEY (accesstoken);
 R   ALTER TABLE ONLY public.oauth_access_data DROP CONSTRAINT oauth_access_data_pkey;
       public         ueoddrluibf30n    false    219    219            �           2606    18514    oauth_auth_data_pkey 
   CONSTRAINT     ]   ALTER TABLE ONLY oauth_auth_data
    ADD CONSTRAINT oauth_auth_data_pkey PRIMARY KEY (code);
 N   ALTER TABLE ONLY public.oauth_auth_data DROP CONSTRAINT oauth_auth_data_pkey;
       public         ueoddrluibf30n    false    218    218            �           2606    18506    oauth_clients_pkey 
   CONSTRAINT     W   ALTER TABLE ONLY oauth_clients
    ADD CONSTRAINT oauth_clients_pkey PRIMARY KEY (id);
 J   ALTER TABLE ONLY public.oauth_clients DROP CONSTRAINT oauth_clients_pkey;
       public         ueoddrluibf30n    false    217    217            �           2606    18095    payment_services_pkey 
   CONSTRAINT     m   ALTER TABLE ONLY payment_services
    ADD CONSTRAINT payment_services_pkey PRIMARY KEY (payment_service_id);
 P   ALTER TABLE ONLY public.payment_services DROP CONSTRAINT payment_services_pkey;
       public         ueoddrluibf30n    false    199    199            �           2606    18097    photos_pkey 
   CONSTRAINT     O   ALTER TABLE ONLY photos
    ADD CONSTRAINT photos_pkey PRIMARY KEY (photo_id);
 <   ALTER TABLE ONLY public.photos DROP CONSTRAINT photos_pkey;
       public         ueoddrluibf30n    false    200    200            �           2606    18099    purchase_items_pkey 
   CONSTRAINT     g   ALTER TABLE ONLY purchase_items
    ADD CONSTRAINT purchase_items_pkey PRIMARY KEY (purchase_item_id);
 L   ALTER TABLE ONLY public.purchase_items DROP CONSTRAINT purchase_items_pkey;
       public         ueoddrluibf30n    false    201    201            �           2606    18101    purchase_purchaseitem_pkey 
   CONSTRAINT     }   ALTER TABLE ONLY purchase_purchaseitem
    ADD CONSTRAINT purchase_purchaseitem_pkey PRIMARY KEY (purcahse_purchaseitem_id);
 Z   ALTER TABLE ONLY public.purchase_purchaseitem DROP CONSTRAINT purchase_purchaseitem_pkey;
       public         ueoddrluibf30n    false    202    202            �           2606    18103    purchases_pkey 
   CONSTRAINT     X   ALTER TABLE ONLY purchases
    ADD CONSTRAINT purchases_pkey PRIMARY KEY (purchase_id);
 B   ALTER TABLE ONLY public.purchases DROP CONSTRAINT purchases_pkey;
       public         ueoddrluibf30n    false    203    203            �           2606    18107    sessions_id_unique 
   CONSTRAINT     M   ALTER TABLE ONLY sessions
    ADD CONSTRAINT sessions_id_unique UNIQUE (id);
 E   ALTER TABLE ONLY public.sessions DROP CONSTRAINT sessions_id_unique;
       public         ueoddrluibf30n    false    204    204                       2606    24593 
   slugs_pkey 
   CONSTRAINT     G   ALTER TABLE ONLY slugs
    ADD CONSTRAINT slugs_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.slugs DROP CONSTRAINT slugs_pkey;
       public         ueoddrluibf30n    false    222    222                       2606    24626    slugs_slug_key1 
   CONSTRAINT     I   ALTER TABLE ONLY slugs
    ADD CONSTRAINT slugs_slug_key1 UNIQUE (slug);
 ?   ALTER TABLE ONLY public.slugs DROP CONSTRAINT slugs_slug_key1;
       public         ueoddrluibf30n    false    222    222            �           2606    18109    tracks_pkey 
   CONSTRAINT     O   ALTER TABLE ONLY tracks
    ADD CONSTRAINT tracks_pkey PRIMARY KEY (track_id);
 <   ALTER TABLE ONLY public.tracks DROP CONSTRAINT tracks_pkey;
       public         ueoddrluibf30n    false    205    205            �           2606    18111    twitter_account_settings_pkey 
   CONSTRAINT     �   ALTER TABLE ONLY twitter_account_settings
    ADD CONSTRAINT twitter_account_settings_pkey PRIMARY KEY (twitter_account_settings_id);
 `   ALTER TABLE ONLY public.twitter_account_settings DROP CONSTRAINT twitter_account_settings_pkey;
       public         ueoddrluibf30n    false    206    206            �           2606    18113    user_checkins_pkey 
   CONSTRAINT     d   ALTER TABLE ONLY user_checkins
    ADD CONSTRAINT user_checkins_pkey PRIMARY KEY (user_checkin_id);
 J   ALTER TABLE ONLY public.user_checkins DROP CONSTRAINT user_checkins_pkey;
       public         ueoddrluibf30n    false    207    207            �           2606    18115    user_media_listens_pkey 
   CONSTRAINT     s   ALTER TABLE ONLY user_media_listens
    ADD CONSTRAINT user_media_listens_pkey PRIMARY KEY (user_media_listen_id);
 T   ALTER TABLE ONLY public.user_media_listens DROP CONSTRAINT user_media_listens_pkey;
       public         ueoddrluibf30n    false    208    208            �           2606    18117    user_notifications_pkey 
   CONSTRAINT     s   ALTER TABLE ONLY user_notifications
    ADD CONSTRAINT user_notifications_pkey PRIMARY KEY (user_notification_id);
 T   ALTER TABLE ONLY public.user_notifications DROP CONSTRAINT user_notifications_pkey;
       public         ueoddrluibf30n    false    209    209            �           2606    18119    user_payments_pkey 
   CONSTRAINT     d   ALTER TABLE ONLY user_payments
    ADD CONSTRAINT user_payments_pkey PRIMARY KEY (user_payment_id);
 J   ALTER TABLE ONLY public.user_payments DROP CONSTRAINT user_payments_pkey;
       public         ueoddrluibf30n    false    210    210            
           2606    24608    usernames_pkey 
   CONSTRAINT     O   ALTER TABLE ONLY usernames
    ADD CONSTRAINT usernames_pkey PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.usernames DROP CONSTRAINT usernames_pkey;
       public         ueoddrluibf30n    false    224    224                       2606    24622    usernames_username_key 
   CONSTRAINT     X   ALTER TABLE ONLY usernames
    ADD CONSTRAINT usernames_username_key UNIQUE (username);
 J   ALTER TABLE ONLY public.usernames DROP CONSTRAINT usernames_username_key;
       public         ueoddrluibf30n    false    224    224            �           2606    18123    users_email_unique 
   CONSTRAINT     M   ALTER TABLE ONLY users
    ADD CONSTRAINT users_email_unique UNIQUE (email);
 B   ALTER TABLE ONLY public.users DROP CONSTRAINT users_email_unique;
       public         ueoddrluibf30n    false    211    211            �           2606    18125 
   users_pkey 
   CONSTRAINT     L   ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public         ueoddrluibf30n    false    211    211            �           2606    18127    venues_pkey 
   CONSTRAINT     O   ALTER TABLE ONLY venues
    ADD CONSTRAINT venues_pkey PRIMARY KEY (venue_id);
 <   ALTER TABLE ONLY public.venues DROP CONSTRAINT venues_pkey;
       public         ueoddrluibf30n    false    213    213            �           2606    18129    voucher_purchase_pkey 
   CONSTRAINT     n   ALTER TABLE ONLY voucher_purchase
    ADD CONSTRAINT voucher_purchase_pkey PRIMARY KEY (voucher_purchase_id);
 P   ALTER TABLE ONLY public.voucher_purchase DROP CONSTRAINT voucher_purchase_pkey;
       public         ueoddrluibf30n    false    215    215            �           2606    18131    vouchers_pkey 
   CONSTRAINT     U   ALTER TABLE ONLY vouchers
    ADD CONSTRAINT vouchers_pkey PRIMARY KEY (voucher_id);
 @   ALTER TABLE ONLY public.vouchers DROP CONSTRAINT vouchers_pkey;
       public         ueoddrluibf30n    false    216    216                       1259    24633    slugs_case_insensitive_slug_key    INDEX     `   CREATE UNIQUE INDEX slugs_case_insensitive_slug_key ON slugs USING btree (lower((slug)::text));
 3   DROP INDEX public.slugs_case_insensitive_slug_key;
       public         ueoddrluibf30n    false    222    222                       1259    24597    slugs_target_artist_id    INDEX     h   CREATE INDEX slugs_target_artist_id ON slugs USING btree ((((target ->> 'artist_id'::text))::integer));
 *   DROP INDEX public.slugs_target_artist_id;
       public         ueoddrluibf30n    false    222    222                       1259    24596    slugs_target_user_id    INDEX     d   CREATE INDEX slugs_target_user_id ON slugs USING btree ((((target ->> 'user_id'::text))::integer));
 (   DROP INDEX public.slugs_target_user_id;
       public         ueoddrluibf30n    false    222    222                       1259    24612    usernames_target_artist_id    INDEX     p   CREATE INDEX usernames_target_artist_id ON usernames USING btree ((((target ->> 'artist_id'::text))::integer));
 .   DROP INDEX public.usernames_target_artist_id;
       public         ueoddrluibf30n    false    224    224                       1259    24611    usernames_target_user_id    INDEX     l   CREATE INDEX usernames_target_user_id ON usernames USING btree ((((target ->> 'user_id'::text))::integer));
 ,   DROP INDEX public.usernames_target_user_id;
       public         ueoddrluibf30n    false    224    224            V           2620    24618    artists_username_insert    TRIGGER     z   CREATE TRIGGER artists_username_insert BEFORE INSERT ON artists FOR EACH ROW EXECUTE PROCEDURE artists_username_insert();
 8   DROP TRIGGER artists_username_insert ON public.artists;
       public       ueoddrluibf30n    false    188    1305            W           2620    24620    artists_username_update    TRIGGER     z   CREATE TRIGGER artists_username_update BEFORE UPDATE ON artists FOR EACH ROW EXECUTE PROCEDURE artists_username_update();
 8   DROP TRIGGER artists_username_update ON public.artists;
       public       ueoddrluibf30n    false    1304    188            X           2620    24614    users_username_insert    TRIGGER     t   CREATE TRIGGER users_username_insert BEFORE INSERT ON users FOR EACH ROW EXECUTE PROCEDURE users_username_insert();
 4   DROP TRIGGER users_username_insert ON public.users;
       public       ueoddrluibf30n    false    1288    211            Y           2620    24616    users_username_update    TRIGGER     t   CREATE TRIGGER users_username_update BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE users_username_update();
 4   DROP TRIGGER users_username_update ON public.users;
       public       ueoddrluibf30n    false    211    1279                       2606    18132    album_track_album_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY album_track
    ADD CONSTRAINT album_track_album_id_foreign FOREIGN KEY (album_id) REFERENCES albums(album_id);
 R   ALTER TABLE ONLY public.album_track DROP CONSTRAINT album_track_album_id_foreign;
       public       ueoddrluibf30n    false    183    4285    184                       2606    18137    album_track_track_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY album_track
    ADD CONSTRAINT album_track_track_id_foreign FOREIGN KEY (track_id) REFERENCES tracks(track_id);
 R   ALTER TABLE ONLY public.album_track DROP CONSTRAINT album_track_track_id_foreign;
       public       ueoddrluibf30n    false    183    205    4327                       2606    18142    albums_artist_id_foreign    FK CONSTRAINT     {   ALTER TABLE ONLY albums
    ADD CONSTRAINT albums_artist_id_foreign FOREIGN KEY (artist_id) REFERENCES artists(artist_id);
 I   ALTER TABLE ONLY public.albums DROP CONSTRAINT albums_artist_id_foreign;
       public       ueoddrluibf30n    false    184    4293    188                       2606    18147    albums_pic_big_id_foreign    FK CONSTRAINT     {   ALTER TABLE ONLY albums
    ADD CONSTRAINT albums_pic_big_id_foreign FOREIGN KEY (pic_big_id) REFERENCES photos(photo_id);
 J   ALTER TABLE ONLY public.albums DROP CONSTRAINT albums_pic_big_id_foreign;
       public       ueoddrluibf30n    false    184    4317    200                       2606    18152    albums_pic_small_id_foreign    FK CONSTRAINT        ALTER TABLE ONLY albums
    ADD CONSTRAINT albums_pic_small_id_foreign FOREIGN KEY (pic_small_id) REFERENCES photos(photo_id);
 L   ALTER TABLE ONLY public.albums DROP CONSTRAINT albums_pic_small_id_foreign;
       public       ueoddrluibf30n    false    4317    200    184                       2606    18157    albums_pic_square_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY albums
    ADD CONSTRAINT albums_pic_square_id_foreign FOREIGN KEY (pic_square_id) REFERENCES photos(photo_id);
 M   ALTER TABLE ONLY public.albums DROP CONSTRAINT albums_pic_square_id_foreign;
       public       ueoddrluibf30n    false    200    4317    184                       2606    18162 !   artist_checkins_artist_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artist_checkins
    ADD CONSTRAINT artist_checkins_artist_id_foreign FOREIGN KEY (artist_id) REFERENCES artists(artist_id);
 [   ALTER TABLE ONLY public.artist_checkins DROP CONSTRAINT artist_checkins_artist_id_foreign;
       public       ueoddrluibf30n    false    185    4293    188                       2606    18167 7   artist_checkins_checkin_location_venue_match_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artist_checkins
    ADD CONSTRAINT artist_checkins_checkin_location_venue_match_id_foreign FOREIGN KEY (checkin_location_venue_match_id) REFERENCES venues(venue_id);
 q   ALTER TABLE ONLY public.artist_checkins DROP CONSTRAINT artist_checkins_checkin_location_venue_match_id_foreign;
       public       ueoddrluibf30n    false    185    4343    213                       2606    18172 "   artist_checkins_pic_big_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artist_checkins
    ADD CONSTRAINT artist_checkins_pic_big_id_foreign FOREIGN KEY (pic_big_id) REFERENCES photos(photo_id);
 \   ALTER TABLE ONLY public.artist_checkins DROP CONSTRAINT artist_checkins_pic_big_id_foreign;
       public       ueoddrluibf30n    false    185    4317    200                       2606    18177 $   artist_checkins_pic_small_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artist_checkins
    ADD CONSTRAINT artist_checkins_pic_small_id_foreign FOREIGN KEY (pic_small_id) REFERENCES photos(photo_id);
 ^   ALTER TABLE ONLY public.artist_checkins DROP CONSTRAINT artist_checkins_pic_small_id_foreign;
       public       ueoddrluibf30n    false    185    4317    200                       2606    18182 %   artist_checkins_pic_square_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artist_checkins
    ADD CONSTRAINT artist_checkins_pic_square_id_foreign FOREIGN KEY (pic_square_id) REFERENCES photos(photo_id);
 _   ALTER TABLE ONLY public.artist_checkins DROP CONSTRAINT artist_checkins_pic_square_id_foreign;
       public       ueoddrluibf30n    false    185    4317    200                       2606    18187 &   artist_notifications_artist_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artist_notifications
    ADD CONSTRAINT artist_notifications_artist_id_foreign FOREIGN KEY (artist_id) REFERENCES artists(artist_id);
 e   ALTER TABLE ONLY public.artist_notifications DROP CONSTRAINT artist_notifications_artist_id_foreign;
       public       ueoddrluibf30n    false    186    4293    188                       2606    18192    artist_user_artist_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artist_user
    ADD CONSTRAINT artist_user_artist_id_foreign FOREIGN KEY (artist_id) REFERENCES artists(artist_id);
 S   ALTER TABLE ONLY public.artist_user DROP CONSTRAINT artist_user_artist_id_foreign;
       public       ueoddrluibf30n    false    4293    188    187                       2606    18197    artist_user_user_id_foreign    FK CONSTRAINT     }   ALTER TABLE ONLY artist_user
    ADD CONSTRAINT artist_user_user_id_foreign FOREIGN KEY (user_id) REFERENCES users(user_id);
 Q   ALTER TABLE ONLY public.artist_user DROP CONSTRAINT artist_user_user_id_foreign;
       public       ueoddrluibf30n    false    211    4341    187                       2606    18202    artists_booking_user_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artists
    ADD CONSTRAINT artists_booking_user_id_foreign FOREIGN KEY (booking_user_id) REFERENCES users(user_id);
 Q   ALTER TABLE ONLY public.artists DROP CONSTRAINT artists_booking_user_id_foreign;
       public       ueoddrluibf30n    false    211    4341    188                       2606    18207    artists_manager_user_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artists
    ADD CONSTRAINT artists_manager_user_id_foreign FOREIGN KEY (manager_user_id) REFERENCES users(user_id);
 Q   ALTER TABLE ONLY public.artists DROP CONSTRAINT artists_manager_user_id_foreign;
       public       ueoddrluibf30n    false    211    4341    188                       2606    18212    artists_pic_big_id_foreign    FK CONSTRAINT     }   ALTER TABLE ONLY artists
    ADD CONSTRAINT artists_pic_big_id_foreign FOREIGN KEY (pic_big_id) REFERENCES photos(photo_id);
 L   ALTER TABLE ONLY public.artists DROP CONSTRAINT artists_pic_big_id_foreign;
       public       ueoddrluibf30n    false    188    4317    200            $           2606    24654    artists_pic_cover_id_fkey    FK CONSTRAINT     ~   ALTER TABLE ONLY artists
    ADD CONSTRAINT artists_pic_cover_id_fkey FOREIGN KEY (pic_cover_id) REFERENCES photos(photo_id);
 K   ALTER TABLE ONLY public.artists DROP CONSTRAINT artists_pic_cover_id_fkey;
       public       ueoddrluibf30n    false    200    4317    188                        2606    18217    artists_pic_id_foreign    FK CONSTRAINT     u   ALTER TABLE ONLY artists
    ADD CONSTRAINT artists_pic_id_foreign FOREIGN KEY (pic_id) REFERENCES photos(photo_id);
 H   ALTER TABLE ONLY public.artists DROP CONSTRAINT artists_pic_id_foreign;
       public       ueoddrluibf30n    false    4317    200    188            !           2606    18222    artists_pic_small_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artists
    ADD CONSTRAINT artists_pic_small_id_foreign FOREIGN KEY (pic_small_id) REFERENCES photos(photo_id);
 N   ALTER TABLE ONLY public.artists DROP CONSTRAINT artists_pic_small_id_foreign;
       public       ueoddrluibf30n    false    188    200    4317            "           2606    18227    artists_pic_square_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artists
    ADD CONSTRAINT artists_pic_square_id_foreign FOREIGN KEY (pic_square_id) REFERENCES photos(photo_id);
 O   ALTER TABLE ONLY public.artists DROP CONSTRAINT artists_pic_square_id_foreign;
       public       ueoddrluibf30n    false    188    200    4317            #           2606    18232 +   artists_twitter_account_settings_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY artists
    ADD CONSTRAINT artists_twitter_account_settings_id_foreign FOREIGN KEY (twitter_account_settings_id) REFERENCES twitter_account_settings(twitter_account_settings_id);
 ]   ALTER TABLE ONLY public.artists DROP CONSTRAINT artists_twitter_account_settings_id_foreign;
       public       ueoddrluibf30n    false    206    188    4329            %           2606    18237 %   event_admin_log_event_live_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY event_admin_log
    ADD CONSTRAINT event_admin_log_event_live_id_foreign FOREIGN KEY (event_live_id) REFERENCES events_live(event_live_id);
 _   ALTER TABLE ONLY public.event_admin_log DROP CONSTRAINT event_admin_log_event_live_id_foreign;
       public       ueoddrluibf30n    false    190    195    4309            &           2606    18242 .   event_live_notifications_event_live_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY event_live_notifications
    ADD CONSTRAINT event_live_notifications_event_live_id_foreign FOREIGN KEY (event_live_id) REFERENCES events_live(event_live_id);
 q   ALTER TABLE ONLY public.event_live_notifications DROP CONSTRAINT event_live_notifications_event_live_id_foreign;
       public       ueoddrluibf30n    false    191    195    4309            '           2606    18247 '   event_live_routes_event_live_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY event_live_routes
    ADD CONSTRAINT event_live_routes_event_live_id_foreign FOREIGN KEY (event_live_id) REFERENCES events_live(event_live_id);
 c   ALTER TABLE ONLY public.event_live_routes DROP CONSTRAINT event_live_routes_event_live_id_foreign;
       public       ueoddrluibf30n    false    193    195    4309            (           2606    18252 =   event_live_routes_event_live_provisioned_broadcast_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY event_live_routes
    ADD CONSTRAINT event_live_routes_event_live_provisioned_broadcast_id_foreign FOREIGN KEY (event_live_provisioned_broadcast_id) REFERENCES event_live_provisioned_broadcasts(event_live_provisioned_broadcast_id);
 y   ALTER TABLE ONLY public.event_live_routes DROP CONSTRAINT event_live_routes_event_live_provisioned_broadcast_id_foreign;
       public       ueoddrluibf30n    false    4301    192    193            )           2606    18257 )   event_live_segments_event_live_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY event_live_segments
    ADD CONSTRAINT event_live_segments_event_live_id_foreign FOREIGN KEY (event_live_id) REFERENCES events_live(event_live_id);
 g   ALTER TABLE ONLY public.event_live_segments DROP CONSTRAINT event_live_segments_event_live_id_foreign;
       public       ueoddrluibf30n    false    194    4309    195            *           2606    18262 $   events_live_archive_track_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY events_live
    ADD CONSTRAINT events_live_archive_track_id_foreign FOREIGN KEY (archive_track_id) REFERENCES tracks(track_id);
 Z   ALTER TABLE ONLY public.events_live DROP CONSTRAINT events_live_archive_track_id_foreign;
       public       ueoddrluibf30n    false    195    205    4327            +           2606    18267    events_live_artist_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY events_live
    ADD CONSTRAINT events_live_artist_id_foreign FOREIGN KEY (artist_id) REFERENCES artists(artist_id);
 S   ALTER TABLE ONLY public.events_live DROP CONSTRAINT events_live_artist_id_foreign;
       public       ueoddrluibf30n    false    195    188    4293            ,           2606    18272    events_live_pic_big_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY events_live
    ADD CONSTRAINT events_live_pic_big_id_foreign FOREIGN KEY (pic_big_id) REFERENCES photos(photo_id);
 T   ALTER TABLE ONLY public.events_live DROP CONSTRAINT events_live_pic_big_id_foreign;
       public       ueoddrluibf30n    false    200    195    4317            -           2606    18277     events_live_pic_small_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY events_live
    ADD CONSTRAINT events_live_pic_small_id_foreign FOREIGN KEY (pic_small_id) REFERENCES photos(photo_id);
 V   ALTER TABLE ONLY public.events_live DROP CONSTRAINT events_live_pic_small_id_foreign;
       public       ueoddrluibf30n    false    195    200    4317            .           2606    18282    events_live_venue_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY events_live
    ADD CONSTRAINT events_live_venue_id_foreign FOREIGN KEY (venue_id) REFERENCES venues(venue_id);
 R   ALTER TABLE ONLY public.events_live DROP CONSTRAINT events_live_venue_id_foreign;
       public       ueoddrluibf30n    false    195    213    4343            /           2606    18287    media_formats_pic_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY media_formats
    ADD CONSTRAINT media_formats_pic_id_foreign FOREIGN KEY (pic_id) REFERENCES photos(photo_id);
 T   ALTER TABLE ONLY public.media_formats DROP CONSTRAINT media_formats_pic_id_foreign;
       public       ueoddrluibf30n    false    196    200    4317            0           2606    18292    medias_media_format_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY media
    ADD CONSTRAINT medias_media_format_id_foreign FOREIGN KEY (media_format_id) REFERENCES media_formats(media_format_id);
 N   ALTER TABLE ONLY public.media DROP CONSTRAINT medias_media_format_id_foreign;
       public       ueoddrluibf30n    false    196    197    4311            1           2606    18297 "   purchase_items_purchase_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY purchase_items
    ADD CONSTRAINT purchase_items_purchase_id_foreign FOREIGN KEY (purchase_id) REFERENCES purchases(purchase_id);
 [   ALTER TABLE ONLY public.purchase_items DROP CONSTRAINT purchase_items_purchase_id_foreign;
       public       ueoddrluibf30n    false    4323    201    203            2           2606    18302 )   purchase_purchaseitem_purchase_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY purchase_purchaseitem
    ADD CONSTRAINT purchase_purchaseitem_purchase_id_foreign FOREIGN KEY (purchase_id) REFERENCES purchases(purchase_id);
 i   ALTER TABLE ONLY public.purchase_purchaseitem DROP CONSTRAINT purchase_purchaseitem_purchase_id_foreign;
       public       ueoddrluibf30n    false    4323    203    202            3           2606    18307 .   purchase_purchaseitem_purchase_item_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY purchase_purchaseitem
    ADD CONSTRAINT purchase_purchaseitem_purchase_item_id_foreign FOREIGN KEY (purchase_item_id) REFERENCES purchase_items(purchase_item_id);
 n   ALTER TABLE ONLY public.purchase_purchaseitem DROP CONSTRAINT purchase_purchaseitem_purchase_item_id_foreign;
       public       ueoddrluibf30n    false    4319    202    201            4           2606    18312    purchases_user_id_foreign    FK CONSTRAINT     y   ALTER TABLE ONLY purchases
    ADD CONSTRAINT purchases_user_id_foreign FOREIGN KEY (user_id) REFERENCES users(user_id);
 M   ALTER TABLE ONLY public.purchases DROP CONSTRAINT purchases_user_id_foreign;
       public       ueoddrluibf30n    false    203    211    4341            5           2606    18317 !   purchases_user_payment_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY purchases
    ADD CONSTRAINT purchases_user_payment_id_foreign FOREIGN KEY (user_payment_id) REFERENCES user_payments(user_payment_id);
 U   ALTER TABLE ONLY public.purchases DROP CONSTRAINT purchases_user_payment_id_foreign;
       public       ueoddrluibf30n    false    210    4337    203            6           2606    18322    sessions_user_id_foreign    FK CONSTRAINT     w   ALTER TABLE ONLY sessions
    ADD CONSTRAINT sessions_user_id_foreign FOREIGN KEY (user_id) REFERENCES users(user_id);
 K   ALTER TABLE ONLY public.sessions DROP CONSTRAINT sessions_user_id_foreign;
       public       ueoddrluibf30n    false    211    4341    204            7           2606    18327    tracks_artist_id_foreign    FK CONSTRAINT     {   ALTER TABLE ONLY tracks
    ADD CONSTRAINT tracks_artist_id_foreign FOREIGN KEY (artist_id) REFERENCES artists(artist_id);
 I   ALTER TABLE ONLY public.tracks DROP CONSTRAINT tracks_artist_id_foreign;
       public       ueoddrluibf30n    false    4293    188    205            8           2606    18332    tracks_media_id_foreign    FK CONSTRAINT     v   ALTER TABLE ONLY tracks
    ADD CONSTRAINT tracks_media_id_foreign FOREIGN KEY (media_id) REFERENCES media(media_id);
 H   ALTER TABLE ONLY public.tracks DROP CONSTRAINT tracks_media_id_foreign;
       public       ueoddrluibf30n    false    205    197    4313            9           2606    18337 #   tracks_must_bundle_album_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY tracks
    ADD CONSTRAINT tracks_must_bundle_album_id_foreign FOREIGN KEY (must_bundle_album_id) REFERENCES albums(album_id);
 T   ALTER TABLE ONLY public.tracks DROP CONSTRAINT tracks_must_bundle_album_id_foreign;
       public       ueoddrluibf30n    false    184    4285    205            :           2606    18342    tracks_pic_big_id_foreign    FK CONSTRAINT     {   ALTER TABLE ONLY tracks
    ADD CONSTRAINT tracks_pic_big_id_foreign FOREIGN KEY (pic_big_id) REFERENCES photos(photo_id);
 J   ALTER TABLE ONLY public.tracks DROP CONSTRAINT tracks_pic_big_id_foreign;
       public       ueoddrluibf30n    false    200    4317    205            ;           2606    18347    tracks_pic_small_id_foreign    FK CONSTRAINT        ALTER TABLE ONLY tracks
    ADD CONSTRAINT tracks_pic_small_id_foreign FOREIGN KEY (pic_small_id) REFERENCES photos(photo_id);
 L   ALTER TABLE ONLY public.tracks DROP CONSTRAINT tracks_pic_small_id_foreign;
       public       ueoddrluibf30n    false    205    200    4317            <           2606    18352    tracks_pic_square_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY tracks
    ADD CONSTRAINT tracks_pic_square_id_foreign FOREIGN KEY (pic_square_id) REFERENCES photos(photo_id);
 M   ALTER TABLE ONLY public.tracks DROP CONSTRAINT tracks_pic_square_id_foreign;
       public       ueoddrluibf30n    false    4317    200    205            =           2606    18357    tracks_preview_media_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY tracks
    ADD CONSTRAINT tracks_preview_media_id_foreign FOREIGN KEY (preview_media_id) REFERENCES media(media_id);
 P   ALTER TABLE ONLY public.tracks DROP CONSTRAINT tracks_preview_media_id_foreign;
       public       ueoddrluibf30n    false    197    205    4313            >           2606    18362 5   user_checkins_checkin_location_venue_match_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY user_checkins
    ADD CONSTRAINT user_checkins_checkin_location_venue_match_id_foreign FOREIGN KEY (checkin_location_venue_match_id) REFERENCES venues(venue_id);
 m   ALTER TABLE ONLY public.user_checkins DROP CONSTRAINT user_checkins_checkin_location_venue_match_id_foreign;
       public       ueoddrluibf30n    false    4343    213    207            ?           2606    18367     user_checkins_pic_big_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY user_checkins
    ADD CONSTRAINT user_checkins_pic_big_id_foreign FOREIGN KEY (pic_big_id) REFERENCES photos(photo_id);
 X   ALTER TABLE ONLY public.user_checkins DROP CONSTRAINT user_checkins_pic_big_id_foreign;
       public       ueoddrluibf30n    false    200    4317    207            @           2606    18372 "   user_checkins_pic_small_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY user_checkins
    ADD CONSTRAINT user_checkins_pic_small_id_foreign FOREIGN KEY (pic_small_id) REFERENCES photos(photo_id);
 Z   ALTER TABLE ONLY public.user_checkins DROP CONSTRAINT user_checkins_pic_small_id_foreign;
       public       ueoddrluibf30n    false    200    207    4317            A           2606    18377 #   user_checkins_pic_square_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY user_checkins
    ADD CONSTRAINT user_checkins_pic_square_id_foreign FOREIGN KEY (pic_square_id) REFERENCES photos(photo_id);
 [   ALTER TABLE ONLY public.user_checkins DROP CONSTRAINT user_checkins_pic_square_id_foreign;
       public       ueoddrluibf30n    false    200    4317    207            B           2606    18382    user_checkins_user_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY user_checkins
    ADD CONSTRAINT user_checkins_user_id_foreign FOREIGN KEY (user_id) REFERENCES users(user_id);
 U   ALTER TABLE ONLY public.user_checkins DROP CONSTRAINT user_checkins_user_id_foreign;
       public       ueoddrluibf30n    false    4341    207    211            C           2606    18387 #   user_media_listens_media_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY user_media_listens
    ADD CONSTRAINT user_media_listens_media_id_foreign FOREIGN KEY (media_id) REFERENCES media(media_id);
 `   ALTER TABLE ONLY public.user_media_listens DROP CONSTRAINT user_media_listens_media_id_foreign;
       public       ueoddrluibf30n    false    4313    197    208            D           2606    18392 "   user_notifications_user_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY user_notifications
    ADD CONSTRAINT user_notifications_user_id_foreign FOREIGN KEY (user_id) REFERENCES users(user_id);
 _   ALTER TABLE ONLY public.user_notifications DROP CONSTRAINT user_notifications_user_id_foreign;
       public       ueoddrluibf30n    false    4341    211    209            E           2606    18397 (   user_payments_payment_service_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY user_payments
    ADD CONSTRAINT user_payments_payment_service_id_foreign FOREIGN KEY (payment_service_id) REFERENCES payment_services(payment_service_id);
 `   ALTER TABLE ONLY public.user_payments DROP CONSTRAINT user_payments_payment_service_id_foreign;
       public       ueoddrluibf30n    false    199    4315    210            F           2606    18402    user_payments_user_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY user_payments
    ADD CONSTRAINT user_payments_user_id_foreign FOREIGN KEY (user_id) REFERENCES users(user_id);
 U   ALTER TABLE ONLY public.user_payments DROP CONSTRAINT user_payments_user_id_foreign;
       public       ueoddrluibf30n    false    210    211    4341            G           2606    18407 %   users_current_listen_media_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY users
    ADD CONSTRAINT users_current_listen_media_id_foreign FOREIGN KEY (current_listen_media_id) REFERENCES media(media_id);
 U   ALTER TABLE ONLY public.users DROP CONSTRAINT users_current_listen_media_id_foreign;
       public       ueoddrluibf30n    false    4313    197    211            H           2606    18412    users_pic_big_id_foreign    FK CONSTRAINT     y   ALTER TABLE ONLY users
    ADD CONSTRAINT users_pic_big_id_foreign FOREIGN KEY (pic_big_id) REFERENCES photos(photo_id);
 H   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pic_big_id_foreign;
       public       ueoddrluibf30n    false    4317    211    200            I           2606    18417    users_pic_cover_id_foreign    FK CONSTRAINT     }   ALTER TABLE ONLY users
    ADD CONSTRAINT users_pic_cover_id_foreign FOREIGN KEY (pic_cover_id) REFERENCES photos(photo_id);
 J   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pic_cover_id_foreign;
       public       ueoddrluibf30n    false    211    4317    200            J           2606    18422    users_pic_small_id_foreign    FK CONSTRAINT     }   ALTER TABLE ONLY users
    ADD CONSTRAINT users_pic_small_id_foreign FOREIGN KEY (pic_small_id) REFERENCES photos(photo_id);
 J   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pic_small_id_foreign;
       public       ueoddrluibf30n    false    4317    200    211            K           2606    18427    users_pic_square_id_foreign    FK CONSTRAINT        ALTER TABLE ONLY users
    ADD CONSTRAINT users_pic_square_id_foreign FOREIGN KEY (pic_square_id) REFERENCES photos(photo_id);
 K   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pic_square_id_foreign;
       public       ueoddrluibf30n    false    4317    211    200            L           2606    18432 )   users_twitter_account_settings_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY users
    ADD CONSTRAINT users_twitter_account_settings_id_foreign FOREIGN KEY (twitter_account_settings_id) REFERENCES twitter_account_settings(twitter_account_settings_id);
 Y   ALTER TABLE ONLY public.users DROP CONSTRAINT users_twitter_account_settings_id_foreign;
       public       ueoddrluibf30n    false    211    4329    206            M           2606    18437    users_user_payment_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY users
    ADD CONSTRAINT users_user_payment_id_foreign FOREIGN KEY (user_payment_id) REFERENCES user_payments(user_payment_id);
 M   ALTER TABLE ONLY public.users DROP CONSTRAINT users_user_payment_id_foreign;
       public       ueoddrluibf30n    false    210    4337    211            N           2606    18442    venues_pic_big_id_foreign    FK CONSTRAINT     {   ALTER TABLE ONLY venues
    ADD CONSTRAINT venues_pic_big_id_foreign FOREIGN KEY (pic_big_id) REFERENCES photos(photo_id);
 J   ALTER TABLE ONLY public.venues DROP CONSTRAINT venues_pic_big_id_foreign;
       public       ueoddrluibf30n    false    200    213    4317            O           2606    18447    venues_pic_cover_id_foreign    FK CONSTRAINT        ALTER TABLE ONLY venues
    ADD CONSTRAINT venues_pic_cover_id_foreign FOREIGN KEY (pic_cover_id) REFERENCES photos(photo_id);
 L   ALTER TABLE ONLY public.venues DROP CONSTRAINT venues_pic_cover_id_foreign;
       public       ueoddrluibf30n    false    4317    213    200            P           2606    18452    venues_pic_id_foreign    FK CONSTRAINT     s   ALTER TABLE ONLY venues
    ADD CONSTRAINT venues_pic_id_foreign FOREIGN KEY (pic_id) REFERENCES photos(photo_id);
 F   ALTER TABLE ONLY public.venues DROP CONSTRAINT venues_pic_id_foreign;
       public       ueoddrluibf30n    false    213    200    4317            Q           2606    18457    venues_pic_small_id_foreign    FK CONSTRAINT        ALTER TABLE ONLY venues
    ADD CONSTRAINT venues_pic_small_id_foreign FOREIGN KEY (pic_small_id) REFERENCES photos(photo_id);
 L   ALTER TABLE ONLY public.venues DROP CONSTRAINT venues_pic_small_id_foreign;
       public       ueoddrluibf30n    false    200    213    4317            R           2606    18462    venues_pic_square_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY venues
    ADD CONSTRAINT venues_pic_square_id_foreign FOREIGN KEY (pic_square_id) REFERENCES photos(photo_id);
 M   ALTER TABLE ONLY public.venues DROP CONSTRAINT venues_pic_square_id_foreign;
       public       ueoddrluibf30n    false    4317    200    213            S           2606    18467 $   voucher_purchase_purchase_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY voucher_purchase
    ADD CONSTRAINT voucher_purchase_purchase_id_foreign FOREIGN KEY (purchase_id) REFERENCES purchases(purchase_id);
 _   ALTER TABLE ONLY public.voucher_purchase DROP CONSTRAINT voucher_purchase_purchase_id_foreign;
       public       ueoddrluibf30n    false    203    4323    215            T           2606    18472 #   voucher_purchase_voucher_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY voucher_purchase
    ADD CONSTRAINT voucher_purchase_voucher_id_foreign FOREIGN KEY (voucher_id) REFERENCES vouchers(voucher_id);
 ^   ALTER TABLE ONLY public.voucher_purchase DROP CONSTRAINT voucher_purchase_voucher_id_foreign;
       public       ueoddrluibf30n    false    215    4347    216            U           2606    18477 !   vouchers_issuer_artist_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY vouchers
    ADD CONSTRAINT vouchers_issuer_artist_id_foreign FOREIGN KEY (issuer_artist_id) REFERENCES artists(artist_id);
 T   ALTER TABLE ONLY public.vouchers DROP CONSTRAINT vouchers_issuer_artist_id_foreign;
       public       ueoddrluibf30n    false    4293    188    216           