CREATE TABLE todo (
  id VARCHAR(36) PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  is_completed bool DEFAULT(false),
  created_at DATETIME
);