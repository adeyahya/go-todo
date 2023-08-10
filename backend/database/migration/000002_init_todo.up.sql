CREATE TABLE "todo" (
  id VARCHAR(12) PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  is_completed bool DEFAULT(false),
  completed_at TIMESTAMP,
  created_at TIMESTAMP,
  author_id VARCHAR(12) NOT NULL,
  CONSTRAINT fk_user
    FOREIGN KEY(author_id)
      REFERENCES "user"(id)
      ON DELETE SET NULL
);