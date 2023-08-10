CREATE TABLE "user" (
  id VARCHAR(12) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(100) NOT NULL UNIQUE,
  password VARCHAR(72) NOT NULL,
  is_email_confirmed bool DEFAULT(false),
  email_confirmed_at TIMESTAMP,
  created_at TIMESTAMP
);