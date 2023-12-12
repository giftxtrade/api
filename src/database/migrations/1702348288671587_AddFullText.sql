ALTER TABLE product
ADD COLUMN product_ts tsvector
    GENERATED ALWAYS AS (
    	setweight(to_tsvector('english', coalesce(title, '')), 'A') ||
     	setweight(to_tsvector('english', coalesce(description, '')), 'B') 
    ) STORED;
