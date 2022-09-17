/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 100138
 Source Host           : localhost:3306
 Source Schema         : orders_by

 Target Server Type    : MySQL
 Target Server Version : 100138
 File Encoding         : 65001

 Date: 17/09/2022 23:58:02
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for items
-- ----------------------------
DROP TABLE IF EXISTS `items`;
CREATE TABLE `items`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `item_code` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `description` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `quantity` int(11) NOT NULL,
  `order_id` int(11) NOT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 66 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of items
-- ----------------------------
INSERT INTO `items` VALUES (62, 'A1', 'Iphone', 2, 28, NULL, '2022-09-17 23:51:35', '2022-09-17 23:51:35');
INSERT INTO `items` VALUES (63, 'A2', 'MacBook', 200, 29, NULL, '2022-09-17 23:53:07', '2022-09-17 23:53:07');
INSERT INTO `items` VALUES (64, 'A3', 'HP Samsung', 3, 29, NULL, '2022-09-17 23:53:07', '2022-09-17 23:53:07');
INSERT INTO `items` VALUES (65, 'A1', 'Iphone', 2, 29, NULL, '2022-09-17 23:53:07', '2022-09-17 23:53:07');

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_name` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `ordered_at` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `person` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_orders_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 30 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of orders
-- ----------------------------
INSERT INTO `orders` VALUES (28, 'Jarvis', '2022-09-17T21:21:46+00:00', '2022-09-17 23:51:35', '2022-09-17 23:51:35', NULL, NULL);
INSERT INTO `orders` VALUES (29, 'Jajang Mulyana', '2022-09-17T21:21:46+00:00', '2022-09-17 23:53:07', '2022-09-17 23:53:07', NULL, NULL);

SET FOREIGN_KEY_CHECKS = 1;
