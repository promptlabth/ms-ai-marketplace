-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS frameworks_id_seq;

-- Table Definition
CREATE TABLE "public"."frameworks" (
    "id" int4 NOT NULL DEFAULT nextval('frameworks_id_seq'::regclass),
    "name" varchar NOT NULL,
    "detail" varchar,
    "component" json,
    "language" varchar,
    PRIMARY KEY ("id")
);