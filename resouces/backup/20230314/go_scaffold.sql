-- MySQL dump 10.13  Distrib 5.7.38, for Linux (x86_64)
--
-- Host: localhost    Database: go_scaffold
-- ------------------------------------------------------
-- Server version	5.7.38

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `go_scaffold`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `go_scaffold` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `go_scaffold`;

--
-- Table structure for table `sys_api_operation_logs`
--

DROP TABLE IF EXISTS `sys_api_operation_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_api_operation_logs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `domain_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '领域ID',
  `ip` varchar(100) NOT NULL COMMENT '请求ip',
  `method` varchar(255) NOT NULL COMMENT '请求方法',
  `path` varchar(255) NOT NULL COMMENT '请求路径',
  `status` int(10) NOT NULL COMMENT '请求状态',
  `latency` int(10) NOT NULL DEFAULT '0' COMMENT '延迟',
  `agent` varchar(255) NOT NULL COMMENT '代理',
  `error` varchar(255) NOT NULL DEFAULT '' COMMENT '错误信息',
  `body` text COMMENT '请求Body',
  `resp` text COMMENT '响应Body',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  PRIMARY KEY (`id`),
  KEY `idx_domain_id_data` (`domain_id`),
  KEY `idx_sys_api_operation_logs_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_api_operation_logs`
--

