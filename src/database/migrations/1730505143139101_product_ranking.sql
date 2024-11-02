ALTER TABLE "product"
ADD COLUMN "ranking" BIGINT NOT NULL
    GENERATED ALWAYS AS (CEIL("total_reviews" * "rating")) STORED;
