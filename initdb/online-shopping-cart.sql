/*
 Navicat Premium Data Transfer

 Source Server         : docker-mysql
 Source Server Type    : MySQL
 Source Server Version : 110002 (11.0.2-MariaDB-1:11.0.2+maria~ubu2204)
 Source Host           : localhost:3306
 Source Schema         : online-shopping-cart

 Target Server Type    : MySQL
 Target Server Version : 110002 (11.0.2-MariaDB-1:11.0.2+maria~ubu2204)
 File Encoding         : 65001

 Date: 07/09/2023 11:31:21
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for accounts
-- ----------------------------
DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `phone_number` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'CUSTOMER',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'SYSTEM',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'SYSTEM',
  `deleted_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT current_timestamp(3),
  `updated_at` datetime(3) NULL DEFAULT current_timestamp(3),
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `email`(`email` ASC) USING BTREE,
  UNIQUE INDEX `phone_number`(`phone_number` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of accounts
-- ----------------------------
INSERT INTO `accounts` VALUES (1, 'Admin', 'admin@example.com', '08544432132', '$2a$05$WfPjV0pMYveDW/iP2AOtVu/xtgkfwY5IISEgjSaKbk5tdDgGUICFy', 'Address Example', 'ADMIN', 'SYSTEM', 'SYSTEM', NULL, '2023-09-06 12:42:53.279', '2023-09-06 12:42:53.279', NULL);
INSERT INTO `accounts` VALUES (2, 'Customer 1', 'customer1@example.com', '085445672341', '$2a$05$VXgglOjbT0cSTq7N.V7T8O2a9.gog0APjAt3Ohgep01JqZTycp3X6', 'Address Customer Example', 'CUSTOMER', 'SYSTEM', 'SYSTEM', NULL, '2023-09-06 12:45:35.579', '2023-09-06 12:45:35.579', NULL);

-- ----------------------------
-- Table structure for carts
-- ----------------------------
DROP TABLE IF EXISTS `carts`;
CREATE TABLE `carts`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `account_id` bigint UNSIGNED NOT NULL,
  `item_id` bigint UNSIGNED NOT NULL,
  `qty` bigint NOT NULL,
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'SYSTEM',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'SYSTEM',
  `deleted_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT current_timestamp(3),
  `updated_at` datetime(3) NULL DEFAULT current_timestamp(3),
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_account_item_key`(`account_id` ASC, `item_id` ASC) USING BTREE,
  INDEX `fk_items_cart`(`item_id` ASC) USING BTREE,
  CONSTRAINT `fk_accounts_cart` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE,
  CONSTRAINT `fk_carts_account` FOREIGN KEY (`account_id`) REFERENCES `items` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE,
  CONSTRAINT `fk_items_cart` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of carts
-- ----------------------------

-- ----------------------------
-- Table structure for items
-- ----------------------------
DROP TABLE IF EXISTS `items`;
CREATE TABLE `items`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `price` double NOT NULL,
  `qty` bigint NOT NULL,
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'SYSTEM',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'SYSTEM',
  `deleted_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT current_timestamp(3),
  `updated_at` datetime(3) NULL DEFAULT current_timestamp(3),
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of items
-- ----------------------------
INSERT INTO `items` VALUES (1, 'Baju', 'description baju', 200000, 61, 'admin@example.com', 'SYSTEM', NULL, '2023-09-06 13:32:14.313', '2023-09-07 00:53:27.811', NULL);
INSERT INTO `items` VALUES (2, 'Celana', 'description celana', 60000, 0, 'admin@example.com', 'admin@example.com', NULL, '2023-09-06 13:33:37.082', '2023-09-06 23:21:02.514', NULL);
INSERT INTO `items` VALUES (3, 'Sepatu', 'description sepatu', 540000, 10, 'admin@example.com', 'SYSTEM', NULL, '2023-09-06 13:35:12.721', '2023-09-06 13:35:12.721', NULL);
INSERT INTO `items` VALUES (4, 'Sandal', 'description sandal', 143000, 18, 'admin@example.com', 'SYSTEM', NULL, '2023-09-06 13:38:06.505', '2023-09-06 13:38:06.505', NULL);
INSERT INTO `items` VALUES (5, 'Hp', 'description hp', 1430000, 18, 'admin@example.com', 'SYSTEM', NULL, '2023-09-06 13:39:12.134', '2023-09-06 13:39:12.134', NULL);
INSERT INTO `items` VALUES (6, 'Daster', 'description daster', 130000, 189, 'admin@example.com', 'SYSTEM', NULL, '2023-09-06 13:44:23.379', '2023-09-06 13:44:23.379', NULL);

-- ----------------------------
-- Table structure for order_items
-- ----------------------------
DROP TABLE IF EXISTS `order_items`;
CREATE TABLE `order_items`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `order_id` bigint UNSIGNED NOT NULL,
  `item_id` bigint UNSIGNED NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `price` double NOT NULL,
  `qty` bigint NOT NULL,
  `total` double NOT NULL,
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'SYSTEM',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'SYSTEM',
  `deleted_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT current_timestamp(3),
  `updated_at` datetime(3) NULL DEFAULT current_timestamp(3),
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_orders_order_item`(`order_id` ASC) USING BTREE,
  INDEX `fk_items_order_item`(`item_id` ASC) USING BTREE,
  CONSTRAINT `fk_items_order_item` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE,
  CONSTRAINT `fk_orders_order_item` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order_items
-- ----------------------------
INSERT INTO `order_items` VALUES (1, 1, 1, 'Baju', 'description baju', 200000, 1, 200000, 'customer1@example.com', 'SYSTEM', NULL, '2023-09-06 17:43:37.892', '2023-09-06 17:43:37.892', NULL);
INSERT INTO `order_items` VALUES (3, 3, 1, 'Baju', 'description baju', 200000, 1, 200000, 'customer1@example.com', 'SYSTEM', NULL, '2023-09-06 17:50:15.945', '2023-09-06 17:50:15.945', NULL);
INSERT INTO `order_items` VALUES (4, 4, 1, 'Baju', 'description baju', 200000, 4, 800000, 'customer1@example.com', 'SYSTEM', NULL, '2023-09-06 17:53:27.808', '2023-09-06 17:53:27.808', NULL);

-- ----------------------------
-- Table structure for order_vouchers
-- ----------------------------
DROP TABLE IF EXISTS `order_vouchers`;
CREATE TABLE `order_vouchers`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `order_id` bigint UNSIGNED NOT NULL,
  `voucher_id` bigint UNSIGNED NOT NULL,
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `percentage` double NOT NULL,
  `max` double NOT NULL,
  `total` double NOT NULL,
  `applied` double NOT NULL,
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'SYSTEM',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'SYSTEM',
  `deleted_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT current_timestamp(3),
  `updated_at` datetime(3) NULL DEFAULT current_timestamp(3),
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_order_voucher_key`(`order_id` ASC, `voucher_id` ASC) USING BTREE,
  INDEX `fk_vouchers_order_voucher`(`voucher_id` ASC) USING BTREE,
  CONSTRAINT `fk_order_vouchers_order` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE,
  CONSTRAINT `fk_vouchers_order_voucher` FOREIGN KEY (`voucher_id`) REFERENCES `vouchers` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order_vouchers
-- ----------------------------
INSERT INTO `order_vouchers` VALUES (1, 3, 1, 'COBA', 'coba', 'oba aja', 100, 10000, 200000, 10000, 'customer1@example.com', 'SYSTEM', NULL, '2023-09-06 17:50:15.954', '2023-09-06 17:50:15.954', NULL);
INSERT INTO `order_vouchers` VALUES (2, 4, 1, 'COBA', 'coba', 'oba aja', 100, 10000, 800000, 10000, 'customer1@example.com', 'SYSTEM', NULL, '2023-09-06 17:53:27.814', '2023-09-06 17:53:27.814', NULL);

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `account_id` bigint UNSIGNED NOT NULL,
  `address` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `phone_number` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `total` double NOT NULL,
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'ORDERED',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'SYSTEM',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'SYSTEM',
  `deleted_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT current_timestamp(3),
  `updated_at` datetime(3) NULL DEFAULT current_timestamp(3),
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_orders_account`(`account_id` ASC) USING BTREE,
  CONSTRAINT `fk_orders_account` FOREIGN KEY (`account_id`) REFERENCES `items` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of orders
-- ----------------------------
INSERT INTO `orders` VALUES (1, 2, 'Address Customer Example', '085445672341', 200000, 'ORDERED', 'customer1@example.com', 'customer1@example.com', NULL, '2023-09-06 17:43:37.890', '2023-09-07 00:43:37.903', NULL);
INSERT INTO `orders` VALUES (3, 2, 'Address Customer Example', '085445672341', 190000, 'ORDERED', 'customer1@example.com', 'customer1@example.com', NULL, '2023-09-06 17:50:15.941', '2023-09-07 00:50:15.957', NULL);
INSERT INTO `orders` VALUES (4, 2, 'Address Customer Example', '085445672341', 790000, 'ORDERED', 'customer1@example.com', 'customer1@example.com', NULL, '2023-09-06 17:53:27.803', '2023-09-07 00:53:27.817', NULL);

-- ----------------------------
-- Table structure for vouchers
-- ----------------------------
DROP TABLE IF EXISTS `vouchers`;
CREATE TABLE `vouchers`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `percentage` double NOT NULL,
  `max` double NOT NULL,
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'SYSTEM',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'SYSTEM',
  `deleted_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT current_timestamp(3),
  `updated_at` datetime(3) NULL DEFAULT current_timestamp(3),
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `code`(`code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of vouchers
-- ----------------------------
INSERT INTO `vouchers` VALUES (1, 'COBA', 'coba', 'oba aja', 100, 10000, 'SYSTEM', 'SYSTEM', NULL, '2023-09-06 17:44:48.953', '2023-09-06 17:44:48.953', NULL);

SET FOREIGN_KEY_CHECKS = 1;
