CREATE DATABASE
IF NOT EXISTS hash;

CREATE TABLE
IF NOT EXISTS hash.users
(
  id int,
  name varchar(64),
  display_name varchar(64),
  email varchar(64),
  password varchar(64),
  created_at datetime  default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp
);
