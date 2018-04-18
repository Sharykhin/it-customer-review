CREATE TABLE `reviews` (
  `id` CHAR(36) NOT NULL,
  `name` VARCHAR(80) NOT NULL,
  `email` VARCHAR(80) NOT NULL,
  `content` TEXT NULL,
  `published` TINYINT(1) NOT NULL DEFAULT 0,
  `score` INT(3) NOT NULL,
  `category` ENUM('positive', 'negative') NOT NULL,
  `created_at` TIMESTAMP NULL,
  `updated_at` TIMESTAMP NULL,
  PRIMARY KEY (`id`));