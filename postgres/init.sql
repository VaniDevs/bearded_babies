DROP TABLE IF EXISTS "agency";
CREATE TABLE "public"."agency" (
    "id" integer NOT NULL,
    "status" integer NOT NULL,
    "name" character varying NOT NULL,
    "phone" character varying NOT NULL,
    "email" character varying NOT NULL,
    "city" character varying NOT NULL,
    "address1" character varying NOT NULL,
    "address2" character varying NOT NULL,
    "contact" character varying NOT NULL,
    CONSTRAINT "agency_id" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "client";
CREATE TABLE "public"."client" (
    "id" integer NOT NULL,
    "status" integer NOT NULL,
    "name" character varying NOT NULL,
    "DOB" date NOT NULL,
    "childDOB" date NOT NULL,
    "phone" character varying NOT NULL,
    "email" character varying NOT NULL,
    "city" character varying NOT NULL,
    "address1" character varying NOT NULL,
    "address2" character varying NOT NULL,
    "notification" integer NOT NULL,
    "agency_id" integer NOT NULL,
    CONSTRAINT "client_id" PRIMARY KEY ("id"),
    CONSTRAINT "client_agency_id_fkey" FOREIGN KEY (agency_id) REFERENCES agency(id) ON UPDATE SET NULL ON DELETE SET NULL NOT DEFERRABLE,
    CONSTRAINT "client_id_fkey" FOREIGN KEY (id) REFERENCES "user"(id) NOT DEFERRABLE
) WITH (oids = false);


DROP TABLE IF EXISTS "user";
DROP SEQUENCE IF EXISTS user_id_seq;
CREATE SEQUENCE user_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1;

CREATE TABLE "public"."user" (
    "id" integer DEFAULT nextval('user_id_seq') NOT NULL,
    "login" character varying NOT NULL,
    "password" character varying NOT NULL,
    "role" integer NOT NULL,
    CONSTRAINT "user_id" PRIMARY KEY ("id")
) WITH (oids = false);