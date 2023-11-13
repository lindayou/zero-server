/*
 Navicat Premium Data Transfer

 Source Server         : 本机mysql
 Source Server Type    : MySQL
 Source Server Version : 80034 (8.0.34)
 Source Host           : localhost:3306
 Source Schema         : gva

 Target Server Type    : MySQL
 Target Server Version : 80034 (8.0.34)
 File Encoding         : 65001

 Date: 13/11/2023 09:48:36
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_dictionary_details
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionary_details`;
CREATE TABLE `sys_dictionary_details`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at`  timestamp NOT NULL   DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL   ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
--   `deleted_at` datetime(3) NULL DEFAULT NULL,
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL  COMMENT '展示值',
  `value` bigint NOT NULL    COMMENT '字典值',
  `status` tinyint(1)  NOT NULL  COMMENT '启用状态',
  `sort` bigint  NOT NULL  COMMENT '排序标记',
  `sys_dictionary_id` bigint UNSIGNED  NOT NULL  COMMENT '关联标记',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionary_details_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dictionary_details
-- ----------------------------
INSERT INTO `sys_dictionary_details` VALUES (1, '2023-08-16 15:04:02.731', '2023-08-16 15:04:02.731', NULL, '男', 1, 1, 1, 1);
INSERT INTO `sys_dictionary_details` VALUES (2, '2023-08-16 15:04:02.731', '2023-08-16 15:04:02.731', NULL, '女', 2, 1, 2, 1);
INSERT INTO `sys_dictionary_details` VALUES (3, '2023-08-16 15:04:02.746', '2023-08-16 15:04:02.746', NULL, 'smallint', 1, 1, 1, 2);
INSERT INTO `sys_dictionary_details` VALUES (4, '2023-08-16 15:04:02.746', '2023-08-16 15:04:02.746', NULL, 'mediumint', 2, 1, 2, 2);
INSERT INTO `sys_dictionary_details` VALUES (5, '2023-08-16 15:04:02.746', '2023-08-16 15:04:02.746', NULL, 'int', 3, 1, 3, 2);
INSERT INTO `sys_dictionary_details` VALUES (6, '2023-08-16 15:04:02.746', '2023-08-16 15:04:02.746', NULL, 'bigint', 4, 1, 4, 2);
INSERT INTO `sys_dictionary_details` VALUES (7, '2023-08-16 15:04:02.753', '2023-08-16 15:04:02.753', NULL, 'date', 0, 1, 0, 3);
INSERT INTO `sys_dictionary_details` VALUES (8, '2023-08-16 15:04:02.753', '2023-08-16 15:04:02.753', NULL, 'time', 1, 1, 1, 3);
INSERT INTO `sys_dictionary_details` VALUES (9, '2023-08-16 15:04:02.753', '2023-08-16 15:04:02.753', NULL, 'year', 2, 1, 2, 3);
INSERT INTO `sys_dictionary_details` VALUES (10, '2023-08-16 15:04:02.753', '2023-08-16 15:04:02.753', NULL, 'datetime', 3, 1, 3, 3);
INSERT INTO `sys_dictionary_details` VALUES (11, '2023-08-16 15:04:02.753', '2023-08-16 15:04:02.753', NULL, 'timestamp', 5, 1, 5, 3);
INSERT INTO `sys_dictionary_details` VALUES (12, '2023-08-16 15:04:02.757', '2023-08-16 15:04:02.757', NULL, 'float', 0, 1, 0, 4);
INSERT INTO `sys_dictionary_details` VALUES (13, '2023-08-16 15:04:02.757', '2023-08-16 15:04:02.757', NULL, 'double', 1, 1, 1, 4);
INSERT INTO `sys_dictionary_details` VALUES (14, '2023-08-16 15:04:02.757', '2023-08-16 15:04:02.757', NULL, 'decimal', 2, 1, 2, 4);
INSERT INTO `sys_dictionary_details` VALUES (15, '2023-08-16 15:04:02.761', '2023-08-16 15:04:02.761', NULL, 'char', 0, 1, 0, 5);
INSERT INTO `sys_dictionary_details` VALUES (16, '2023-08-16 15:04:02.761', '2023-08-16 15:04:02.761', NULL, 'varchar', 1, 1, 1, 5);
INSERT INTO `sys_dictionary_details` VALUES (17, '2023-08-16 15:04:02.761', '2023-08-16 15:04:02.761', NULL, 'tinyblob', 2, 1, 2, 5);
INSERT INTO `sys_dictionary_details` VALUES (18, '2023-08-16 15:04:02.761', '2023-08-16 15:04:02.761', NULL, 'tinytext', 3, 1, 3, 5);
INSERT INTO `sys_dictionary_details` VALUES (19, '2023-08-16 15:04:02.761', '2023-08-16 15:04:02.761', NULL, 'text', 4, 1, 4, 5);
INSERT INTO `sys_dictionary_details` VALUES (20, '2023-08-16 15:04:02.761', '2023-08-16 15:04:02.761', NULL, 'blob', 5, 1, 5, 5);
INSERT INTO `sys_dictionary_details` VALUES (21, '2023-08-16 15:04:02.761', '2023-08-16 15:04:02.761', NULL, 'mediumblob', 6, 1, 6, 5);
INSERT INTO `sys_dictionary_details` VALUES (22, '2023-08-16 15:04:02.761', '2023-08-16 15:04:02.761', NULL, 'mediumtext', 7, 1, 7, 5);
INSERT INTO `sys_dictionary_details` VALUES (23, '2023-08-16 15:04:02.761', '2023-08-16 15:04:02.761', NULL, 'longblob', 8, 1, 8, 5);
INSERT INTO `sys_dictionary_details` VALUES (24, '2023-08-16 15:04:02.761', '2023-08-16 15:04:02.761', NULL, 'longtext', 9, 1, 9, 5);
INSERT INTO `sys_dictionary_details` VALUES (25, '2023-08-16 15:04:02.765', '2023-08-16 15:04:02.765', NULL, 'tinyint', 0, 1, 0, 6);
INSERT INTO `sys_dictionary_details` VALUES (26, '2023-11-13 09:40:22.621', '2023-11-13 09:40:22.621', NULL, '测试1', 1, 1, 1, 7);
INSERT INTO `sys_dictionary_details` VALUES (27, '2023-11-13 09:40:31.885', '2023-11-13 09:40:31.885', NULL, '测试2', 2, 1, 2, 7);

SET FOREIGN_KEY_CHECKS = 1;
