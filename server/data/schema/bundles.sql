DROP TABLE IF EXISTS bundles;
DROP SEQUENCE IF EXISTS bundle_id_seq;

CREATE SEQUENCE bundle_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE bundles (
    id bigint DEFAULT nextval('bundle_id_seq'::regclass) NOT NULL,
    release_id bigint NOT NULL,
    hash varchar(128) NOT NULL,
    name varchar(64) NOT NULL,
    type int NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL
);

ALTER TABLE ONLY bundles ADD CONSTRAINT bundles_pkey PRIMARY KEY (id);
ALTER TABLE bundles ADD FOREIGN KEY ("release_id") REFERENCES releases("id") ON DELETE CASCADE ON UPDATE CASCADE;
