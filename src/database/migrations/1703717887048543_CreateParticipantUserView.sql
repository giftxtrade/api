CREATE VIEW participant_user AS (
    SELECT 
        "participant".*,
        "user"."name" AS "user_name",
        "user"."email" AS "user_email",
        "user"."image_url" AS "user_image_url"
    FROM "participant"
    LEFT JOIN "user" ON "user"."id" = "participant"."user_id"
);
