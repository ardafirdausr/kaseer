CREATE TABLE `users` (
  `id` int(11) AUTO_INCREMENT NOT NULL,
  `name` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `photo_url` varchar(250) DEFAULT NULL,
  `password` varchar(250) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `users` (`id`, `name`, `email`, `photo_url`, `password`, `created_at`, `updated_at`) VALUES
(1, 'Admin', 'admin@mail.com', NULL, '5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8', '2021-06-14 11:14:52', '2021-06-14 11:14:52'),
(2, 'Staff', 'staff@mail.com', '/static/image/shb.jpg', '2c6ba64c0fee083131cf80539097a9fcb04960f1', '2021-06-14 11:14:55', '2021-06-14 11:14:55');

CREATE TABLE `products` (
  `id` int(11) AUTO_INCREMENT NOT NULL,
  `code` varchar(15) NOT NULL,
  `name` varchar(50) NOT NULL,
  `price` int(11) NOT NULL DEFAULT 0,
  `stock` int(11) NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE `code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `products` (`id`, `code`, `name`, `price`, `stock`, `created_at`, `updated_at`) VALUES
(1, '001', 'Indomie Goreng', 3000, 90, '2021-06-14 11:15:06', '2021-06-14 11:15:06'),
(2, '002', 'Indomie Soto', 3300, 78, '2021-06-14 11:15:06', '2021-06-14 11:15:06'),
(3, '003', 'Gula 1KG', 12000, 19, '2021-06-14 11:15:06', '2021-06-14 11:15:06'),
(4, '004', 'Telor 1KG', 13000, 13, '2021-06-14 11:15:06', '2021-06-14 11:15:06'),
(5, '005', 'Beras 5KG', 54000, 10, '2021-06-14 11:15:06', '2021-06-14 11:15:06');


CREATE TABLE `orders` (
  `id` int(11) AUTO_INCREMENT NOT NULL,
  `total` int(11) NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `order_items` (
  `id` int(11) AUTO_INCREMENT NOT NULL,
  `order_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `quantity` int(11) NOT NULL DEFAULT 0,
  `subtotal` int(11) NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  FOREIGN KEY `fk_order_id` (`order_id`) REFERENCES `orders`(`id`),
  FOREIGN KEY `fk_product_id` (`product_id`) REFERENCES `products`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
