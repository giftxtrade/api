--
-- Table structure for table categorie
--
CREATE TABLE "category" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "description" TEXT,
  "category_url" TEXT,

  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);

--
-- Table structure for table product
--
CREATE TYPE "currency_type" AS ENUM (
  'USD',
  'CAD'
);

CREATE TABLE "product" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY,
  "title" TEXT NOT NULL,
  "description" TEXT,
  "product_key" VARCHAR(255) UNIQUE NOT NULL,
  "image_url" TEXT NOT NULL,
  "total_reviews" INT NOT NULL,
  "rating" REAL NOT NULL,
  "price" MONEY NOT NULL,
  "currency" "currency_type" NOT NULL DEFAULT 'USD',
  "modified" TIMESTAMPTZ NOT NULL,
  "website" TEXT NOT NULL,
  "category_id" BIGINT REFERENCES "category"("id") ON DELETE SET NULL,

  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);

--
-- Table structure for table user
--
CREATE TABLE "user" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) UNIQUE NOT NULL,
  "image_url" VARCHAR(255) NOT NULL,
  "phone" VARCHAR(255),
  "admin" BOOLEAN NOT NULL DEFAULT false,
  "active" BOOLEAN NOT NULL DEFAULT false,

  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);

--
-- Table structure for table event
--
CREATE TABLE "event" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "description" TEXT,
  "budget" MONEY NOT NULL,
  "invitation_message" TEXT NOT NULL,
  "draw_at" TIMESTAMPTZ NOT NULL,
  "close_at" TIMESTAMPTZ NOT NULL,

  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);

--
-- Table structure for table link
--
CREATE TABLE "link" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY,
  "code" VARCHAR(255) NOT NULL,
  "expiration_date" TIMESTAMPTZ NOT NULL,
  "event_id" BIGINT REFERENCES "event"("id") ON DELETE CASCADE NOT NULL,

  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);

--
-- Table structure for table participant
--
CREATE TABLE "participant" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) NOT NULL,
  "address" VARCHAR(255) NOT NULL,
  "organizer" BOOLEAN NOT NULL DEFAULT false,
  "participates" BOOLEAN NOT NULL DEFAULT true,
  "accepted" BOOLEAN NOT NULL DEFAULT false,
  "event_id" BIGINT REFERENCES "event"("id") ON DELETE CASCADE NOT NULL,
  "user_id" BIGINT REFERENCES "user"("id") ON DELETE SET NULL,

  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);

--
-- Table structure for table draw
--
CREATE TABLE "draw" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY,
  "drawer_id" BIGINT REFERENCES "participant"("id") ON DELETE CASCADE NOT NULL,
  "drawee_id" BIGINT REFERENCES "participant"("id") ON DELETE CASCADE NOT NULL,
  "event_id" BIGINT REFERENCES "event"("id") NOT NULL,

  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);

--
-- Table structure for table wishes
--
CREATE TABLE "wish" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY,
  "user_id" BIGINT REFERENCES "user"("id") ON DELETE CASCADE NOT NULL,
  "participant_id" BIGINT REFERENCES "participant"("id") ON DELETE CASCADE NOT NULL,
  "product_id" BIGINT REFERENCES "product"("id") ON DELETE SET NULL,
  "event_id" BIGINT REFERENCES "event"("id") ON DELETE CASCADE NOT NULL,

  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);
