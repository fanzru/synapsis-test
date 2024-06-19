CREATE TABLE `order` (
    `id` INT(11) PRIMARY KEY AUTO_INCREMENT,
    `user_id` INT(11) NOT NULL,
    `total_amount` float NOT NULL,
    `status` varchar(255) NOT NULL,
    `created_at` timestamp DEFAULT (current_timestamp),
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
);