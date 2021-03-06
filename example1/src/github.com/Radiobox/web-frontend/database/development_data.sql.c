PGDMP     -                    r           radiobox    9.3.4    9.3.4 1    �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                       false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                       false            ~          0    47681    photos 
   TABLE DATA               �   COPY photos (photo_id, caption, file_store_id, file_store_key, format, height, "media_MD5", media_mimetype, owner_id, owner_type, secret_url, url, width, created_at, updated_at, options) FROM stdin;
    public       radiobox    false    208   �;       x          0    47641    media_formats 
   TABLE DATA               �   COPY media_formats (media_format_id, bitrate, channels, format, is_lossless, name, pic_id, pic_url, created_at, updated_at) FROM stdin;
    public       radiobox    false    202   �;       w          0    47631    media 
   TABLE DATA               �   COPY media (media_id, file_store_id, file_store_key, is_preview, media_format_id, "media_MD5", secret_url, title, created_at, updated_at, original_file_name, bucket_name, content_type, content_length, status, bitrate, duration) FROM stdin;
    public       radiobox    false    201   �;       }          0    47672    payment_services 
   TABLE DATA               f   COPY payment_services (payment_service_id, name, script_location, created_at, updated_at) FROM stdin;
    public       radiobox    false    207   �;       �          0    47743    twitter_account_settings 
   TABLE DATA               �   COPY twitter_account_settings (twitter_account_settings_id, handle, token1_encrypted, token2_encrypted, created_at, updated_at) FROM stdin;
    public       radiobox    false    216   <       �          0    47780    user_payments 
   TABLE DATA               g  COPY user_payments (user_payment_id, payment_service_id, token, token_salt, payment_identifier_safe_string, payment_identifier_name, payment_expiration, bill_name_first, bill_name_mi, bill_name_last, bill_address_1, bill_address_2, bill_address_verified, bill_city, bill_country, bill_state, bill_notes, currency, user_id, created_at, updated_at) FROM stdin;
    public       radiobox    false    220   $<       �          0    47797    users 
   TABLE DATA               4  COPY users (user_id, birthday_date, current_address, current_address_point, current_listen_media_id, current_location, current_location_point, contact_phone, contact_phone_country_code, contact_phone_extension, email_proxy, email, email_pending_update, facebook_token_encrypted, facebook_user, first_name, last_name, interests, keywords, locale, pass_hash, pic_cover, pic_cover_id, pic_square, pic_square_id, political, profile_blurb, user_payment_id, relationship_status, sex, rb_subscriber_count, facebook_subscriber_count, facebook_last_data_pull, twitter_subscriber_count, twitter_last_data_pull, twitter_account_settings_id, linkedin_token_encrypted, linkedin_user, linkedin_last_data_pull, timezone, username, wall_count, work, is_active, is_verified, created_at, updated_at, current_location_updated) FROM stdin;
    public       radiobox    false    223   A<       j          0    47531    artists 
   TABLE DATA               �  COPY artists (artist_id, band_user_ids, booking_user_id, contact_email, contact_phone, contact_phone_country_code, contact_phone_extension, facebook_page_id, facebook_page_token_encrypted, hometown_address, hometown_address_point, keywords, manager_user_id, pic_cover, pic_square, pic_square_id, record_label_id, subscriber_count, timezone, twitter_account_settings_id, username, website, is_active, is_verified, created_at, updated_at, description, pic_cover_id) FROM stdin;
    public       radiobox    false    188   �<       f          0    47487    albums 
   TABLE DATA               �   COPY albums (album_id, artist_id, downloads, title, pic_square, pic_square_id, price_usd, price_usd_min, price_nyo, purchases, time_length, is_active, created_at, updated_at) FROM stdin;
    public       radiobox    false    184   �<       �          0    47727    tracks 
   TABLE DATA               <  COPY tracks (track_id, artist_id, checksum_md5, downloads, format_keyval, pic_square, pic_square_id, media_id, must_bundle_album_id, title, preview_media_id, preview_only, price_usd, price_usd_min, price_nyo, purchases, time_length, is_active, created_at, updated_at, encoding_media_id, encoding_status) FROM stdin;
    public       radiobox    false    215   =       e          0    47481    album_track 
   TABLE DATA               h   COPY album_track (album_track_id, album_id, track_id, created_at, updated_at, track_number) FROM stdin;
    public       radiobox    false    183   3=       �          0    47810    venues 
   TABLE DATA               +  COPY venues (venue_id, address, email, facebook_token, facebook_user, address_point, name, contact_phone, pic_cover, pic_cover_id, pic_square, pic_square_id, url, venue_blurb, is_active, is_verified, created_at, updated_at, timezone, contact_phone_country_code, contact_phone_extension) FROM stdin;
    public       radiobox    false    225   P=       g          0    47503    artist_checkins 
   TABLE DATA               p  COPY artist_checkins (artist_checkin_id, checkin_via, is_location_matched, checkin_scan_code, checkin_location_point, checkin_location_match_service, checkin_location_match_name, checkin_location_venue_match_id, checkin_message, pic_big, pic_big_id, pic_small, pic_small_id, pic_square, pic_square_id, checkin_timestamp, artist_id, created_at, updated_at) FROM stdin;
    public       radiobox    false    185   m=       h          0    47512    artist_notifications 
   TABLE DATA               p   COPY artist_notifications (artist_notification_id, artist_id, notification, created_at, updated_at) FROM stdin;
    public       radiobox    false    186   �=       i          0    47521    artist_user 
   TABLE DATA               j   COPY artist_user (artist_user_id, artist_id, user_id, is_admin, role, created_at, updated_at) FROM stdin;
    public       radiobox    false    187   �=       �           0    0    artists_artist_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('artists_artist_id_seq', 1, false);
            public       radiobox    false    189            l          0    47544    beta_signups 
   TABLE DATA               C   COPY beta_signups (beta_signup_id, artist_name, email) FROM stdin;
    public       radiobox    false    190   �=       �           0    0    beta_signups_beta_signup_id_seq    SEQUENCE SET     G   SELECT pg_catalog.setval('beta_signups_beta_signup_id_seq', 1, false);
            public       radiobox    false    191            n          0    47552    email_templates 
   TABLE DATA               �   COPY email_templates (email_template_id, name, body_plaintext_link, body_html_link, subject, language, notes, version) FROM stdin;
    public       radiobox    false    192   �=       t          0    47611    events_live 
   TABLE DATA               �  COPY events_live (event_live_id, actual_end, actual_start, archive_track_id, archiving, artist_id, title, title_from, artist_collaborators_id, is_in_progress, is_concluded, off_schedule_end_seconds, off_schedule_start_seconds, scheduled_end, scheduled_start, location_address, location_address_point, venue_id, listeners_max, is_active, created_at, updated_at, caption, is_standby, messages, standby_message, pic_square_id, pic_square) FROM stdin;
    public       radiobox    false    198   �=       o          0    47559    event_admin_log 
   TABLE DATA               c   COPY event_admin_log (event_admin_log_id, event_live_id, data, created_at, updated_at) FROM stdin;
    public       radiobox    false    193   >       p          0    47568    event_live_notifications 
   TABLE DATA               �   COPY event_live_notifications (event_live_notification_id, event_live_id, notification, title, created_at, updated_at) FROM stdin;
    public       radiobox    false    194   8>       q          0    47577 !   event_live_provisioned_broadcasts 
   TABLE DATA               �  COPY event_live_provisioned_broadcasts (event_live_provisioned_broadcast_id, type, is_video, broadcast_url_1_port, broadcast_url_2_port, broadcast_stream_name, provider_stream_id, broadcast_url_1, broadcast_url_2, broadcast_method, broadcast_username, broadcast_password, encode_suggested_params, encode_accepted_params, client_hds, client_hls, client_hdflash1, client_shoutcast_url, reserved_until, is_in_progress, is_concluded, is_available, listeners_max, listeners_now, created_at, updated_at) FROM stdin;
    public       radiobox    false    195   U>       r          0    47592    event_live_routes 
   TABLE DATA               �   COPY event_live_routes (event_live_route_id, event_live_id, event_live_provisioned_broadcast_id, priority, is_primary, listeners_max, listeners_now, created_at, updated_at) FROM stdin;
    public       radiobox    false    196   r>       s          0    47602    event_live_segments 
   TABLE DATA               �   COPY event_live_segments (event_live_segment_id, end_time, event_live_id, media_id, segment_name, start_time, created_at, updated_at) FROM stdin;
    public       radiobox    false    197   �>       u          0    47621    logs 
   TABLE DATA               N   COPY logs (log_id, log_type, log_details, created_at, updated_at) FROM stdin;
    public       radiobox    false    199   �>       �           0    0    logs_log_id_seq    SEQUENCE SET     6   SELECT pg_catalog.setval('logs_log_id_seq', 1, true);
            public       radiobox    false    200            y          0    47651 
   migrations 
   TABLE DATA               /   COPY migrations (migration, batch) FROM stdin;
    public       radiobox    false    203   j?       z          0    47654    oauth_access_data 
   TABLE DATA               �   COPY oauth_access_data (client, authorizedata, accessdata, accesstoken, refreshtoken, expiresin, scope, redirecturi, createdat, userid) FROM stdin;
    public       radiobox    false    204   �?       {          0    47660    oauth_auth_data 
   TABLE DATA               i   COPY oauth_auth_data (client, code, expiresin, scope, redirecturi, state, createdat, userid) FROM stdin;
    public       radiobox    false    205   �?       |          0    47666    oauth_clients 
   TABLE DATA               9   COPY oauth_clients (id, secret, redirecturi) FROM stdin;
    public       radiobox    false    206   �?       �          0    47702 	   purchases 
   TABLE DATA               �  COPY purchases (purchase_id, begin_cart, is_complete, is_paid, is_paypal, is_bitcoin, is_creditcard, pay_grand_total, pay_shipping, pay_sub_total, pay_discount, pay_tax, user_payment_id, receipt_email, ship_name_first, ship_name_mi, ship_name_last, ship_address_1, ship_address_2, ship_address_verified, ship_city, ship_country, ship_state, ship_notes, user_id, created_at, updated_at) FROM stdin;
    public       radiobox    false    211   �?                 0    47690    purchase_items 
   TABLE DATA               s   COPY purchase_items (purchase_item_id, purchase_id, purchase_type, unit_price, created_at, updated_at) FROM stdin;
    public       radiobox    false    209   @       �          0    47696    purchase_purchaseitem 
   TABLE DATA               y   COPY purchase_purchaseitem (purcahse_purchaseitem_id, purchase_id, purchase_item_id, created_at, updated_at) FROM stdin;
    public       radiobox    false    210   $@       �          0    47711    slugs 
   TABLE DATA               *   COPY slugs (id, target, slug) FROM stdin;
    public       radiobox    false    212   A@       �           0    0    slugs_id_seq    SEQUENCE SET     3   SELECT pg_catalog.setval('slugs_id_seq', 1, true);
            public       radiobox    false    213            �          0    46448    spatial_ref_sys 
   TABLE DATA               Q   COPY spatial_ref_sys (srid, auth_name, auth_srid, srtext, proj4text) FROM stdin;
    public       postgres    false    171   r@       �          0    47719    tokens 
   TABLE DATA               >   COPY tokens (token_id, user_id, expires_at, type) FROM stdin;
    public       radiobox    false    214   �@       �          0    47752    user_checkins 
   TABLE DATA               j  COPY user_checkins (user_checkin_id, checkin_via, is_location_matched, checkin_scan_code, checkin_location_point, checkin_location_match_service, checkin_location_match_name, checkin_location_venue_match_id, checkin_message, pic_big, pic_big_id, pic_small, pic_small_id, pic_square, pic_square_id, checkin_timestamp, user_id, created_at, updated_at) FROM stdin;
    public       radiobox    false    217   �@       �          0    47761    user_listens 
   TABLE DATA               �   COPY user_listens (user_listen_id, is_concluded, is_live, play_max_position, play_cursor_position, user_id, created_at, updated_at, asset) FROM stdin;
    public       radiobox    false    218   A       �          0    47771    user_notifications 
   TABLE DATA               w   COPY user_notifications (user_notification_id, link, notification, title, user_id, created_at, updated_at) FROM stdin;
    public       radiobox    false    219   5A       �          0    47789 	   usernames 
   TABLE DATA               2   COPY usernames (id, target, username) FROM stdin;
    public       radiobox    false    221   RA       �           0    0    usernames_id_seq    SEQUENCE SET     7   SELECT pg_catalog.setval('usernames_id_seq', 1, true);
            public       radiobox    false    222            �           0    0    users_user_id_seq    SEQUENCE SET     8   SELECT pg_catalog.setval('users_user_id_seq', 1, true);
            public       radiobox    false    224            �           0    0    venues_venue_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('venues_venue_id_seq', 1, false);
            public       radiobox    false    226            �          0    47828    vouchers 
   TABLE DATA               �   COPY vouchers (voucher_id, amount, issuer_artist_id, is_closed_out, expiration, bill_artist, bill_radiobox, created_at, updated_at) FROM stdin;
    public       radiobox    false    228   �A       �          0    47822    voucher_purchase 
   TABLE DATA               i   COPY voucher_purchase (voucher_purchase_id, voucher_id, purchase_id, created_at, updated_at) FROM stdin;
    public       radiobox    false    227   �A       ~      x������ � �      x      x������ � �      w      x������ � �      }      x������ � �      �      x������ � �      �      x������ � �      �   �   x�3��Â�tbq��%�cWD*F�*�*��)��aUa9���a�f���F��E%�ő��)��I�!����YIɑ��!���8%�����J8�8�MtLt�,M�L,�����L�-��K��q��qqq ��@�      j      x������ � �      f      x������ � �      �      x������ � �      e      x������ � �      �      x������ � �      g      x������ � �      h      x������ � �      i      x������ � �      l      x������ � �      n      x������ � �      t      x������ � �      o      x������ � �      p      x������ � �      q      x������ � �      r      x������ � �      s      x������ � �      u   �   x�}��
�0D��W��چ�F�9��(�E(�]1�&�D)���������Yfe���nL��(���0�0��C!g��n'�m���~�5��{ec��<�F3[�+��o��	_�юu���J,e���h)�=�&����J�z>U'�j�j���;��Z���J��F���]9c�"7F�      y      x������ � �      z      x������ � �      {      x������ � �      |      x�30 CΒ��0����� IJ�      �      x������ � �            x������ � �      �      x������ � �      �   !   x�3�V*-N-��LQ�2��L,N����� b�      �      x������ � �      �   \   x���
�  �g���c�֦ٷak��"���6Ե� �Ԁ��j�2�&N��QD��@�C��L��H�>Fggm�������m�����O'�      �      x������ � �      �      x������ � �      �      x������ � �      �   !   x�3�V*-N-��LQ�2��L,N����� b�      �      x������ � �      �      x������ � �     