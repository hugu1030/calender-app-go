-- +migrate Up
CREATE TABLE IF NOT EXISTS schedules (
  id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(255) NOT NULL,
  sub_title VARCHAR(255),
  place VARCHAR(255),
  from_date DATETIME NOT NULL,
  to_date DATETIME NOT NULL,
  updated_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +migrate Down
DROP TABLE IF EXISTS schedules;
