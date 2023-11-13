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

 Date: 13/11/2023 09:48:25
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_dictionaries
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionaries`;
CREATE TABLE `sys_dictionaries`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at`  timestamp NOT NULL   DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL   ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
--   `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL  COMMENT '字典名（中）',
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL  COMMENT '字典名（英）',
  `status` tinyint(1)  NOT NULL  COMMENT '状态',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL  COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionaries_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dictionaries
-- ----------------------------
INSERT INTO `sys_dictionaries` VALUES (1, '2023-08-16 15:04:02.726', '2023-08-16 15:04:02.730', NULL, '性别', 'gender', 1, '性别字典');
INSERT INTO `sys_dictionaries` VALUES (2, '2023-08-16 15:04:02.726', '2023-08-16 15:04:02.745', NULL, '数据库int类型', 'int', 1, 'int类型对应的数据库类型');
INSERT INTO `sys_dictionaries` VALUES (3, '2023-08-16 15:04:02.726', '2023-08-16 15:04:02.752', NULL, '数据库时间日期类型', 'time.Time', 1, '数据库时间日期类型');
INSERT INTO `sys_dictionaries` VALUES (4, '2023-08-16 15:04:02.726', '2023-08-16 15:04:02.756', NULL, '数据库浮点型', 'float64', 1, '数据库浮点型');
INSERT INTO `sys_dictionaries` VALUES (5, '2023-08-16 15:04:02.726', '2023-08-16 15:04:02.760', NULL, '数据库字符串', 'string', 1, '数据库字符串');
INSERT INTO `sys_dictionaries` VALUES (6, '2023-08-16 15:04:02.726', '2023-08-16 15:04:02.765', NULL, '数据库bool类型', 'bool', 1, '数据库bool类型');
INSERT INTO `sys_dictionaries` VALUES (7, '2023-11-13 09:40:03.937', '2023-11-13 09:40:03.937', NULL, '测试', 'test', 1, '测试用的');

SET FOREIGN_KEY_CHECKS = 1;
