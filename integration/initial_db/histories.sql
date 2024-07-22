-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS histories_id_seq;

-- Table Definition
CREATE TABLE "public"."histories" (
    "id" int4 NOT NULL DEFAULT nextval('histories_id_seq'::regclass),
    "firebase_id" text NOT NULL,
    "agent_id" int4 NOT NULL,
    "framework_id" int4 NOT NULL,
    "prompt" text NOT NULL,
    "style_message_id" int4 NOT NULL,
    "language" text NOT NULL,
    "result" text NOT NULL,
    "time_stamp" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);