CREATE TABLE `profiles` (
    `user_id` VARCHAR(255) NOT NULL,
    `nick_name` VARCHAR(255) NOT NULL,
    `date_of_birth` DATE NOT NULL,
    `fishing_started_date` DATE NOT NULL,
    `image` VARCHAR(255) DEFAULT NULL,
    `created_at` TIMESTAMP NULL DEFAULT NULL,
    `updated_at` TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
