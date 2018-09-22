-- Adminer 4.6.2 PostgreSQL dump

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


INSERT INTO "user" ("id", "login", "password", "role") VALUES
(1,	'admin',	'21232f297a57a5a743894a0e4a801fc3',	1);


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
    CONSTRAINT "agency_id" PRIMARY KEY ("id"),
    CONSTRAINT "agency_id_fkey" FOREIGN KEY (id) REFERENCES "user"(id) ON UPDATE SET NULL ON DELETE SET NULL NOT DEFERRABLE
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
    "unemployed" smallint DEFAULT '0' NOT NULL,
    "newcomer" smallint DEFAULT '0' NOT NULL,
    "homeless" smallint DEFAULT '0' NOT NULL,
    "special_needs" smallint DEFAULT '0' NOT NULL,
    CONSTRAINT "client_id" PRIMARY KEY ("id"),
    CONSTRAINT "client_agency_id_fkey" FOREIGN KEY (agency_id) REFERENCES agency(id) ON UPDATE SET NULL ON DELETE SET NULL NOT DEFERRABLE,
    CONSTRAINT "client_id_fkey" FOREIGN KEY (id) REFERENCES "user"(id) ON UPDATE SET NULL ON DELETE SET NULL NOT DEFERRABLE
) WITH (oids = false);


DROP TABLE IF EXISTS "gear";
DROP SEQUENCE IF EXISTS gear_id_seq;
CREATE SEQUENCE gear_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1;

CREATE TABLE "public"."gear" (
    "id" integer DEFAULT nextval('gear_id_seq') NOT NULL,
    "name" character varying NOT NULL,
    CONSTRAINT "gear_id" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "referral";
DROP SEQUENCE IF EXISTS referral_id_seq;
CREATE SEQUENCE referral_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1;

CREATE TABLE "public"."referral" (
    "id" integer DEFAULT nextval('referral_id_seq') NOT NULL,
    "client_id" integer NOT NULL,
    "appointment1" timestamp NOT NULL,
    "appointment2" timestamp NOT NULL,
    CONSTRAINT "referral_id" PRIMARY KEY ("id"),
    CONSTRAINT "referral_client_id_fkey" FOREIGN KEY (client_id) REFERENCES client(id) ON UPDATE SET NULL ON DELETE SET NULL NOT DEFERRABLE
) WITH (oids = false);


DROP TABLE IF EXISTS "referral_gear";
CREATE TABLE "public"."referral_gear" (
    "referral_id" integer NOT NULL,
    "gear_id" integer NOT NULL,
    "status" integer NOT NULL,
    CONSTRAINT "referral_gear_gear_id_fkey" FOREIGN KEY (gear_id) REFERENCES gear(id) ON UPDATE SET NULL ON DELETE SET NULL NOT DEFERRABLE,
    CONSTRAINT "referral_gear_referral_id_fkey" FOREIGN KEY (referral_id) REFERENCES referral(id) ON UPDATE SET NULL ON DELETE SET NULL NOT DEFERRABLE
) WITH (oids = false);


-- 2018-09-22 19:48:01.263279+00