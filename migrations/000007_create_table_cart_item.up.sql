CREATE TABLE `card_item` (
  `id` INT(11) PRIMARY KEY AUTO_INCREMENT,
  `cart_id` INT(11) NOT NULL,
  `product_id` INT(11) NOT NULL,
  `quantity` int NOT NULL,
  FOREIGN KEY (`cart_id`) REFERENCES `cart` (`id`),
  FOREIGN KEY (`product_id`) REFERENCES `product` (`id`)
);
