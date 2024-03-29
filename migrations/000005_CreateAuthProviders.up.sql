CREATE TABLE auth_providers (
  provider VARCHAR(255) NOT NULL,
  user_id INT unsigned NOT NULL,
  uid VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now(),

  FOREIGN KEY (user_id) REFERENCES `users`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  UNIQUE INDEX `auth_providers_idx` (`provider`,`user_id`)
);