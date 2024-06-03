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
  name VARCHAR(255) UNIQUE
);

CREATE TABLE roles (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) UNIQUE
);

CREATE TABLE frameworks (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE,
  detail VARCHAR(255),
  component JSONB
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
