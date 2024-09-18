CREATE TABLE IF NOT EXISTS Urls(
  id BIGSERIAL PRIMARY KEY, 
  original_url VARCHAR(255) NOT NULL,
  shortened_url VARCHAR(255) NOT NULL UNIQUE,
  created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW()
);

CREATE INDEX shortened_url_idx ON Urls(shortened_url);
