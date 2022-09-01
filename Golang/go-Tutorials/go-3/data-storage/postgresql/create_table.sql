CREATE TABLE public.teacher
(
    id serial,
    create_time timestamp without time zone,
    update_time timestamp without time zone,
    firstname character varying(255),
    lastname character varying(255),
    PRIMARY KEY (id)
);