DROP TABLE IF EXISTS apps;
DROP SEQUENCE IF EXISTS app_id_seq;

CREATE SEQUENCE app_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE apps (
    id bigint DEFAULT nextval('app_id_seq'::regclass) NOT NULL,
    name varchar(256) NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL
);

ALTER TABLE ONLY apps ADD CONSTRAINT apps_pkey PRIMARY KEY (id);
