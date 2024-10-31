-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS users_id_seq;

-- Table Definition
CREATE TABLE "public"."users" (
    "id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "firebase_id" text,
    "name" text,
    "email" text,
    "platform" text,
    "stripe_id" text,
    "plan_id" text,
    "datetime_last_active" timestamp,
    "profile_pic" text,
    "access_token" text,
    PRIMARY KEY ("id")
);