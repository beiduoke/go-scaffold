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

 Date: 12/01/2023 11:59:58
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_api_operation_logs
-- ----------------------------
DROP TABLE IF EXISTS `sys_api_operation_logs`;
CREATE TABLE `sys_api_operation_logs`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `domain_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '领域ID',
  `ip` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求ip',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求方法',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求路径',
  `status` int(10) NOT NULL COMMENT '请求状态',
  `latency` int(10) NOT NULL DEFAULT 0 COMMENT '延迟',
  `agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '代理',
  `error` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '错误信息',
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求Body',
  `resp` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '响应Body',
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_domain_id_data`(`domain_id`) USING BTREE,
  INDEX `idx_sys_api_operation_logs_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_casbin_rules
-- ----------------------------
DROP TABLE IF EXISTS `sys_casbin_rules`;
CREATE TABLE `sys_casbin_rules`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_sys_casbin_rules`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_casbin_rules
-- ----------------------------
INSERT INTO `sys_casbin_rules` VALUES (1, 'g', '1', '1', '0', '1', '', '');
INSERT INTO `sys_casbin_rules` VALUES (13, 'g', '2', '1', '1', '0', '', '');
INSERT INTO `sys_casbin_rules` VALUES (12, 'p', '1', '1', '/api.admin.v1.Admin/ListApi', '*', '', '');

-- ----------------------------
-- Table structure for sys_depts
-- ----------------------------
DROP TABLE IF EXISTS `sys_depts`;
CREATE TABLE `sys_depts`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `domain_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '领域ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `parent_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '父角色ID',
  `sort` int(10) NOT NULL DEFAULT 100 COMMENT '排序',
  `remarks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `state` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态 0 未指定  1 启用 2 停用',
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_depts_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_domain_id_data`(`domain_id`) USING BTREE,
  INDEX `idx_sys_depts_state`(`state`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_depts
-- ----------------------------
INSERT INTO `sys_depts` VALUES (1, '2022-12-08 06:40:02.892', '2022-12-08 06:40:02.892', NULL, 0, '产品部', 0, 100, '', 1, 0);

-- ----------------------------
-- Table structure for sys_domain_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_domain_menus`;
CREATE TABLE `sys_domain_menus`  (
  `sys_menu_id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `sys_domain_id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
  PRIMARY KEY (`sys_menu_id`, `sys_domain_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_domain_menus
-- ----------------------------
INSERT INTO `sys_domain_menus` VALUES (1, 1);
INSERT INTO `sys_domain_menus` VALUES (1, 2);
INSERT INTO `sys_domain_menus` VALUES (2, 1);
INSERT INTO `sys_domain_menus` VALUES (2, 2);
INSERT INTO `sys_domain_menus` VALUES (3, 1);
INSERT INTO `sys_domain_menus` VALUES (3, 2);
INSERT INTO `sys_domain_menus` VALUES (5, 1);
INSERT INTO `sys_domain_menus` VALUES (5, 2);
INSERT INTO `sys_domain_menus` VALUES (6, 1);
INSERT INTO `sys_domain_menus` VALUES (6, 2);
INSERT INTO `sys_domain_menus` VALUES (7, 1);
INSERT INTO `sys_domain_menus` VALUES (7, 2);
INSERT INTO `sys_domain_menus` VALUES (8, 1);
INSERT INTO `sys_domain_menus` VALUES (8, 2);
INSERT INTO `sys_domain_menus` VALUES (9, 1);
INSERT INTO `sys_domain_menus` VALUES (10, 1);
INSERT INTO `sys_domain_menus` VALUES (10, 2);
INSERT INTO `sys_domain_menus` VALUES (12, 1);
INSERT INTO `sys_domain_menus` VALUES (12, 2);
INSERT INTO `sys_domain_menus` VALUES (14, 1);
INSERT INTO `sys_domain_menus` VALUES (14, 2);
INSERT INTO `sys_domain_menus` VALUES (16, 2);

-- ----------------------------
-- Table structure for sys_domains
-- ----------------------------
DROP TABLE IF EXISTS `sys_domains`;
CREATE TABLE `sys_domains`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '领域编码',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `parent_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父角色ID',
  `sort` int(10) NOT NULL DEFAULT 100 COMMENT '排序',
  `state` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态 0 未指定  1 启用 2 停用',
  `default_role_id` bigint(20) NOT NULL COMMENT '默认角色',
  `remarks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_domains_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_sys_domains_code`(`code`) USING BTREE,
  INDEX `idx_sys_domains_state`(`state`) USING BTREE,
  INDEX `idx_sys_domains_default_role_id`(`default_role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_domains
-- ----------------------------
INSERT INTO `sys_domains` VALUES (0, '2023-01-04 17:39:54.000', '2023-01-04 17:39:56.000', NULL, '1557658909480354584', '中台管理平台', 0, 100, 1, 1, '');
INSERT INTO `sys_domains` VALUES (1, '2023-01-04 17:39:54.000', '2023-01-04 17:39:56.000', NULL, '1557658909480354583', '中台运营平台', 1, 100, 1, 1, '');

-- ----------------------------
-- Table structure for sys_menu_buttons
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu_buttons`;
CREATE TABLE `sys_menu_buttons`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `menu_id` bigint(20) UNSIGNED NOT NULL COMMENT '菜单ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '按钮关键key',
  `remarks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '按钮备注',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_menu_buttons_menu_id`(`menu_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_menu_parameters
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu_parameters`;
CREATE TABLE `sys_menu_parameters`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `menu_id` bigint(20) UNSIGNED NOT NULL COMMENT '菜单ID',
  `type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '地址栏携带参类型 0 未指定 1 params 2 query',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地址栏携带参数的名称',
  `value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_menu_parameters_menu_id`(`menu_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_menu_resources
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu_resources`;
CREATE TABLE `sys_menu_resources`  (
  `sys_menu_id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `sys_resource_id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
  PRIMARY KEY (`sys_menu_id`, `sys_resource_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_menus`;
CREATE TABLE `sys_menus`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
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
  `ext_type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '附加属性',
  `remarks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_menu_type_name`(`name`, `type`) USING BTREE,
  INDEX `idx_sys_menus_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_sys_menus_parent_id`(`parent_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menus
-- ----------------------------
INSERT INTO `sys_menus` VALUES (1, '2022-12-08 03:03:34.028', '2022-12-08 07:36:41.721', NULL, 'Dashboard', 1, 0, '/dashboard', 2, 'LAYOUT', '', 100, 'ion:grid-outline', '仪表盘', 1, 1, 1, 1, '');
INSERT INTO `sys_menus` VALUES (2, '2022-12-08 03:06:12.140', '2023-01-11 03:27:35.110', NULL, 'Analysis', 2, 1, 'analysis', 2, '/dashboard/analysis/index', 'dashboardAnalyse', 100, 'ant-design:alert-twotone', '分析页', 2, 2, 1, 1, '');
INSERT INTO `sys_menus` VALUES (3, '2022-12-08 03:50:55.084', '2022-12-08 06:35:34.566', NULL, 'Add', 3, 2, '', 2, '', 'add', 100, 'ant-design:user-add-outlined', '新增', 1, 1, 1, 1, '');
INSERT INTO `sys_menus` VALUES (5, '2022-12-08 05:56:56.234', '2022-12-08 06:35:43.458', NULL, 'Edit', 3, 2, '', 2, '', 'edit', 100, 'ant-design:edit-filled', '编辑', 1, 1, 1, 1, '');
INSERT INTO `sys_menus` VALUES (6, '2022-12-08 06:18:26.477', '2022-12-08 06:35:51.860', NULL, 'Delete', 3, 2, '', 2, '', 'delete', 100, 'ant-design:delete-twotone', '删除', 1, 1, 1, 1, '');
INSERT INTO `sys_menus` VALUES (7, '2022-12-08 06:19:15.292', '2022-12-08 06:36:05.088', NULL, 'Update', 3, 2, '', 2, '', 'update', 100, 'ant-design:upload-outlined', '修改', 1, 1, 1, 1, '');
INSERT INTO `sys_menus` VALUES (8, '2022-12-08 07:40:29.095', '2022-12-08 07:40:29.095', NULL, 'Workbench', 2, 1, 'workbench', 2, '/dashboard/workbench/index', 'workbench', 100, '', '工作台', 2, 2, 1, 1, '');
INSERT INTO `sys_menus` VALUES (9, '2022-12-08 07:41:41.753', '2022-12-08 07:41:41.753', NULL, 'System', 1, 0, '/system', 2, 'LAYOUT', '', 100, 'ion:settings-outline', '系统管理', 1, 1, 1, 1, '');
INSERT INTO `sys_menus` VALUES (10, '2022-12-08 07:45:05.125', '2022-12-08 07:45:05.125', NULL, 'AccountManagement', 2, 9, 'account', 2, '/sys/account/index', 'account', 100, '', '账号管理', 2, 2, 1, 1, '');
INSERT INTO `sys_menus` VALUES (11, '2022-12-08 07:47:50.582', '2022-12-08 07:47:50.582', NULL, 'AccountDetail', 2, 10, 'account_detail/:id', 1, '/sys/account/AccountDetail', 'account_detail', 100, '', '账号详情', 2, 2, 1, 1, '');
INSERT INTO `sys_menus` VALUES (12, '2022-12-08 07:48:47.697', '2022-12-08 07:48:47.697', NULL, 'RoleManagement', 2, 9, 'role', 2, '/sys/role/index', 'role', 100, '', '角色管理', 2, 2, 1, 1, '');
INSERT INTO `sys_menus` VALUES (13, '2022-12-08 07:49:18.253', '2023-01-11 03:26:09.922', NULL, 'MenuManagement', 2, 17, 'menu', 2, '/sys/menu/index', 'menu', 100, '', '菜单管理', 2, 2, 1, 1, '');
INSERT INTO `sys_menus` VALUES (14, '2022-12-08 07:49:43.133', '2022-12-08 07:49:43.133', NULL, 'DeptManagement', 2, 9, 'dept', 2, '/sys/dept/index', 'dept', 100, '', '部门管理', 2, 2, 1, 1, '');
INSERT INTO `sys_menus` VALUES (15, '2022-12-08 07:50:54.820', '2022-12-08 07:54:54.651', NULL, 'ChangePassword', 2, 10, 'changePassword/:id', 1, '/sys/password/index', 'changePassword', 100, '', '修改密码', 2, 2, 1, 1, '');
INSERT INTO `sys_menus` VALUES (16, '2023-01-10 02:12:34.805', '2023-01-10 02:12:34.805', NULL, 'Post', 2, 9, 'post', 2, '/sys/post/index', 'post', 100, '', '职位管理', 2, 2, 1, 1, '');
INSERT INTO `sys_menus` VALUES (17, '2023-01-11 03:22:07.673', '2023-01-11 03:22:07.673', NULL, 'Middleground', 1, 0, '/middleground', 2, 'LAYOUT', '', 100, 'ant-design:setting-outlined', '中台管理', 1, 1, 1, 1, '');
INSERT INTO `sys_menus` VALUES (18, '2023-01-11 03:24:59.502', '2023-01-11 03:24:59.502', NULL, 'Domain', 2, 17, 'domain', 2, '/sys/domain/index', 'domain', 100, '', '租户管理', 2, 2, 1, 1, '');
INSERT INTO `sys_menus` VALUES (19, '2023-01-11 03:25:50.886', '2023-01-11 03:25:50.886', NULL, 'Resource', 2, 17, 'resource', 2, '/sys/resource/index', 'resource', 100, '', '资源管理', 2, 2, 1, 1, '');

-- ----------------------------
-- Table structure for sys_posts
-- ----------------------------
DROP TABLE IF EXISTS `sys_posts`;
CREATE TABLE `sys_posts`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `domain_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '领域ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '职位名称',
  `sort` int(10) NOT NULL DEFAULT 100 COMMENT '排序',
  `state` tinyint(1) NOT NULL DEFAULT 1 COMMENT '职位状态 0 未指定  1 启用 2 停用',
  `remarks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '职位编码',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_posts_state`(`state`) USING BTREE,
  INDEX `idx_sys_posts_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_domain_id_data`(`domain_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_posts
-- ----------------------------
INSERT INTO `sys_posts` VALUES (1, '2023-01-10 02:21:41.604', '2023-01-10 07:07:40.646', NULL, 0, '董事长', 100, 1, '', '');
INSERT INTO `sys_posts` VALUES (2, '2023-01-10 07:26:19.430', '2023-01-10 07:26:20.292', NULL, 0, '总经理', 99, 1, '总经理职位', '1612708008079396864');

-- ----------------------------
-- Table structure for sys_resources
-- ----------------------------
DROP TABLE IF EXISTS `sys_resources`;
CREATE TABLE `sys_resources`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求路径',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'POST' COMMENT '方法',
  `operation` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求动作',
  `group` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'api分组',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'api描述',
  `remarks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_api_path_method`(`path`, `method`, `operation`) USING BTREE,
  INDEX `idx_sys_resources_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_resources
-- ----------------------------
INSERT INTO `sys_resources` VALUES (1, '2022-11-05 07:38:29.503', '2022-12-24 07:03:25.647', NULL, '资源列表', '/api.admin.v1.Admin/ListResource', '*', '/v1/admin/resources', '默认', '', '');
INSERT INTO `sys_resources` VALUES (2, '2022-12-24 07:02:54.516', '2022-12-24 07:30:14.589', NULL, '菜单列表', '/api.admin.v1.Admin/ListMenu', '*', '/v1/admin/menus', '默认', '', '');
INSERT INTO `sys_resources` VALUES (3, '2022-12-24 07:32:38.919', '2022-12-24 07:33:01.874', NULL, '领域列表', '/api.admin.v1.Admin/ListDomain', '*', '/v1/admin/domains', '领域', '', '');

-- ----------------------------
-- Table structure for sys_role_depts
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_depts`;
CREATE TABLE `sys_role_depts`  (
  `sys_dept_id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `sys_role_id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
  PRIMARY KEY (`sys_dept_id`, `sys_role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_role_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menus`;
CREATE TABLE `sys_role_menus`  (
  `created_at` datetime(3) NULL DEFAULT NULL,
  `sys_role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `sys_menu_id` bigint(20) UNSIGNED NOT NULL COMMENT '菜单ID',
  `sys_menu_button` json NULL COMMENT '菜单按钮',
  `sys_menu_parameter` json NULL COMMENT '菜单参数',
  UNIQUE INDEX `idx_role_menu_role_id_menu_id`(`sys_role_id`, `sys_menu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_menus
-- ----------------------------
INSERT INTO `sys_role_menus` VALUES ('2023-01-10 02:13:09.581', 1, 1, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2023-01-10 02:13:09.581', 1, 2, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2023-01-10 02:13:09.581', 1, 3, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2023-01-10 02:13:09.581', 1, 5, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2023-01-10 02:13:09.581', 1, 6, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2023-01-10 02:13:09.581', 1, 7, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2023-01-10 02:13:09.581', 1, 8, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2023-01-10 02:13:09.581', 1, 9, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2023-01-10 02:13:09.581', 1, 10, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2023-01-10 02:13:09.581', 1, 12, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2023-01-10 02:13:09.581', 1, 13, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2023-01-10 02:13:09.581', 1, 14, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2023-01-10 02:13:09.581', 1, 16, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2023-01-10 02:13:09.581', 1, 17, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2023-01-10 02:13:09.581', 1, 18, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2022-12-23 02:05:39.053', 2, 3, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2022-12-23 02:05:39.053', 2, 6, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2022-12-23 02:05:39.053', 2, 7, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2022-12-23 02:05:39.053', 2, 8, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2022-12-23 02:05:39.053', 2, 9, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2022-12-23 02:05:39.053', 2, 10, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2022-12-23 02:05:39.053', 2, 12, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2022-12-23 02:05:39.053', 2, 13, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2022-12-23 02:05:39.053', 2, 14, '[]', '[]');
INSERT INTO `sys_role_menus` VALUES ('2022-12-23 02:05:39.053', 2, 16, '[]', '[]');

-- ----------------------------
-- Table structure for sys_role_relations
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_relations`;
CREATE TABLE `sys_role_relations`  (
  `sys_role_id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `role_id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
  PRIMARY KEY (`sys_role_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_role_resources
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_resources`;
CREATE TABLE `sys_role_resources`  (
  `sys_role_id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `sys_resource_id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
  PRIMARY KEY (`sys_role_id`, `sys_resource_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_roles
-- ----------------------------
DROP TABLE IF EXISTS `sys_roles`;
CREATE TABLE `sys_roles`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `domain_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '领域ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `parent_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父角色ID',
  `default_router` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '/dashboard' COMMENT '默认路由',
  `sort` int(10) NOT NULL DEFAULT 100 COMMENT '排序',
  `state` tinyint(1) NOT NULL DEFAULT 1 COMMENT '角色状态 0 未指定  1 启用 2 停用',
  `remarks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_roles_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_domain_id_data`(`domain_id`, `name`) USING BTREE,
  INDEX `idx_sys_roles_state`(`state`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_roles
-- ----------------------------
INSERT INTO `sys_roles` VALUES (1, '2022-11-05 06:21:06.260', '2023-01-11 03:17:04.761', NULL, 0, '超级管理员', 0, '', 0, 1, '');
INSERT INTO `sys_roles` VALUES (2, '2022-11-05 09:07:56.985', '2022-12-24 01:44:20.580', NULL, 0, 'default', 0, '', 0, 1, '');

-- ----------------------------
-- Table structure for sys_user_posts
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_posts`;
CREATE TABLE `sys_user_posts`  (
  `sys_post_id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `sys_user_id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
  PRIMARY KEY (`sys_post_id`, `sys_user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `nick_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `real_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '实名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `birthday` datetime(0) NULL DEFAULT NULL COMMENT '生日',
  `gender` tinyint(1) NOT NULL DEFAULT 1 COMMENT '性别 0 未指定 1 男 2 女',
  `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `state` tinyint(1) NOT NULL DEFAULT 1 COMMENT '用户状态 0 未指定  1 启用 2 停用',
  `last_use_domain` bigint(20) NOT NULL DEFAULT 0 COMMENT '最后使用租户',
  `last_login_at` datetime(0) NULL DEFAULT NULL COMMENT '最后登录时间',
  `last_login_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `remarks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_users_mobile_email`(`mobile`, `email`) USING BTREE,
  INDEX `idx_sys_users_state`(`state`) USING BTREE,
  INDEX `idx_sys_users_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_users_name_nick_name_real_name`(`name`, `real_name`, `nick_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES (1, '2022-11-05 06:17:21.364', '2022-11-05 06:17:21.364', NULL, '', 'xiong', 'xiong', '熊龙军', '$2a$10$rxZ7O60l2fN1DJuXBNKWgub50t320mN4.7innl5b6Gv34XQkU/IZa', NULL, 1, '18584565115', 'xiong@qq.com', 1, 0, NULL, '127.0.0.1', '');
INSERT INTO `sys_users` VALUES (2, '2022-12-24 02:01:13.833', '2022-12-24 02:01:13.833', NULL, '', 'jayden', 'jayden', '熊军', '$2a$10$JePrEVFbXjYv/BsDsowslOzjLqzhwEy1CJ1YomxBIyWPmSnImhxP2', NULL, 1, '18584565116', 'jayden@qq.com', 1, 0, NULL, '172.16.0.1', '');

SET FOREIGN_KEY_CHECKS = 1;
