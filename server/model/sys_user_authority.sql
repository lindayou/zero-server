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

 Date: 16/10/2023 10:28:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_user_authority
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_authority`;
CREATE TABLE `sys_user_authority`  (
  `sys_user_id` bigint UNSIGNED NOT NULL,
  `sys_authority_authority_id` bigint UNSIGNED NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_authority
-- ----------------------------
INSERT INTO `sys_user_authority` VALUES (1, 888);
INSERT INTO `sys_user_authority` VALUES (1, 8881);
INSERT INTO `sys_user_authority` VALUES (1, 9528);
INSERT INTO `sys_user_authority` VALUES (2, 888);

SET FOREIGN_KEY_CHECKS = 1;
