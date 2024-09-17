CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE content
(
    id    INT,
    title VARCHAR(63),
    text  VARCHAR(4095),
    url   VARCHAR(2047)
);
ALTER TABLE content ADD CONSTRAINT content_id_pk PRIMARY KEY (id);
ALTER TABLE content ALTER title SET NOT NULL;
ALTER TABLE content ALTER text SET NOT NULL;
ALTER TABLE content ALTER url SET NOT NULL;
/*TODO: добавить serial для article*/
CREATE TABLE tag
(
    id INT
);
ALTER TABLE tag ADD CONSTRAINT tag_id_pk PRIMARY KEY (id);

CREATE TABLE feature
(
    id INT
);
ALTER TABLE feature ADD CONSTRAINT feature_id_pk PRIMARY KEY (id);

CREATE TABLE banner
(
    id              INT,
    content_id      INT,
    is_active       BOOLEAN,
    created_at      TIMESTAMP,
    updated_at      TIMESTAMP,
    current_version INT,
    article         INT
);

ALTER TABLE banner ADD CONSTRAINT banner_id_pk PRIMARY KEY (id);
ALTER TABLE banner
    ADD CONSTRAINT banner_contentid_fk_content FOREIGN KEY (content_id) REFERENCES content (id);
ALTER TABLE banner ALTER COLUMN is_active SET NOT NULL;
ALTER TABLE banner ALTER COLUMN created_at SET NOT NULL;
ALTER TABLE banner ALTER COLUMN current_version SET NOT NULL;
ALTER TABLE banner ALTER COLUMN article SET NOT NULL;

CREATE TABLE banner_tag_feature_associative
(
    id         INT, /*optional*/
    banner_id  INT,
    tag_id     INT,
    feature_id INT,
    article    INT
);
ALTER TABLE banner_tag_feature_associative
    ADD CONSTRAINT bannertagfeatureassociative_tagid_feautereid_unique
        UNIQUE(tag_id, feature_id);
ALTER TABLE banner_tag_feature_associative
    ADD CONSTRAINT bannertagfeatureassociative_id_pk PRIMARY KEY (id);
ALTER TABLE banner_tag_feature_associative
    ADD CONSTRAINT bannertagfeatureassociative_bannerid_fk_banner
        FOREIGN KEY (banner_id) REFERENCES banner (id);
ALTER TABLE banner_tag_feature_associative
    ADD CONSTRAINT bannertagfeatureassociative_featureid_fk_feature
        FOREIGN KEY (feature_id) REFERENCES feature (id);
ALTER TABLE banner_tag_feature_associative
    ADD CONSTRAINT bannertagfeatureassociative_tagid_fk_tag
        FOREIGN KEY (tag_id) REFERENCES tag (id);
ALTER TABLE banner_tag_feature_associative ALTER COLUMN article SET NOT NULL;
ALTER TABLE banner_tag_feature_associative ALTER COLUMN feature_id SET NOT NULL;
ALTER TABLE banner_tag_feature_associative ALTER COLUMN tag_id SET NOT NULL;
ALTER TABLE banner_tag_feature_associative ALTER COLUMN banner_id SET NOT NULL;

CREATE TABLE banner_actual_version
(
    id              INT,
    banner_id       INT,
    article         INT,
    current_version INT
);
ALTER TABLE banner_actual_version ADD CONSTRAINT banneractualversion_id_pk PRIMARY KEY (id);
ALTER TABLE banner_actual_version
    ADD CONSTRAINT banneractualversion_bannerid_fk_banner
        FOREIGN KEY (banner_id) REFERENCES banner (id);
ALTER TABLE banner_actual_version ALTER banner_id SET NOT NULL;
ALTER TABLE banner_actual_version ALTER article SET NOT NULL;
ALTER TABLE banner_actual_version ALTER current_version SET NOT NULL;

CREATE TABLE account
(
    id       UUID,
    login    VARCHAR(63),
    password bytea,
    is_admin BOOLEAN,
    tag_id   INT
);
ALTER TABLE account ALTER COLUMN id SET DEFAULT uuid_generate_v4();
ALTER TABLE account ADD CONSTRAINT account_id_pk PRIMARY KEY (id);
/*ALTER TABLE account ADD CONSTRAINT account_tagid_fk_tag FOREIGN KEY (tag_id) REFERENCES tag (id);*/
ALTER TABLE account ALTER login SET NOT NULL;
ALTER TABLE account ALTER password SET NOT NULL;
ALTER TABLE account ALTER is_admin SET NOT NULL;
/*ALTER TABLE account ALTER tag_id SET NOT NULL;*/

CREATE TABLE refresh_token
(
    token_id UUID,
    user_id UUID,
    exp_date TIMESTAMP
);
ALTER TABLE refresh_token ADD CONSTRAINT refreshtoken_id_pk PRIMARY KEY (token_id);
ALTER TABLE refresh_token ADD CONSTRAINT refreshtoken_userid_fk_account FOREIGN KEY (user_id) REFERENCES account (id);
ALTER TABLE refresh_token ADD CONSTRAINT refreshtoken_tokentid_userid_unique UNIQUE(token_id, user_id);
ALTER TABLE refresh_token ALTER user_id SET NOT NULL;
ALTER TABLE refresh_token ALTER exp_date SET NOT NULL;
