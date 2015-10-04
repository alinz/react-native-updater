CREATE SEQUENCE release_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE releases (
    id bigint DEFAULT nextval('release_id_seq'::regclass) NOT NULL,
    app_id bigint NOT NULL,
    platform int NOT NULL,
    note text DEFAULT '',
    version bigint NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL
);

ALTER TABLE ONLY releases ADD CONSTRAINT releases_pkey PRIMARY KEY (id);
