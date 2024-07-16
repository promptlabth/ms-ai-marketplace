-- CREATE TABLE agent_details (
--   id INT AUTO_INCREMENT ,
--   name VARCHAR(255) NOT NULL UNIQUE,
--   description VARCHAR(255),
--   image_url VARCHAR(255),
--   prompt JSON,
--   user_id VARCHAR(255),
--   framework_id INTEGER,
--   role_framework_id INTEGER
-- );

-- CREATE TABLE role_frameworks (
--   id INT AUTO_INCREMENT ,
--   name VARCHAR(255) UNIQUE
-- );

-- CREATE TABLE roles (
--   id INT AUTO_INCREMENT ,
--   name VARCHAR(255) UNIQUE
-- );

-- CREATE TABLE frameworks (
--   id INT AUTO_INCREMENT ,
--   name VARCHAR(255) NOT NULL UNIQUE,
--   detail VARCHAR(255),
--   component JSON
-- );

-- CREATE TABLE users (
--   id INT AUTO_INCREMENT ,
--   firebase_id VARCHAR(255),
--   name VARCHAR(255),
--   profile_pic VARCHAR(255),
--   access_token VARCHAR(255),
--   stripe_id VARCHAR(255),
--   plan_id VARCHAR(255),
--   password VARCHAR(255),
--   datetime_last_active TIMESTAMP
-- );


CREATE TABLE agent_details (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE,
  description VARCHAR(255),
  image_url VARCHAR(255),
  prompt JSONB,
  user_id VARCHAR(255),
  framework_id INTEGER,
  role_framework_id INTEGER
);

CREATE TABLE role_frameworks (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  language VARCHAR(255)
);

CREATE TABLE roles (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  language VARCHAR(255)
);

CREATE TABLE frameworks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    detail VARCHAR(255),
    component JSONB,
    language VARCHAR(255)
);

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  firebase_id VARCHAR(255),
  name VARCHAR(255),
  profile_pic VARCHAR(255),
  access_token VARCHAR(255),
  stripe_id VARCHAR(255),
  plan_id VARCHAR(255),
  password VARCHAR(255),
  datetime_last_active TIMESTAMP
);

CREATE TABLE histohistoriesrys (
    id SERIAL PRIMARY KEY,
    user_id TEXT NOT NULL,
    agent_id INT NOT NULL,
    framework_id INT NOT NULL,
    prompt TEXT NOT NULL,
    style_message_id INT NOT NULL,
    language TEXT NOT NULL,
    result TEXT NOT NULL,
    time_stamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE style_prompts (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    language TEXT NOT NULL
);