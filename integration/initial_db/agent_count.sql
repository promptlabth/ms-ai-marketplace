-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS agent_counts_agent_id_seq;

-- Table Definition
CREATE TABLE "public"."agent_counts" (
    "agent_id" int4 NOT NULL DEFAULT nextval('agent_counts_agent_id_seq'::regclass),
    "count_used" int4,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("agent_id")
);