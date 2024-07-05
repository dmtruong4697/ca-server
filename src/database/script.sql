
CREATE TABLE users(
  id INT NOT NULL,
  user_name VARCHAR,
  email VARCHAR,
  `password` VARCHAR,
  phone_number VARCHAR,
  avatar_image VARCHAR,
  hashtag_name VARCHAR,
  create_at DATETIME,
  account_status VARCHAR,
  validate_code VARCHAR,
  last_active DATETIME,
  gender VARCHAR,
  date_of_birth DATETIME,
  device_token VARCHAR,
  PRIMARY KEY(id)
);

CREATE TABLE messages(
  id INT NOT NULL,
  sender_id INT,
  group_id INT,
  create_at DATETIME,
  content VARCHAR,
  is_edited BOOLEAN,
  last_update DATETIME,
  `status` VARCHAR,
  PRIMARY KEY(id)
);

CREATE TABLE media_message(
  id INT NOT NULL,
  `url` VARCHAR,
  create_at DATETIME,
  sender_id INT,
  message_id INT,
  `type` VARCHAR,
  PRIMARY KEY(id)
);

CREATE TABLE `groups`(
  id INT NOT NULL,
  creator_id INT,
  `name` VARCHAR,
  `status` VARCHAR,
  create_at DATETIME,
  last_message_id INT,
  group_image_url VARCHAR,
  `type` VARCHAR,
  invite_code VARCHAR,
  is_allow_invite_code BOOLEAN,
  PRIMARY KEY(id)
);

CREATE TABLE group_member(
  id INT NOT NULL,
  user_id INT,
  group_id INT,
  join_at DATETIME,
  `status` VARCHAR,
  ingroup_name VARCHAR,
  PRIMARY KEY(id)
);

CREATE TABLE relationships(
  id INT NOT NULL,
  sender_id INT,
  receiver_id INT,
  create_at DATETIME,
  `status` VARCHAR,
  PRIMARY KEY(id)
);

CREATE TABLE notifications(
  id INT NOT NULL,
  user_id INT,
  create_at DATETIME,
  is_read BOOLEAN,
  content VARCHAR,
  PRIMARY KEY(id)
);

ALTER TABLE media_message
  ADD CONSTRAINT id_message_id FOREIGN KEY (message_id) REFERENCES messages (id)
  ;

ALTER TABLE messages
  ADD CONSTRAINT id_sender_id FOREIGN KEY (sender_id) REFERENCES users (id);

ALTER TABLE media_message
  ADD CONSTRAINT id_sender_id FOREIGN KEY (sender_id) REFERENCES users (id);

ALTER TABLE group_member
  ADD CONSTRAINT id_user_id FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE group_member
  ADD CONSTRAINT id_group_id FOREIGN KEY (group_id) REFERENCES `groups` (id);

ALTER TABLE messages
  ADD CONSTRAINT id_group_id FOREIGN KEY (group_id) REFERENCES `groups` (id);

ALTER TABLE relationships
  ADD CONSTRAINT id_sender_id FOREIGN KEY (sender_id) REFERENCES users (id);

ALTER TABLE relationships
  ADD CONSTRAINT id_receiver_id FOREIGN KEY (receiver_id) REFERENCES users (id);

ALTER TABLE notifications
  ADD CONSTRAINT id_user_id FOREIGN KEY (user_id) REFERENCES users (id);
Main Diagram
user_name
