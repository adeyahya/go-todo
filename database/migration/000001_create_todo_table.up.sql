CREATE TABLE todo (
  id VARCHAR(36) PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description string,
  is_completed bool DEFAULT(false)
);