CREATE TABLE `product` (
    `id` INT(11) PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `description` text,
    `price` float NOT NULL,
    `stock` int NOT NULL,
    `category_id` INT(11) NOT NULL,
    FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
);