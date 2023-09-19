CREATE TABLE IF NOT EXISTS `users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `full_name` VARCHAR(45) NULL,
  `phone` VARCHAR(15) NULL,
  `gender` VARCHAR(15) NULL,
  `email` VARCHAR(100) NOT NULL,
  `password` VARCHAR(100) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  `role` VARCHAR(100) NOT NULL DEFAULT "Customer",
  PRIMARY KEY (id),
  CONSTRAINT `UQ_user_email` UNIQUE (`email`),
  CONSTRAINT `UQ_user_phone` UNIQUE (`phone`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
