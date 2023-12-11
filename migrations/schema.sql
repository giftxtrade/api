--
-- Table structure for table categories
--
CREATE TABLE categories (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  description text,
  category_url varchar(255) NOT NULL,
  PRIMARY KEY (id)
);

--
-- Table structure for table draws
--
CREATE TABLE draws (
  id int NOT NULL AUTO_INCREMENT,
  created_at datetime NOT NULL,
  drawer_id int DEFAULT NULL,
  drawee_id int DEFAULT NULL,
  event_id int DEFAULT NULL,
  PRIMARY KEY (id),
  KEY IDX_5c6f8dcbad4254f1541d043b52 (drawer_id),
  KEY IDX_4a912e98dc9402aa438ff6e5ca (event_id),
  KEY FK_13049315c8546dd62939d537e97 (drawee_id),
  CONSTRAINT FK_13049315c8546dd62939d537e97 FOREIGN KEY (drawee_id) REFERENCES participants (id) ON DELETE CASCADE,
  CONSTRAINT FK_4a912e98dc9402aa438ff6e5ca5 FOREIGN KEY (event_id) REFERENCES events (id) ON DELETE CASCADE,
  CONSTRAINT FK_5c6f8dcbad4254f1541d043b52b FOREIGN KEY (drawer_id) REFERENCES participants (id) ON DELETE CASCADE
);

--
-- Table structure for table events
--
CREATE TABLE events (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  description text,
  budget decimal(10,0) NOT NULL,
  invitation_message text NOT NULL,
  created_at datetime NOT NULL,
  draw_at datetime NOT NULL,
  close_at datetime NOT NULL,
  PRIMARY KEY (id)
);

--
-- Table structure for table links
--
CREATE TABLE links (
  id int NOT NULL AUTO_INCREMENT,
  code varchar(255) NOT NULL,
  created_at datetime NOT NULL,
  expiration_date datetime NOT NULL,
  event_id int DEFAULT NULL,
  PRIMARY KEY (id),
  KEY IDX_52a3fa2a2c27a987ed58fd2ea4 (code),
  KEY IDX_121cadf4bb99f7ba5dbac54997 (event_id),
  CONSTRAINT FK_121cadf4bb99f7ba5dbac54997a FOREIGN KEY (event_id) REFERENCES events (id) ON DELETE CASCADE
);

--
-- Table structure for table participants
--
CREATE TABLE participants (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  address varchar(255) NOT NULL,
  organizer tinyint NOT NULL,
  participates tinyint NOT NULL,
  accepted tinyint NOT NULL,
  event_id int DEFAULT NULL,
  user_id int DEFAULT NULL,
  PRIMARY KEY (id),
  KEY IDX_b77ad0832a0f8ec526c1f40a84 (email),
  KEY IDX_a622804301e735196918e6a47e (event_id),
  KEY IDX_5fc9cddc801b973cd9edcdda42 (user_id),
  CONSTRAINT FK_5fc9cddc801b973cd9edcdda42a FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
  CONSTRAINT FK_a622804301e735196918e6a47e5 FOREIGN KEY (event_id) REFERENCES events (id) ON DELETE CASCADE
);

--
-- Table structure for table products
--
CREATE TABLE products (
  id int NOT NULL AUTO_INCREMENT,
  title text NOT NULL,
  description text NOT NULL,
  product_key varchar(255) NOT NULL,
  image_url text NOT NULL,
  rating double NOT NULL,
  price double NOT NULL,
  currency varchar(255) NOT NULL,
  modified datetime NOT NULL,
  website text NOT NULL,
  categoryId int DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY IDX_add1c4bbc945b8fc7c3fdb4689 (product_key),
  KEY IDX_2394b26ba3206c1c59f461d3dc (rating),
  KEY IDX_75895eeb1903f8a17816dafe0a (price),
  KEY IDX_ff56834e735fa78a15d0cf2192 (categoryId),
  FULLTEXT KEY IDX_c30f00a871de74c8e8c213acc4 (title),
  CONSTRAINT FK_ff56834e735fa78a15d0cf21926 FOREIGN KEY (categoryId) REFERENCES categories (id) ON DELETE CASCADE
);

--
-- Table structure for table users
--
CREATE TABLE users (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  image_url varchar(255) NOT NULL,
  phone varchar(255) DEFAULT NULL,
  password text,
  PRIMARY KEY (id),
  UNIQUE KEY IDX_97672ac88f789774dd47f7c8be (email)
);

--
-- Table structure for table wishes
--
CREATE TABLE wishes (
  id int NOT NULL AUTO_INCREMENT,
  created_at datetime NOT NULL,
  user_id int DEFAULT NULL,
  participant_id int DEFAULT NULL,
  product_key int DEFAULT NULL,
  event_id int DEFAULT NULL,
  PRIMARY KEY (id),
  KEY IDX_4a6e5770133910acfc3c16a499 (user_id),
  KEY IDX_12156d1b005e5d22483fd4fc08 (participant_id),
  KEY IDX_95215d774cbe2079c9d90560dc (product_key),
  KEY IDX_a9ed3aaf623b52ec4d11e272f3 (event_id),
  CONSTRAINT FK_12156d1b005e5d22483fd4fc086 FOREIGN KEY (participant_id) REFERENCES participants (id) ON DELETE CASCADE,
  CONSTRAINT FK_4a6e5770133910acfc3c16a499b FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
  CONSTRAINT FK_95215d774cbe2079c9d90560dc5 FOREIGN KEY (product_key) REFERENCES products (id) ON DELETE CASCADE,
  CONSTRAINT FK_a9ed3aaf623b52ec4d11e272f3f FOREIGN KEY (event_id) REFERENCES events (id) ON DELETE CASCADE
);
