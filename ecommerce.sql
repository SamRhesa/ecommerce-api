-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 03, 2026 at 12:55 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `ecommerce`
--

-- --------------------------------------------------------

--
-- Table structure for table `addresses`
--

CREATE TABLE `addresses` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL,
  `detail` longtext DEFAULT NULL,
  `city` longtext DEFAULT NULL,
  `province` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `addresses`
--

INSERT INTO `addresses` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `detail`, `city`, `province`) VALUES
(1, '2026-05-03 15:23:23.513', '2026-05-03 15:28:35.395', '2026-05-03 15:31:21.601', 6, 'Jl. Melati No. 20', 'Bandung', 'Jawa Barat'),
(2, '2026-05-03 16:27:21.805', '2026-05-03 16:27:21.805', NULL, 6, 'Jl. Mawar No. 10', 'Jakarta', 'DKI Jakarta');

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

CREATE TABLE `categories` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `categories`
--

INSERT INTO `categories` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`) VALUES
(1, '2026-05-03 15:50:56.920', '2026-05-03 15:57:34.813', '2026-05-03 15:57:48.096', 'Fashion'),
(2, '2026-05-03 17:22:31.021', '2026-05-03 17:22:31.021', NULL, 'Elektronik'),
(3, '2026-05-03 17:22:48.804', '2026-05-03 17:22:48.804', NULL, 'Makanan');

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `price` bigint(20) DEFAULT NULL,
  `stock` bigint(20) DEFAULT NULL,
  `store_id` bigint(20) UNSIGNED DEFAULT NULL,
  `category_id` bigint(20) UNSIGNED DEFAULT NULL,
  `image` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `price`, `stock`, `store_id`, `category_id`, `image`) VALUES
(1, '2026-05-03 16:03:30.486', '2026-05-03 17:31:15.433', NULL, 'Laptop', 10000000, 5, 2, 2, 'uploads/1777802697_foto.png'),
(2, '2026-05-03 16:04:01.754', '2026-05-03 17:33:01.343', NULL, 'smarthphone', 15000000, 5, 2, 2, ''),
(3, '2026-05-03 17:43:57.062', '2026-05-03 17:43:57.062', NULL, 'keyboard', 500000, 10, 2, 2, '');

-- --------------------------------------------------------

--
-- Table structure for table `stores`
--

CREATE TABLE `stores` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `stores`
--

INSERT INTO `stores` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `user_id`) VALUES
(1, '2026-05-01 13:53:11.189', '2026-05-01 13:53:11.189', NULL, 'siti\'s Store', 5),
(2, '2026-05-01 14:02:25.746', '2026-05-03 15:13:47.612', NULL, 'toko budi', 6);

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL,
  `address_id` bigint(20) UNSIGNED DEFAULT NULL,
  `total` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `address_id`, `total`) VALUES
(1, '2026-05-03 16:27:46.704', '2026-05-03 16:27:46.709', NULL, 6, 2, 40000000),
(2, '2026-05-03 16:29:16.335', '2026-05-03 16:29:16.348', NULL, 6, 2, 60000000);

-- --------------------------------------------------------

--
-- Table structure for table `transaction_items`
--

CREATE TABLE `transaction_items` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `transaction_id` bigint(20) UNSIGNED DEFAULT NULL,
  `product_id` bigint(20) UNSIGNED DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `price` bigint(20) DEFAULT NULL,
  `qty` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `transaction_items`
--

INSERT INTO `transaction_items` (`id`, `created_at`, `updated_at`, `deleted_at`, `transaction_id`, `product_id`, `name`, `price`, `qty`) VALUES
(1, '2026-05-03 16:27:46.706', '2026-05-03 16:27:46.706', NULL, 1, 1, 'Laptop', 10000000, 1),
(2, '2026-05-03 16:27:46.708', '2026-05-03 16:27:46.708', NULL, 1, 2, 'SmartPhone', 15000000, 2),
(3, '2026-05-03 16:29:16.345', '2026-05-03 16:29:16.345', NULL, 2, 2, 'SmartPhone', 15000000, 2),
(4, '2026-05-03 16:29:16.347', '2026-05-03 16:29:16.347', NULL, 2, 2, 'SmartPhone', 15000000, 2);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `email` varchar(191) DEFAULT NULL,
  `phone` varchar(191) DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `role` varchar(191) DEFAULT 'user'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `email`, `phone`, `password`, `role`) VALUES
(4, '2026-05-01 13:42:59.061', '2026-05-01 13:42:59.061', NULL, 'Andi', 'andi@mail.com', '08123455589', '$2a$10$O0tvFVTo6N.ybQYVou01/u9zwGcrGNI80eDLJjDTLtUwghdcuLjmq', 'user'),
(5, '2026-05-01 13:53:11.187', '2026-05-01 13:53:11.187', NULL, 'siti', 'siti@mail.com', '081656455589', '$2a$10$EUuJ1otIiFpc2sSgTGWgI.1xbtL31sibBonQn7E9cgIGjgZJpu2t2', 'user'),
(6, '2026-05-01 14:02:25.744', '2026-05-03 15:06:49.078', NULL, 'budiUpdate', 'budi@mail.com', '0899999999', '$2a$10$W/WyY6h65A0CvEzZNujceOIGKjNjAmmGn7MQJKonDiIgOs/QBTlVe', 'admin');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `addresses`
--
ALTER TABLE `addresses`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_addresses_deleted_at` (`deleted_at`);

--
-- Indexes for table `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_categories_deleted_at` (`deleted_at`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_products_deleted_at` (`deleted_at`);

--
-- Indexes for table `stores`
--
ALTER TABLE `stores`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_stores_deleted_at` (`deleted_at`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_transactions_deleted_at` (`deleted_at`);

--
-- Indexes for table `transaction_items`
--
ALTER TABLE `transaction_items`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_transaction_items_deleted_at` (`deleted_at`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_users_email` (`email`),
  ADD UNIQUE KEY `uni_users_phone` (`phone`),
  ADD KEY `idx_users_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `addresses`
--
ALTER TABLE `addresses`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `categories`
--
ALTER TABLE `categories`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `stores`
--
ALTER TABLE `stores`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `transaction_items`
--
ALTER TABLE `transaction_items`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
