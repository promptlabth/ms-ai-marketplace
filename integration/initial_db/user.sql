-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS user_seq;

-- Table Definition
CREATE TABLE "public"."user" (
    "id" int4 NOT NULL DEFAULT nextval('user_seq'::regclass),
    "firebase_id" varchar,
    "name" varchar,
    "email" json,
    "profilepic" varchar,
    "platform" varchar ,
    "access_token" varchar,
    "stripe_id" varchar,
    "plan" varchar,
    
    PRIMARY KEY ("id")
);