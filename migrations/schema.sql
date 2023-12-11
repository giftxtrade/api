--
-- Table structure for table categories
--
CREATE TABLE category (
  id BIGSERIAL UNIQUE PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  category_url VARCHAR(255) NOT NULL
);

--
-- Table structure for table draws
--
CREATE TABLE draw (
  id BIGSERIAL UNIQUE PRIMARY KEY,
  created_at DATETIME NOT NULL,
  drawer_id BIGINT REFERENCES participant(id) NOT NULL,
  drawee_id BIGINT REFERENCES participant(id) NOT NULL,
  event_id BIGINT REFERENCES event(id) NOT NULL
);

--
-- Table structure for table events
--
CREATE TABLE event (
  id BIGSERIAL UNIQUE PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  budget DECIMAL(10,0) NOT NULL,
  invitation_message TEXT NOT NULL,
  created_at DATETIME NOT NULL,
  draw_at DATETIME NOT NULL,
  close_at DATETIME NOT NULL
);

--
-- Table structure for table links
--
CREATE TABLE link (
  id BIGSERIAL UNIQUE PRIMARY KEY,
  code VARCHAR(255) NOT NULL,
  created_at DATETIME NOT NULL,
  expiration_date DATETIME NOT NULL,
  event_id BIGINT REFERENCES event(id) NOT NULL
);

--
-- Table structure for table participants
--
CREATE TABLE participant (
  id BIGSERIAL UNIQUE PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  address VARCHAR(255) NOT NULL,
  organizer BOOLEAN DEFAULT false,
  participates BOOLEAN DEFAULT true,
  accepted BOOLEAN DEFAULT false,
  event_id BIGINT REFERENCES event(id) NOT NULL,
  user_id BIGINT REFERNCES user(id) NOT NULL
);

--
-- Table structure for table products
--
CREATE TABLE product (
  id BIGSERIAL UNIQUE PRIMARY KEY,
  title TEXT NOT NULL,
  description TEXT,
  product_key VARCHAR(255) UNIQUE NOT NULL,
  image_url TEXT NOT NULL,
  total_reviews INT NOT NULL,
  rating DOUBLE NOT NULL,
  price DOUBLE NOT NULL,
  currency VARCHAR(255) NOT NULL,
  modified DATETIME NOT NULL,
  website TEXT NOT NULL,
  category_id BIGINT REFERENCES category(id) NOT NULL
);

ALTER TABLE product
ADD COLUMN title_ts tsvector
    GENERATED ALWAYS AS (
    	setweight(to_tsvector('english', coalesce(title, '')), 'A') ||
     	setweight(to_tsvector('english', coalesce(description, '')), 'B') 
    ) STORED;

--
-- Table structure for table users
--
CREATE TABLE user (
  id BIGSERIAL UNIQUE PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  image_url VARCHAR(255) NOT NULL,
  phone VARCHAR(255) DEFAULT NULL,
  admin BOOLEAN DEFAULT false
);

--
-- Table structure for table wishes
--
CREATE TABLE wish (
  id BIGSERIAL UNIQUE PRIMARY KEY,
  created_at DATETIME NOT NULL,
  user_id BIGINT REFERENCES user(id) NOT NULL,
  participant_id BIGINT REFERENCES participant(id) NOT NULL,
  product_id BIGINT REFERENCES product(id) NOT NULL,
  event_id BIGINT REFERENCES event(id) NOT NULL
);
