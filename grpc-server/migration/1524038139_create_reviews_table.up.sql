CREATE TABLE `reviews` (
  `id` CHAR(36) NOT NULL,
  `name` VARCHAR(80) NOT NULL,
  `email` VARCHAR(80) NOT NULL,
  `content` TEXT NOT NULL,
  `published` TINYINT(1) NOT NULL DEFAULT 0,
  `score` INT(3),
  `category` ENUM('positive', 'negative'),
  `created_at` TIMESTAMP NULL,
  `updated_at` TIMESTAMP NULL,
  PRIMARY KEY (`id`));