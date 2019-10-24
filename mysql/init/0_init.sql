CREATE DATABASE
IF NOT EXISTS hash;

CREATE TABLE
IF NOT EXISTS hash.users
(
  id int PRIMARY KEY AUTO_INCREMENT,
  name varchar(64) NOT NULL,
  display_name varchar(64) NOT NULL,
  email varchar(64) NOT NULL,
  password varchar(64) NOT NULL,
  created_at datetime  default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp
);

CREATE TABLE
IF NOT EXISTS hash.tags
(
  id int PRIMARY KEY,
  name varchar(64) NOT NULL,
  created_at datetime  default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp
);

CREATE TABLE
IF NOT EXISTS hash.user_tags
(
  user_id int NOT NULL,
  tag_id int NOT NULL,
  created_at datetime  default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp
);