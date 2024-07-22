-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS agent_details_id_seq;

-- Table Definition
CREATE TABLE "public"."agent_details" (
    "id" int4 NOT NULL DEFAULT nextval('agent_details_id_seq'::regclass),
    "name" varchar NOT NULL,
    "description" varchar,
    "image_url" varchar,
    "prompt" jsonb,
    "firebase_id" varchar,
    "framework_id" int4,
    "role_framework_id" int4,
    "total_used" int4,
    "created_at" timestamp,
    "updated_at" timestamp,
    "deleted_at" timestamp,
    PRIMARY KEY ("id")
);