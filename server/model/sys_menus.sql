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

 Date: 10/10/2023 16:25:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_base_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menus`;
CREATE TABLE `sys_base_menus`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at`  timestamp NOT NULL   DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL   ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
--   `deleted_at` timestamp  NULL DEFAULT NULL,
  `menu_level` bigint UNSIGNED NOT NULL DEFAULT NULL,
  `parent_id` varchar(255) NOT NULL COMMENT '父菜单ID',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) NULL DEFAULT NOT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint NULL DEFAULT NOT NULL COMMENT '排序标记',
  `active_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT NULL COMMENT '附加属性',
  `keep_alive` tinyint(1) NULL DEFAULT NOT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) NULL DEFAULT NOT NULL COMMENT '附加属性',
  `title` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) NULL DEFAULT NOT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menus_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 32 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
