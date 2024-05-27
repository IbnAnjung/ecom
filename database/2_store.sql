CREATE DATABASE store;

use store;

-- shop.stores definition

CREATE TABLE `stores` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `seller_user_id` bigint NOT NULL,
  `name` varchar(50) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- shop.orders definition

CREATE TABLE `orders` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `store_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `total_price` double(14,2) DEFAULT NULL,
  `created_time` datetime NOT NULL COMMENT 'time created order',
  `expired_time` datetime DEFAULT NULL COMMENT 'time order to expire',
  `payment_time` datetime DEFAULT NULL COMMENT 'time of payment',
  `status` tinyint NOT NULL COMMENT '0: order created, 1:order paid, 9:expired/stock already back',
  PRIMARY KEY (`id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- shop.order_details definition

CREATE TABLE `order_details` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `order_id` bigint NOT NULL,
  `product_id` varchar(255) NOT NULL,
  `quantity` double(12,2) NOT NULL,
  `price` double(14,2) NOT NULL,
  `total_price` double(14,2) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- shop.warehouses definition

CREATE TABLE `warehouses` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `store_id` bigint NOT NULL,
  `name` varchar(50) NOT NULL,
  `status` tinyint NOT NULL DEFAULT '1',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_store_id` (`store_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- shop.warehouse_products definition

CREATE TABLE `warehouse_products` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `warehouse_id` bigint NOT NULL,
  `product_id` varchar(255) NOT NULL,
  `stock` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unq_warehouse_id_product_id` (`warehouse_id`,`product_id`),
  KEY `idx_product_id` (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;