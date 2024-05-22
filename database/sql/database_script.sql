CREATE TABLE agent_details (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE,
  description VARCHAR(255),
  image_url VARCHAR(255),
  prompt JSON,
  user_id VARCHAR(255),
  framework_id VARCHAR(255),
  role_framework_id VARCHAR(255)
);

CREATE TABLE role_frameworks (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) UNIQUE
);

CREATE TABLE frameworks (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE,
  detail VARCHAR(255),
  component JSON
);

CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  firebase_id VARCHAR(255),
  name VARCHAR(255),
  profile_pic VARCHAR(255),
  access_token VARCHAR(255),
  stripe_id VARCHAR(255),
  plan_id VARCHAR(255),
  password VARCHAR(255),
  datetime_last_active TIMESTAMP
);
