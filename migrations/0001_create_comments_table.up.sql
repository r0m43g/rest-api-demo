CREATE TABLE IF NOT EXISTS comments (
  ID uuid PRIMARY KEY,
  Slug text,
  Author text,
  Body text
);
