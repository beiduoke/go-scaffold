/*
 Navicat Premium Data Transfer

 Source Server         : 本地Docker数据库
 Source Server Type    : MySQL
 Source Server Version : 50735
 Source Host           : localhost:3306
 Source Schema         : go_scaffold

 Target Server Type    : MySQL
 Target Server Version : 50735
 File Encoding         : 65001

 Date: 08/12/2022 18:24:01
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_menus`;
CREATE TABLE `sys_menus`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `domain_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '领域ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由名称',
  `type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '菜单类型 0 无指定 1 目录 2 菜单 3 功能(按钮等)',
  `parent_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '父菜单ID',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由path',
  `hidden` tinyint(1) NOT NULL DEFAULT 1 COMMENT '隐藏 0 无指定 1 是 2 否',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '对应前端文件路径',
  `permission` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '权限标识',
  `sort` int(10) NOT NULL DEFAULT 10 COMMENT '排序标记',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '附加属性',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '附加属性',
  `keep_alive` tinyint(1) NOT NULL DEFAULT 1 COMMENT '附加属性',
  `base_menu` tinyint(1) NOT NULL DEFAULT 1 COMMENT '附加属性',
  `close_tab` tinyint(1) NOT NULL DEFAULT 1 COMMENT '附加属性',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_menu_type_name`(`name`, `type`) USING BTREE,
  INDEX `idx_sys_menus_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_domain_id_data`(`domain_id`) USING BTREE,
  INDEX `idx_sys_menus_parent_id`(`parent_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menus
-- ----------------------------
INSERT INTO `sys_menus` VALUES (1, '2022-12-08 03:03:34.028', '2022-12-08 07:36:41.721', NULL, 1, 'Dashboard', 1, 0, '/dashboard', 2, 'LAYOUT', '', 100, 'ion:grid-outline', '仪表盘', 1, 1, 1);
INSERT INTO `sys_menus` VALUES (2, '2022-12-08 03:06:12.140', '2022-12-08 07:36:57.116', NULL, 1, 'analysis', 2, 1, '/dashboard/analysis', 2, '/dashboard/analyse/index', 'dashboardAnalyse', 100, 'ant-design:alert-twotone', '分析页', 2, 2, 1);
INSERT INTO `sys_menus` VALUES (3, '2022-12-08 03:50:55.084', '2022-12-08 06:35:34.566', NULL, 1, 'add', 3, 2, '', 2, '', 'add', 100, 'ant-design:user-add-outlined', '新增', 1, 1, 1);
INSERT INTO `sys_menus` VALUES (5, '2022-12-08 05:56:56.234', '2022-12-08 06:35:43.458', NULL, 1, 'edit', 3, 2, '', 2, '', 'edit', 100, 'ant-design:edit-filled', '编辑', 1, 1, 1);
INSERT INTO `sys_menus` VALUES (6, '2022-12-08 06:18:26.477', '2022-12-08 06:35:51.860', NULL, 1, 'delete', 3, 2, '', 2, '', 'delete', 100, 'ant-design:delete-twotone', '删除', 1, 1, 1);
INSERT INTO `sys_menus` VALUES (7, '2022-12-08 06:19:15.292', '2022-12-08 06:36:05.088', NULL, 1, 'update', 3, 2, '', 2, '', 'update', 100, 'ant-design:upload-outlined', '修改', 1, 1, 1);
INSERT INTO `sys_menus` VALUES (8, '2022-12-08 07:40:29.095', '2022-12-08 07:40:29.095', NULL, 1, 'workbench', 2, 1, 'workbench', 2, '/dashboard/workbench/index', 'workbench', 100, '', '工作台', 2, 2, 1);
INSERT INTO `sys_menus` VALUES (9, '2022-12-08 07:41:41.753', '2022-12-08 07:41:41.753', NULL, 1, 'System', 1, 0, '/system', 2, 'LAYOUT', '', 100, 'ion:settings-outline', '系统管理', 1, 1, 1);
INSERT INTO `sys_menus` VALUES (10, '2022-12-08 07:45:05.125', '2022-12-08 07:45:05.125', NULL, 1, 'AccountManagement', 2, 9, 'account', 2, '/sys/account/index', 'account', 100, '', '账号管理', 2, 2, 1);
INSERT INTO `sys_menus` VALUES (11, '2022-12-08 07:47:50.582', '2022-12-08 07:47:50.582', NULL, 1, 'AccountDetail', 2, 10, 'account_detail/:id', 1, '/sys/account/AccountDetail', 'account_detail', 100, '', '账号详情', 2, 2, 1);
INSERT INTO `sys_menus` VALUES (12, '2022-12-08 07:48:47.697', '2022-12-08 07:48:47.697', NULL, 1, 'RoleManagement', 2, 9, 'role', 2, '/sys/role/index', 'role', 100, '', '角色管理', 2, 2, 1);
INSERT INTO `sys_menus` VALUES (13, '2022-12-08 07:49:18.253', '2022-12-08 07:49:18.253', NULL, 1, 'MenuManagement', 2, 9, 'menu', 2, '/sys/menu/index', 'menu', 100, '', '菜单管理', 2, 2, 1);
INSERT INTO `sys_menus` VALUES (14, '2022-12-08 07:49:43.133', '2022-12-08 07:49:43.133', NULL, 1, 'DeptManagement', 2, 9, 'dept', 2, '/sys/dept/index', 'dept', 100, '', '部门管理', 2, 2, 1);
INSERT INTO `sys_menus` VALUES (15, '2022-12-08 07:50:54.820', '2022-12-08 07:54:54.651', NULL, 1, 'ChangePassword', 2, 10, 'changePassword/:id', 1, '/sys/password/index', 'changePassword', 100, '', '修改密码', 2, 2, 1);

SET FOREIGN_KEY_CHECKS = 1;
