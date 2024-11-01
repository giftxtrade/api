ALTER TABLE "product"
    RENAME COLUMN "website" TO "url";

ALTER TABLE "product"
    ADD COLUMN "origin" VARCHAR(50) NOT NULL;