-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS style_prompts_id_seq;

-- Table Definition
CREATE TABLE "public"."style_prompts" (
    "id" int4 NOT NULL DEFAULT nextval('style_prompts_id_seq'::regclass),
    "name" text NOT NULL,
    "language" text NOT NULL,
    PRIMARY KEY ("id")
);

-- insert to table
INSERT INTO
    style_prompts (name, language)
VALUES ('fun', 'en'),
    ('Confident', 'en'),
    ('Professional', 'en'),
    ('Luxurious', 'en'),
    ('Educated', 'en'),
    ('Happy', 'en'),
    ('Modern', 'en'),
    ('Nostalgic', 'en'),
    ('สนุกสนาน', 'th'),
    ('มั่นใจ', 'th'),
    ('มืออาชีพ', 'th'),
    ('หรูหรา', 'th'),
    ('มีการศึกษา', 'th'),
    ('มีความสุข', 'th'),
    ('ทันสมัย', 'th'),
    ('ย้อนยุค', 'th');