UPDATE articles SET published=false;
ALTER TABLE articles ALTER COLUMN published SET DEFAULT false;
ALTER TABLE articles ALTER COLUMN published SET NOT NULL;
