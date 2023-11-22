CREATE TABLE users (
  id  INT unsigned NOT NULL  AUTO_INCREMENT,
  first_name VARCHAR(255) NULL,
  last_name VARCHAR(255) NULL,
  email VARCHAR(255) NOT NULL,
  is_confirmed TINYINT(1) DEFAULT 0,
  accepted_tos TINYINT(1) DEFAULT 0,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now(),
  deleted_at TIMESTAMP NULL,
  PRIMARY KEY(id),
  UNIQUE (email)
);