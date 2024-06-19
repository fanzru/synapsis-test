CREATE TABLE `order_item` (
    `id` INT(11) PRIMARY KEY AUTO_INCREMENT,
    `order_id` INT(11) NOT NULL,
    `product_id` INT(11) NOT NULL,
    `quantity` int NOT NULL,
    `price` float NOT NULL,
    FOREIGN KEY (`order_id`) REFERENCES `order` (`id`),
    FOREIGN KEY (`product_id`) REFERENCES `product` (`id`)
);