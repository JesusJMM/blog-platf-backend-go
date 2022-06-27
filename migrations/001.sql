CREATE TABLE IF NOT EXISTS users (
  user_id SERIAL,
  name VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  img TEXT,
  PRIMARY KEY(user_id)
);

CREATE TABLE IF NOT EXISTS articles (
  article_id SERIAL,
  title VARCHAR(255) NOT NULL,
  description VARCHAR(255),
  content TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  slug VARCHAR(255) NOT NULL,
  sm_img TEXT,
  lg_img TEXT,
  user_id INT NOT NULL,
  PRIMARY KEY(article_id),
  FOREIGN KEY(user_id) REFERENCES users(user_id)
);
