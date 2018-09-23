-- Adminer 4.6.2 PostgreSQL dump

DROP TABLE IF EXISTS "agency";
DROP SEQUENCE IF EXISTS agency_id_seq;
CREATE SEQUENCE agency_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1;

CREATE TABLE "public"."agency" (
    "id" integer DEFAULT nextval('agency_id_seq') NOT NULL,
    "login" character varying NOT NULL,
    "password" character varying NOT NULL,
    "role" integer NOT NULL,
    "name" character varying NOT NULL,
    "phone" character varying NOT NULL,
    "email" character varying NOT NULL,
    "city" character varying NOT NULL,
    "address1" character varying NOT NULL,
    "address2" character varying NOT NULL,
    "contact" character varying NOT NULL,
    CONSTRAINT "agency_id" PRIMARY KEY ("id")
) WITH (oids = false);

INSERT INTO "agency" ("login", "password", "role", "name", "phone", "email", "city", "address1", "address2", "contact") VALUES
('admin', '21232f297a57a5a743894a0e4a801fc3',	1,	'BabyGoRound',	'',	'',	'',	'',	'',	'');

DROP TABLE IF EXISTS "client";
DROP SEQUENCE IF EXISTS client_id_seq;
CREATE SEQUENCE client_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1;

CREATE TABLE "public"."client" (
    "id" integer DEFAULT nextval('client_id_seq') NOT NULL,
    "status" integer NOT NULL,
    "name" character varying NOT NULL,
    "dob" date NOT NULL,
    "childdob" date NOT NULL,
    "phone" character varying NOT NULL,
    "email" character varying NOT NULL,
    "city" character varying NOT NULL,
    "address1" character varying NOT NULL,
    "address2" character varying NOT NULL,
    "agency_id" integer NOT NULL,
    "unemployed" smallint DEFAULT '0' NOT NULL,
    "newcomer" smallint DEFAULT '0' NOT NULL,
    "homeless" smallint DEFAULT '0' NOT NULL,
    "special_needs" smallint DEFAULT '0' NOT NULL,
    CONSTRAINT "client_id" PRIMARY KEY ("id"),
    CONSTRAINT "client_agency_id_fkey" FOREIGN KEY (agency_id) REFERENCES agency(id) ON UPDATE SET NULL ON DELETE SET NULL NOT DEFERRABLE
) WITH (oids = false);


DROP TABLE IF EXISTS "gear";
DROP SEQUENCE IF EXISTS gear_id_seq;
CREATE SEQUENCE gear_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1;

CREATE TABLE "public"."gear" (
    "id" integer DEFAULT nextval('gear_id_seq') NOT NULL,
    "name" character varying NOT NULL,
    CONSTRAINT "gear_id" PRIMARY KEY ("id")
) WITH (oids = false);

INSERT INTO "gear" ("name") VALUES
('Crib'),
('Bassinet'),
('Pack ''n play'),
('Single stroller'),
('Double stroller'),
('Front carrier'),
('Bouncy chair'),
('Swing'),
('Exersaucer'),
('Jolly jumper'),
('Bumbo'),
('High chair'),
('Bathtub'),
('Diapers'),
('Diaper bag'),
('Blankets'),
('Crib bedding'),
('Sleepsacks'),
('Clothing'),
('Toys/books'),
('Nursing pillow'),
('Safety gate'),
('Bottles'),
('Feeding accessories'),
('Monitor'),
('Safety gear'),
('Breast pump');

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


-- 2018-09-22 20:36:44.110215+00

-- Demo
INSERT INTO "agency" ("login", "password", "role", "name", "phone", "email", "city", "address1", "address2", "contact") VALUES
('agency', '21232f297a57a5a743894a0e4a801fc3',	2,	'Awesome Agency',	'5551234567',	'agency@gmail.com',	'Vancouver',	'123 Main st.',	'',	'Agent');

INSERT INTO "client" (status, name, dob, childdob, phone, email, city, address1, address2, agency_id, unemployed, newcomer, homeless, special_needs) VALUES
(0, 'John Smith',	'1990-01-01', '2017-01-01',	'5551234567',	'john@gmail.com',	'Richmond',	'123 No.3 Rd.',	'',	2, 1, 1, 0, 0);
--