LOCK TABLES `sys_api_operation_logs` WRITE;
/*!40000 ALTER TABLE `sys_api_operation_logs` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_api_operation_logs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_casbin_rules`
--

DROP TABLE IF EXISTS `sys_casbin_rules`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_casbin_rules` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_sys_casbin_rules` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_casbin_rules`
--

LOCK TABLES `sys_casbin_rules` WRITE;
/*!40000 ALTER TABLE `sys_casbin_rules` DISABLE KEYS */;
INSERT INTO `sys_casbin_rules` VALUES (1,'g','1','1','1','','',''),(13,'g','2','1','1','','',''),(12,'p','1','1','/api.admin.v1.Admin/ListApi','*','','');
/*!40000 ALTER TABLE `sys_casbin_rules` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_depts`
--

DROP TABLE IF EXISTS `sys_depts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_depts` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `domain_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '领域ID',
  `name` varchar(255) NOT NULL COMMENT '资源名称',
  `ancestors` varchar(100) NOT NULL DEFAULT '0' COMMENT '祖级列表',
  `parent_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '父角色ID',
  `sort` int(10) NOT NULL DEFAULT '100' COMMENT '排序',
  `remarks` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `state` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 0 未指定  1 启用 2 停用',
  `leader_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '负责人id',
  PRIMARY KEY (`id`),
  KEY `idx_sys_depts_deleted_at` (`deleted_at`),
  KEY `idx_domain_id_data` (`domain_id`),
  KEY `idx_sys_depts_state` (`state`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_depts`
--

LOCK TABLES `sys_depts` WRITE;
/*!40000 ALTER TABLE `sys_depts` DISABLE KEYS */;
INSERT INTO `sys_depts` VALUES (1,'2022-12-08 06:40:02.892','2022-12-08 06:40:02.892',NULL,1,'产品部','0',0,100,'',1,0),(2,'2023-02-04 08:55:43.809','2023-02-04 08:55:43.809',NULL,1,'研发部','0',0,100,'',1,0),(3,'2023-02-04 08:55:53.432','2023-02-04 08:55:53.432',NULL,1,'设计部','0,2',2,100,'',1,0);
/*!40000 ALTER TABLE `sys_depts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_domain_menus`
--

DROP TABLE IF EXISTS `sys_domain_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_domain_menus` (
  `sys_menu_id` bigint(20) unsigned NOT NULL COMMENT '主键ID',
  `sys_domain_id` bigint(20) unsigned NOT NULL COMMENT '主键ID',
  PRIMARY KEY (`sys_menu_id`,`sys_domain_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_domain_menus`
--

LOCK TABLES `sys_domain_menus` WRITE;
/*!40000 ALTER TABLE `sys_domain_menus` DISABLE KEYS */;
INSERT INTO `sys_domain_menus` VALUES (1,1),(1,2),(2,1),(2,2),(3,1),(3,2),(5,1),(5,2),(6,1),(6,2),(7,1),(7,2),(8,1),(8,2),(9,1),(9,3),(10,1),(10,2),(10,3),(11,1),(11,3),(12,1),(12,2),(12,3),(13,1),(14,1),(14,2),(14,3),(15,1),(15,3),(16,1),(16,2),(16,3),(17,1),(18,1),(19,1),(20,1);
/*!40000 ALTER TABLE `sys_domain_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_domains`
--

DROP TABLE IF EXISTS `sys_domains`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_domains` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `name` varchar(255) NOT NULL COMMENT '领域名称',
  `parent_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '父角色ID',
  `code` varchar(100) NOT NULL COMMENT '领域编码',
  `sort` int(10) NOT NULL DEFAULT '100' COMMENT '排序',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '领域标题',
  `logo` varchar(255) NOT NULL DEFAULT '' COMMENT '领域LOGO',
  `pic` varchar(255) NOT NULL DEFAULT '' COMMENT '领域主图',
  `keywords` varchar(255) NOT NULL DEFAULT '' COMMENT '领域关键字',
  `description` text COMMENT '描述',
  `state` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 0 未指定  1 启用 2 停用',
  `remarks` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_sys_domains_deleted_at` (`deleted_at`),
  KEY `idx_sys_domains_code` (`code`),
  KEY `idx_sys_domains_state` (`state`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_domains`
--

LOCK TABLES `sys_domains` WRITE;
/*!40000 ALTER TABLE `sys_domains` DISABLE KEYS */;
INSERT INTO `sys_domains` VALUES (1,'2023-01-04 17:39:54.000','2023-02-09 09:08:46.269',NULL,'中台管理平台',0,'1557658909480354584',100,'','','','',NULL,1,''),(2,'2023-02-17 09:36:35.479','2023-02-17 09:46:26.868',NULL,'中台运营平台',1,'1557658909480354583',100,'','','','',NULL,1,''),(3,'2023-01-30 07:25:28.254','2023-02-17 09:37:06.563',NULL,'众缘教育',0,'1619959967261921280',100,'','','','',NULL,1,''),(4,'2023-01-30 07:25:34.733','2023-02-17 09:46:43.835',NULL,'美墨瑞教育',3,'1619959994436816896',100,'','','','',NULL,1,'');
/*!40000 ALTER TABLE `sys_domains` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_menu_buttons`
--

DROP TABLE IF EXISTS `sys_menu_buttons`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_menu_buttons` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `menu_id` bigint(20) unsigned NOT NULL COMMENT '菜单ID',
  `name` varchar(255) NOT NULL COMMENT '按钮关键key',
  `remarks` varchar(255) NOT NULL DEFAULT '' COMMENT '按钮备注',
  PRIMARY KEY (`id`),
  KEY `idx_sys_menu_buttons_menu_id` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_menu_buttons`
--

LOCK TABLES `sys_menu_buttons` WRITE;
/*!40000 ALTER TABLE `sys_menu_buttons` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_menu_buttons` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_menu_parameters`
--

DROP TABLE IF EXISTS `sys_menu_parameters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_menu_parameters` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `menu_id` bigint(20) unsigned NOT NULL COMMENT '菜单ID',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '地址栏携带参类型 0 未指定 1 params 2 query',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '地址栏携带参数的名称',
  `value` varchar(255) NOT NULL DEFAULT '' COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`),
  KEY `idx_sys_menu_parameters_menu_id` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_menu_parameters`
--

LOCK TABLES `sys_menu_parameters` WRITE;
/*!40000 ALTER TABLE `sys_menu_parameters` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_menu_parameters` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_menu_resources`
--

DROP TABLE IF EXISTS `sys_menu_resources`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_menu_resources` (
  `sys_menu_id` bigint(20) unsigned NOT NULL COMMENT '主键ID',
  `sys_resource_id` bigint(20) unsigned NOT NULL COMMENT '主键ID',
  PRIMARY KEY (`sys_menu_id`,`sys_resource_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_menu_resources`
--

LOCK TABLES `sys_menu_resources` WRITE;
/*!40000 ALTER TABLE `sys_menu_resources` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_menu_resources` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_menus`
--

DROP TABLE IF EXISTS `sys_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_menus` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '菜单/路由名称',
  `parent_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '父菜单ID',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '菜单类型 0 无指定 1 目录 2 菜单 3 功能(按钮等)',
  `path` varchar(255) NOT NULL DEFAULT '' COMMENT '路由地址',
  `component` varchar(255) NOT NULL DEFAULT '' COMMENT '组件路径',
  `redirect` varchar(255) NOT NULL DEFAULT '' COMMENT '重定向地址',
  `permission` varchar(255) NOT NULL DEFAULT '' COMMENT '权限标识',
  `sort` int(10) NOT NULL DEFAULT '10' COMMENT '排序标记',
  `icon` varchar(255) NOT NULL DEFAULT '' COMMENT '图标',
  `title` varchar(255) NOT NULL COMMENT '菜单标题',
  `is_hidden` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否隐藏 0 无指定 1 是 2 否',
  `is_cache` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否缓存 0 无指定 1 是 2 否',
  `is_affix` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否固定 0 无指定 1 是 2 否',
  `link_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '外链类型  0 无指定 1 无 2 内嵌 3 跳转',
  `link_url` varchar(255) NOT NULL DEFAULT '' COMMENT '链接地址',
  `remarks` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_menu_type_name` (`name`,`type`),
  KEY `idx_sys_menus_deleted_at` (`deleted_at`),
  KEY `idx_sys_menus_parent_id` (`parent_id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_menus`
--

LOCK TABLES `sys_menus` WRITE;
/*!40000 ALTER TABLE `sys_menus` DISABLE KEYS */;
INSERT INTO `sys_menus` VALUES (1,'2022-12-08 03:03:34.028','2022-12-08 07:36:41.721',NULL,'Dashboard',0,1,'/dashboard','LAYOUT','','',100,'ion:grid-outline','仪表盘',2,2,2,1,'',''),(2,'2022-12-08 03:06:12.140','2023-01-11 03:27:35.110',NULL,'Analysis',1,2,'analysis','/dashboard/analysis/index','','dashboardAnalyse',100,'ant-design:alert-twotone','分析页',2,2,2,1,'',''),(3,'2022-12-08 03:50:55.084','2022-12-08 06:35:34.566',NULL,'Add',10,3,'','','','userAdd',100,'ant-design:user-add-outlined','新增',2,2,2,1,'',''),(5,'2022-12-08 05:56:56.234','2022-12-08 06:35:43.458',NULL,'Edit',10,3,'','','','userEdit',100,'ant-design:edit-filled','编辑',2,2,2,1,'',''),(6,'2022-12-08 06:18:26.477','2022-12-08 06:35:51.860',NULL,'Delete',10,3,'','','','userDelete',100,'ant-design:delete-twotone','删除',2,2,2,1,'',''),(7,'2022-12-08 06:19:15.292','2022-12-08 06:36:05.088',NULL,'Update',10,3,'','','','userUpdate',100,'ant-design:upload-outlined','修改',2,2,2,1,'',''),(8,'2022-12-08 07:40:29.095','2023-02-09 08:52:30.195',NULL,'Workbench',1,2,'workbench','/dashboard/workbench/index','','workbench',100,'','工作台',2,2,2,1,'http://localhost',''),(9,'2022-12-08 07:41:41.753','2022-12-08 07:41:41.753',NULL,'System',0,1,'/system','LAYOUT','','',100,'ion:settings-outline','系统管理',2,2,2,1,'',''),(10,'2022-12-08 07:45:05.125','2022-12-08 07:45:05.125',NULL,'UserManagement',9,2,'user','/core/user/index','','user',100,'','用户管理',2,2,2,1,'',''),(11,'2022-12-08 07:47:50.582','2023-02-17 03:45:59.118',NULL,'UserDetail',10,2,'user_detail/:id','/core/user/UserDetail','','userDetail',100,'','用户详情',1,2,2,1,'',''),(12,'2022-12-08 07:48:47.697','2022-12-08 07:48:47.697',NULL,'RoleManagement',9,2,'role','/core/role/index','','role',100,'','角色管理',2,2,2,1,'',''),(13,'2022-12-08 07:49:18.253','2023-01-11 03:26:09.922',NULL,'MenuManagement',17,2,'menu','/core/menu/index','','menu',100,'','菜单管理',2,2,2,1,'',''),(14,'2022-12-08 07:49:43.133','2022-12-08 07:49:43.133',NULL,'DeptManagement',9,2,'dept','/core/dept/index','','dept',100,'','部门管理',2,2,2,1,'',''),(15,'2022-12-08 07:50:54.820','2023-02-17 03:46:05.894',NULL,'ChangePassword',10,2,'change_password/:id','/core/password/index','','changePassword',100,'','修改密码',1,2,2,1,'',''),(16,'2023-01-10 02:12:34.805','2023-01-10 02:12:34.805',NULL,'PostManagement',9,2,'post','/core/post/index','','post',100,'','岗位管理',2,2,2,1,'',''),(17,'2023-01-11 03:22:07.673','2023-01-11 03:22:07.673',NULL,'Middleground',0,1,'/middleground','LAYOUT','','',100,'ant-design:setting-outlined','中台管理',2,2,2,1,'',''),(18,'2023-01-11 03:24:59.502','2023-01-11 03:24:59.502',NULL,'DomainManagement',17,2,'domain','/core/domain/index','','domain',100,'','租户管理',2,2,2,1,'',''),(19,'2023-01-11 03:25:50.886','2023-01-11 03:25:50.886',NULL,'ResourceManagement',17,2,'resource','/core/resource/index','','resource',100,'','资源管理',2,2,2,1,'',''),(20,'2022-12-08 03:06:12.140','2023-02-09 09:48:40.475',NULL,'Step',0,2,'step','/demo/page/form/step/index','','step',100,'ant-design:alert-twotone','步骤演示',2,2,2,1,'http://localhost',''),(21,'2023-02-14 09:36:07.747','2023-03-09 13:05:01.523',NULL,'baidu',0,2,'baidu','','','',100,'','百度',2,2,2,2,'http://www.baidu.com',''),(22,'2023-02-14 10:59:18.875','2023-02-14 10:59:18.875',NULL,'create',12,3,'','','','roleCreate',100,'','新增',2,1,1,1,'','');
/*!40000 ALTER TABLE `sys_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_posts`
--

DROP TABLE IF EXISTS `sys_posts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_posts` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `domain_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '领域ID',
  `name` varchar(255) NOT NULL COMMENT '岗位名称',
  `sort` int(10) NOT NULL DEFAULT '100' COMMENT '排序',
  `state` tinyint(1) NOT NULL DEFAULT '1' COMMENT '岗位状态 0 未指定  1 启用 2 停用',
  `remarks` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `code` varchar(100) NOT NULL DEFAULT '' COMMENT '岗位编码',
  PRIMARY KEY (`id`),
  KEY `idx_sys_posts_state` (`state`),
  KEY `idx_sys_posts_deleted_at` (`deleted_at`),
  KEY `idx_domain_id_data` (`domain_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_posts`
--

LOCK TABLES `sys_posts` WRITE;
/*!40000 ALTER TABLE `sys_posts` DISABLE KEYS */;
INSERT INTO `sys_posts` VALUES (1,'2023-01-10 02:21:41.604','2023-01-10 07:07:40.646',NULL,1,'董事长',100,1,'',''),(2,'2023-01-10 07:26:19.430','2023-01-10 07:26:20.292',NULL,1,'总经理',99,1,'总经理岗位','1612708008079396864');
/*!40000 ALTER TABLE `sys_posts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_resources`
--

DROP TABLE IF EXISTS `sys_resources`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_resources` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `name` varchar(255) NOT NULL COMMENT '资源名称',
  `path` varchar(255) NOT NULL COMMENT '请求路径',
  `method` varchar(255) NOT NULL DEFAULT 'POST' COMMENT '方法',
  `operation` varchar(255) NOT NULL DEFAULT '' COMMENT '请求动作',
  `group` varchar(255) NOT NULL COMMENT 'api分组',
  `description` varchar(255) NOT NULL DEFAULT '' COMMENT 'api描述',
  `remarks` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_api_path_method` (`path`,`method`,`operation`),
  KEY `idx_sys_resources_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_resources`
--

LOCK TABLES `sys_resources` WRITE;
/*!40000 ALTER TABLE `sys_resources` DISABLE KEYS */;
INSERT INTO `sys_resources` VALUES (1,'2022-11-05 07:38:29.503','2022-12-24 07:03:25.647',NULL,'资源列表','/api.admin.v1.Admin/ListResource','*','/v1/admin/resources','默认','',''),(2,'2022-12-24 07:02:54.516','2022-12-24 07:30:14.589',NULL,'菜单列表','/api.admin.v1.Admin/ListMenu','*','/v1/admin/menus','默认','',''),(3,'2022-12-24 07:32:38.919','2022-12-24 07:33:01.874',NULL,'领域列表','/api.admin.v1.Admin/ListDomain','*','/v1/admin/domains','领域','','');
/*!40000 ALTER TABLE `sys_resources` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_depts`
--

DROP TABLE IF EXISTS `sys_role_depts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role_depts` (
  `sys_dept_id` bigint(20) unsigned NOT NULL COMMENT '主键ID',
  `sys_role_id` bigint(20) unsigned NOT NULL COMMENT '主键ID',
  PRIMARY KEY (`sys_dept_id`,`sys_role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_depts`
--

LOCK TABLES `sys_role_depts` WRITE;
/*!40000 ALTER TABLE `sys_role_depts` DISABLE KEYS */;
INSERT INTO `sys_role_depts` VALUES (1,1);
/*!40000 ALTER TABLE `sys_role_depts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_menus`
--

DROP TABLE IF EXISTS `sys_role_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role_menus` (
  `created_at` datetime(3) DEFAULT NULL,
  `sys_role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `sys_menu_id` bigint(20) unsigned NOT NULL COMMENT '菜单ID',
  `sys_menu_button` json DEFAULT NULL COMMENT '菜单按钮',
  `sys_menu_parameter` json DEFAULT NULL COMMENT '菜单参数',
  UNIQUE KEY `idx_role_menu_role_id_menu_id` (`sys_role_id`,`sys_menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_menus`
--

LOCK TABLES `sys_role_menus` WRITE;
/*!40000 ALTER TABLE `sys_role_menus` DISABLE KEYS */;
INSERT INTO `sys_role_menus` VALUES ('2023-02-09 09:37:46.385',1,1,'[]','[]'),('2023-02-09 09:37:46.385',1,2,'[]','[]'),('2023-02-09 09:37:46.385',1,3,'[]','[]'),('2023-02-09 09:37:46.385',1,5,'[]','[]'),('2023-02-09 09:37:46.385',1,6,'[]','[]'),('2023-02-09 09:37:46.385',1,7,'[]','[]'),('2023-02-09 09:37:46.385',1,8,'[]','[]'),('2023-02-09 09:37:46.385',1,9,'[]','[]'),('2023-02-09 09:37:46.385',1,10,'[]','[]'),('2023-02-09 09:37:46.385',1,11,'[]','[]'),('2023-02-09 09:37:46.385',1,12,'[]','[]'),('2023-02-09 09:37:46.385',1,13,'[]','[]'),('2023-02-09 09:37:46.385',1,14,'[]','[]'),('2023-02-09 09:37:46.385',1,15,'[]','[]'),('2023-02-09 09:37:46.385',1,16,'[]','[]'),('2023-02-09 09:37:46.385',1,17,'[]','[]'),('2023-02-09 09:37:46.385',1,18,'[]','[]'),('2023-02-09 09:37:46.385',1,19,'[]','[]'),('2023-02-09 09:37:46.385',1,20,'[]','[]'),('2022-12-23 02:05:39.053',2,3,'[]','[]'),('2022-12-23 02:05:39.053',2,6,'[]','[]'),('2022-12-23 02:05:39.053',2,7,'[]','[]'),('2022-12-23 02:05:39.053',2,8,'[]','[]'),('2022-12-23 02:05:39.053',2,9,'[]','[]'),('2022-12-23 02:05:39.053',2,10,'[]','[]'),('2022-12-23 02:05:39.053',2,12,'[]','[]'),('2022-12-23 02:05:39.053',2,13,'[]','[]'),('2022-12-23 02:05:39.053',2,14,'[]','[]'),('2022-12-23 02:05:39.053',2,16,'[]','[]');
/*!40000 ALTER TABLE `sys_role_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_relations`
--

DROP TABLE IF EXISTS `sys_role_relations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role_relations` (
  `sys_role_id` bigint(20) unsigned NOT NULL COMMENT '主键ID',
  `role_id` bigint(20) unsigned NOT NULL COMMENT '主键ID',
  PRIMARY KEY (`sys_role_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_relations`
--

LOCK TABLES `sys_role_relations` WRITE;
/*!40000 ALTER TABLE `sys_role_relations` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_role_relations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_resources`
--

DROP TABLE IF EXISTS `sys_role_resources`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role_resources` (
  `sys_role_id` bigint(20) unsigned NOT NULL COMMENT '主键ID',
  `sys_resource_id` bigint(20) unsigned NOT NULL COMMENT '主键ID',
  PRIMARY KEY (`sys_role_id`,`sys_resource_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_resources`
--

LOCK TABLES `sys_role_resources` WRITE;
/*!40000 ALTER TABLE `sys_role_resources` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_role_resources` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_roles`
--

DROP TABLE IF EXISTS `sys_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_roles` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `domain_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '领域ID',
  `name` varchar(255) NOT NULL COMMENT '角色名称',
  `parent_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '父角色ID',
  `default_router` varchar(255) NOT NULL DEFAULT '/dashboard' COMMENT '默认路由',
  `sort` int(10) NOT NULL DEFAULT '100' COMMENT '排序',
  `state` tinyint(1) NOT NULL DEFAULT '1' COMMENT '角色状态 0 未指定  1 启用 2 停用',
  `remarks` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `data_scope` tinyint(2) NOT NULL DEFAULT '1' COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `menu_check_strictly` tinyint(2) NOT NULL DEFAULT '1' COMMENT '菜单树选择项是否关联显示',
  `dept_check_strictly` tinyint(2) NOT NULL DEFAULT '1' COMMENT '部门树选择项是否关联显示',
  PRIMARY KEY (`id`),
  KEY `idx_sys_roles_deleted_at` (`deleted_at`),
  KEY `idx_domain_id_data` (`domain_id`,`name`),
  KEY `idx_sys_roles_state` (`state`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_roles`
--

LOCK TABLES `sys_roles` WRITE;
/*!40000 ALTER TABLE `sys_roles` DISABLE KEYS */;
INSERT INTO `sys_roles` VALUES (1,'2022-11-05 06:21:06.260','2023-02-09 09:37:46.359',NULL,1,'超级管理员',0,'',0,1,'',0,0,0),(2,'2022-11-05 09:07:56.985','2022-12-24 01:44:20.580',NULL,1,'default',0,'',0,1,'',1,1,1),(3,'2023-01-30 07:25:28.261','2023-02-04 08:59:36.011','2023-02-17 09:16:16.290',1,'default',0,'/dashboard',0,2,'',1,1,1),(4,'2023-01-30 07:25:34.735','2023-02-04 08:59:35.424','2023-02-17 09:16:01.274',1,'default',0,'/dashboard',0,2,'',1,1,1);
/*!40000 ALTER TABLE `sys_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_posts`
--

DROP TABLE IF EXISTS `sys_user_posts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_user_posts` (
  `sys_post_id` bigint(20) unsigned NOT NULL COMMENT '主键ID',
  `sys_user_id` bigint(20) unsigned NOT NULL COMMENT '主键ID',
  PRIMARY KEY (`sys_post_id`,`sys_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_posts`
--

LOCK TABLES `sys_user_posts` WRITE;
/*!40000 ALTER TABLE `sys_user_posts` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_user_posts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_users`
--

DROP TABLE IF EXISTS `sys_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `domain_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '领域ID',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `name` varchar(255) NOT NULL COMMENT '用户名称',
  `nick_name` varchar(255) NOT NULL DEFAULT '' COMMENT '昵称',
  `real_name` varchar(100) NOT NULL DEFAULT '' COMMENT '实名',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `birthday` datetime DEFAULT NULL COMMENT '生日',
  `gender` tinyint(1) NOT NULL DEFAULT '1' COMMENT '性别 0 未指定 1 男 2 女',
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `state` tinyint(1) NOT NULL DEFAULT '1' COMMENT '用户状态 0 未指定  1 启用 2 停用',
  `last_use_role` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后使用角色',
  `last_login_at` datetime DEFAULT NULL COMMENT '最后登录时间',
  `last_login_ip` varchar(50) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `remarks` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_users_mobile_email` (`phone`,`email`),
  KEY `idx_sys_users_state` (`state`),
  KEY `idx_sys_users_deleted_at` (`deleted_at`),
  KEY `idx_users_name_nick_name_real_name` (`name`,`real_name`,`nick_name`),
  KEY `idx_users_phone_email` (`phone`,`email`),
  KEY `idx_domain_id_data` (`domain_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_users`
--

LOCK TABLES `sys_users` WRITE;
/*!40000 ALTER TABLE `sys_users` DISABLE KEYS */;
INSERT INTO `sys_users` VALUES (1,'2022-11-05 06:17:21.364','2023-03-09 12:48:49.807',NULL,1,'','xiong','xiong','熊龙军','$2a$10$rxZ7O60l2fN1DJuXBNKWgub50t320mN4.7innl5b6Gv34XQkU/IZa',NULL,1,'18584565115','xiong@qq.com',1,1,'2023-03-09 12:48:50','0.0.0.0',''),(2,'2022-12-24 02:01:13.833','2022-12-24 02:01:13.833',NULL,1,'','jayden','jayden','熊军','$2a$10$JePrEVFbXjYv/BsDsowslOzjLqzhwEy1CJ1YomxBIyWPmSnImhxP2',NULL,1,'18584565116','jayden@qq.com',1,2,NULL,'172.16.0.1','');
/*!40000 ALTER TABLE `sys_users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-03-14 18:49:15
