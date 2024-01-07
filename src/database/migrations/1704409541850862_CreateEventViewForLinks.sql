CREATE VIEW event_link AS (
    SELECT
        "event".*,
        "link"."id" as "link_id",
        "link"."code" as "link_code",
        "link"."expiration_date" as "link_expiration_date"
    FROM "event"
    LEFT JOIN "link" ON "link"."event_id" = "event"."id"
);
