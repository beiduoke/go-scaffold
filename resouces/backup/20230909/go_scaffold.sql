-- MySQL dump 10.13  Distrib 5.7.42, for Linux (x86_64)
--
-- Host: localhost    Database: go_scaffold
-- ------------------------------------------------------
-- Server version	5.7.42

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
-- Table structure for table `bus_demos`
--

DROP TABLE IF EXISTS `bus_demos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `bus_demos` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `creator` varchar(64) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '修改时间',
  `updater` varchar(64) NOT NULL DEFAULT '' COMMENT '更新者',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `domain_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '租户ID',
  `name` varchar(255) NOT NULL COMMENT '名称',
  `sort` int(10) NOT NULL DEFAULT '100' COMMENT '排序',
  `state` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 0 未指定  1 启用 2 停用',
  `remarks` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_bus_demos_deleted_at` (`deleted_at`),
  KEY `idx_domain_id_data` (`domain_id`),
  KEY `idx_bus_demos_state` (`state`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bus_demos`
--

LOCK TABLES `bus_demos` WRITE;
/*!40000 ALTER TABLE `bus_demos` DISABLE KEYS */;
/*!40000 ALTER TABLE `bus_demos` ENABLE KEYS */;
UNLOCK TABLES;

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
  `domain_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '租户ID',
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
  `creator` varchar(64) NOT NULL DEFAULT '' COMMENT '创建者',
  `updater` varchar(64) NOT NULL DEFAULT '' COMMENT '更新者',
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
) ENGINE=InnoDB AUTO_INCREMENT=1263 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_casbin_rules`
--

LOCK TABLES `sys_casbin_rules` WRITE;
/*!40000 ALTER TABLE `sys_casbin_rules` DISABLE KEYS */;
INSERT INTO `sys_casbin_rules` VALUES (1112,'g','1','1','1','','',''),(1110,'g','12','3','1','','',''),(1104,'g','13','5','1','','',''),(1108,'g','2','2','1','','',''),(1113,'g','3','3','1','','',''),(1111,'g','4','4','1','','',''),(1238,'p','2','/api.server.v1.Api/CreateDept','*','1','',''),(1245,'p','2','/api.server.v1.Api/CreateDict','*','1','',''),(1249,'p','2','/api.server.v1.Api/CreateDictData','*','1','',''),(1230,'p','2','/api.server.v1.Api/CreateDomain','*','1','',''),(1241,'p','2','/api.server.v1.Api/CreateMenu','*','1','',''),(1233,'p','2','/api.server.v1.Api/CreatePost','*','1','',''),(1227,'p','2','/api.server.v1.Api/CreateRole','*','1','',''),(1215,'p','2','/api.server.v1.Api/CreateUser','*','1','',''),(1214,'p','2','/api.server.v1.Api/DashboardAnalyse','*','1','',''),(1218,'p','2','/api.server.v1.Api/DashboardWorkbench','*','1','',''),(1240,'p','2','/api.server.v1.Api/DeleteDept','*','1','',''),(1247,'p','2','/api.server.v1.Api/DeleteDict','*','1','',''),(1251,'p','2','/api.server.v1.Api/DeleteDictData','*','1','',''),(1232,'p','2','/api.server.v1.Api/DeleteDomain','*','1','',''),(1243,'p','2','/api.server.v1.Api/DeleteMenu','*','1','',''),(1235,'p','2','/api.server.v1.Api/DeletePost','*','1','',''),(1237,'p','2','/api.server.v1.Api/DeleteRole','*','1','',''),(1216,'p','2','/api.server.v1.Api/DeleteUser','*','1','',''),(1261,'p','2','/api.server.v1.Api/ExistUserName','*','1','',''),(1257,'p','2','/api.server.v1.Api/GetDept','*','1','',''),(1260,'p','2','/api.server.v1.Api/GetDomain','*','1','',''),(1259,'p','2','/api.server.v1.Api/GetMenu','*','1','',''),(1258,'p','2','/api.server.v1.Api/GetPost','*','1','',''),(1256,'p','2','/api.server.v1.Api/GetRole','*','1','',''),(1220,'p','2','/api.server.v1.Api/GetUser','*','1','',''),(1252,'p','2','/api.server.v1.Api/GetUserInfo','*','1','',''),(1229,'p','2','/api.server.v1.Api/HandleRoleDataScope','*','1','',''),(1228,'p','2','/api.server.v1.Api/HandleRoleMenu','*','1','',''),(1223,'p','2','/api.server.v1.Api/ListDeptTree','*','1','',''),(1244,'p','2','/api.server.v1.Api/ListDict','*','1','',''),(1248,'p','2','/api.server.v1.Api/ListDictData','*','1','',''),(1226,'p','2','/api.server.v1.Api/ListDomainTree','*','1','',''),(1222,'p','2','/api.server.v1.Api/ListMenuTree','*','1','',''),(1225,'p','2','/api.server.v1.Api/ListPost','*','1','',''),(1221,'p','2','/api.server.v1.Api/ListRole','*','1','',''),(1219,'p','2','/api.server.v1.Api/ListUser','*','1','',''),(1253,'p','2','/api.server.v1.Api/ListUserRole','*','1','',''),(1254,'p','2','/api.server.v1.Api/ListUserRoleMenuRouterTree','*','1','',''),(1255,'p','2','/api.server.v1.Api/ListUserRolePermission','*','1','',''),(1262,'p','2','/api.server.v1.Api/Logout','*','1','',''),(1239,'p','2','/api.server.v1.Api/UpdateDept','*','1','',''),(1246,'p','2','/api.server.v1.Api/UpdateDict','*','1','',''),(1250,'p','2','/api.server.v1.Api/UpdateDictData','*','1','',''),(1231,'p','2','/api.server.v1.Api/UpdateDomain','*','1','',''),(1242,'p','2','/api.server.v1.Api/UpdateMenu','*','1','',''),(1234,'p','2','/api.server.v1.Api/UpdatePost','*','1','',''),(1236,'p','2','/api.server.v1.Api/UpdateRole','*','1','',''),(1217,'p','2','/api.server.v1.Api/UpdateUser','*','1','',''),(1224,'p','2','/api.server.v1.Api/UpdateUserPassword','*','1','',''),(1202,'p','3','/api.server.v1.Api/CreateDept','*','1','',''),(1197,'p','3','/api.server.v1.Api/CreatePost','*','1','',''),(1194,'p','3','/api.server.v1.Api/CreateRole','*','1','',''),(1185,'p','3','/api.server.v1.Api/CreateUser','*','1','',''),(1204,'p','3','/api.server.v1.Api/DeleteDept','*','1','',''),(1199,'p','3','/api.server.v1.Api/DeletePost','*','1','',''),(1201,'p','3','/api.server.v1.Api/DeleteRole','*','1','',''),(1186,'p','3','/api.server.v1.Api/DeleteUser','*','1','',''),(1212,'p','3','/api.server.v1.Api/ExistUserName','*','1','',''),(1210,'p','3','/api.server.v1.Api/GetDept','*','1','',''),(1211,'p','3','/api.server.v1.Api/GetPost','*','1','',''),(1209,'p','3','/api.server.v1.Api/GetRole','*','1','',''),(1189,'p','3','/api.server.v1.Api/GetUser','*','1','',''),(1205,'p','3','/api.server.v1.Api/GetUserInfo','*','1','',''),(1196,'p','3','/api.server.v1.Api/HandleRoleDataScope','*','1','',''),(1195,'p','3','/api.server.v1.Api/HandleRoleMenu','*','1','',''),(1191,'p','3','/api.server.v1.Api/ListDeptTree','*','1','',''),(1193,'p','3','/api.server.v1.Api/ListPost','*','1','',''),(1190,'p','3','/api.server.v1.Api/ListRole','*','1','',''),(1188,'p','3','/api.server.v1.Api/ListUser','*','1','',''),(1206,'p','3','/api.server.v1.Api/ListUserRole','*','1','',''),(1207,'p','3','/api.server.v1.Api/ListUserRoleMenuRouterTree','*','1','',''),(1208,'p','3','/api.server.v1.Api/ListUserRolePermission','*','1','',''),(1213,'p','3','/api.server.v1.Api/Logout','*','1','',''),(1203,'p','3','/api.server.v1.Api/UpdateDept','*','1','',''),(1198,'p','3','/api.server.v1.Api/UpdatePost','*','1','',''),(1200,'p','3','/api.server.v1.Api/UpdateRole','*','1','',''),(1187,'p','3','/api.server.v1.Api/UpdateUser','*','1','',''),(1192,'p','3','/api.server.v1.Api/UpdateUserPassword','*','1','',''),(1076,'p','4','/api.server.v1.Api/GetUserInfo','*','1','',''),(1077,'p','4','/api.server.v1.Api/ListUserRole','*','1','',''),(1078,'p','4','/api.server.v1.Api/ListUserRoleMenuRouterTree','*','1','',''),(1079,'p','4','/api.server.v1.Api/ListUserRolePermission','*','1','',''),(1151,'p','5','/api.server.v1.Api/DashboardAnalyse','*','1','',''),(1152,'p','5','/api.server.v1.Api/GetUserInfo','*','1','',''),(1153,'p','5','/api.server.v1.Api/ListUserRole','*','1','',''),(1154,'p','5','/api.server.v1.Api/ListUserRoleMenuRouterTree','*','1','',''),(1155,'p','5','/api.server.v1.Api/ListUserRolePermission','*','1','',''),(1156,'p','5','/api.server.v1.Api/Logout','*','1','','');
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
  `domain_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '租户ID',
  `name` varchar(255) NOT NULL COMMENT '资源名称',
  `ancestors` varchar(100) NOT NULL DEFAULT '0' COMMENT '祖级列表',
  `parent_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '父角色ID',
  `sort` int(10) NOT NULL DEFAULT '100' COMMENT '排序',
  `remarks` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `state` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 0 未指定  1 启用 2 停用',
  `leader_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '负责人id',
  `creator` varchar(64) NOT NULL DEFAULT '' COMMENT '创建者',
  `updater` varchar(64) NOT NULL DEFAULT '' COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_sys_depts_deleted_at` (`deleted_at`),
  KEY `idx_domain_id_data` (`domain_id`),
  KEY `idx_sys_depts_state` (`state`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_depts`
--

LOCK TABLES `sys_depts` WRITE;
/*!40000 ALTER TABLE `sys_depts` DISABLE KEYS */;
INSERT INTO `sys_depts` VALUES (1,'2022-12-08 06:40:02.892','2023-03-24 09:38:40.233',NULL,1,'产品部','0,2',2,100,'',1,0,'',''),(2,'2023-02-04 08:55:43.809','2023-02-04 08:55:43.809',NULL,1,'研发部','0',0,100,'',1,0,'',''),(3,'2023-02-04 08:55:53.432','2023-02-04 08:55:53.432',NULL,1,'设计部','0,2',2,100,'',1,0,'',''),(4,'2023-07-11 08:30:12.584','2023-07-11 08:30:13.722',NULL,1,'前端部','0,2',2,99,'',1,0,'',''),(5,'2023-03-24 06:03:38.833','2023-03-24 06:03:38.833',NULL,1,'后端部','0,2',2,100,'',1,0,'',''),(6,'2023-03-24 09:35:15.497','2023-03-24 09:35:15.497',NULL,1,'UI设计','0,2,3',3,100,'',1,0,'',''),(7,'2023-07-11 08:36:25.979','2023-07-11 08:36:27.653',NULL,1,'财务部','0',0,100,'',1,0,'','');
/*!40000 ALTER TABLE `sys_depts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dict_data`
--

DROP TABLE IF EXISTS `sys_dict_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_dict_data` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `creator` varchar(64) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '修改时间',
  `updater` varchar(64) NOT NULL DEFAULT '' COMMENT '更新者',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `dict_type` varchar(100) NOT NULL COMMENT '字典类型',
  `label` varchar(255) NOT NULL COMMENT '字典标签',
  `value` varchar(255) NOT NULL DEFAULT '' COMMENT '字典键值',
  `color_type` varchar(100) NOT NULL DEFAULT '' COMMENT '颜色类型',
  `css_class` varchar(100) NOT NULL DEFAULT '' COMMENT 'CSS样式',
  `sort` int(10) NOT NULL DEFAULT '100' COMMENT '排序',
  `state` tinyint(1) NOT NULL DEFAULT '1' COMMENT '字典状态 0 未指定  1 启用 2 停用',
  `remarks` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dict_data_deleted_at` (`deleted_at`),
  KEY `idx_sys_dict_data_state` (`state`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dict_data`
--

LOCK TABLES `sys_dict_data` WRITE;
/*!40000 ALTER TABLE `sys_dict_data` DISABLE KEYS */;
INSERT INTO `sys_dict_data` VALUES (1,'2023-07-11 08:40:49.211','','2023-07-11 08:40:50.290','',NULL,'system_user_gender','11','11','','',100,1,''),(3,'2023-06-08 09:00:47.035','','2023-06-08 09:00:47.977','',NULL,'system_user_age','11岁-13岁','11-13','','',100,1,'');
/*!40000 ALTER TABLE `sys_dict_data` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dicts`
--

DROP TABLE IF EXISTS `sys_dicts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_dicts` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `creator` varchar(64) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '修改时间',
  `updater` varchar(64) NOT NULL DEFAULT '' COMMENT '更新者',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `name` varchar(255) NOT NULL COMMENT '字典名称',
  `type` varchar(100) NOT NULL COMMENT '字典类型',
  `sort` int(10) NOT NULL DEFAULT '100' COMMENT '排序',
  `state` tinyint(1) NOT NULL DEFAULT '1' COMMENT '字典状态 0 未指定  1 启用 2 停用',
  `remarks` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dicts_deleted_at` (`deleted_at`),
  KEY `idx_sys_dicts_state` (`state`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dicts`
--

LOCK TABLES `sys_dicts` WRITE;
/*!40000 ALTER TABLE `sys_dicts` DISABLE KEYS */;
INSERT INTO `sys_dicts` VALUES (1,'2023-07-11 08:01:14.865','','2023-07-11 08:01:15.703','',NULL,'用户性别','system_user_gender',100,1,''),(2,'2023-06-08 08:59:03.537','','2023-06-08 08:59:03.537','',NULL,'年龄段','system_user_age',100,1,'');
/*!40000 ALTER TABLE `sys_dicts` ENABLE KEYS */;
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
INSERT INTO `sys_domain_menus` VALUES (1,1),(2,1),(3,1),(6,1),(7,1),(8,1),(9,1),(10,1),(11,1),(12,1),(13,1),(14,1),(14,2),(15,1),(16,1),(17,1),(18,1),(20,5),(21,1),(22,1),(23,1),(23,5),(25,1),(27,1),(28,1),(29,1),(30,1),(31,1),(32,1),(33,1),(34,1),(34,3),(36,1),(37,1),(38,1),(39,1),(39,2),(39,4),(40,1),(40,2),(41,1),(41,2),(42,1),(42,2),(43,1),(44,1),(45,1),(46,1),(47,1),(48,1),(49,1),(50,1),(51,1),(52,1),(53,1),(54,1),(55,1),(56,1),(57,1),(58,1),(59,1),(60,1),(61,1),(62,1),(63,1),(64,1),(65,1),(66,1),(67,1),(68,1),(69,1),(70,1);
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
  `name` varchar(255) NOT NULL COMMENT '租户名称',
  `parent_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '父角色ID',
  `super_user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '超级用户ID',
  `code` varchar(100) NOT NULL COMMENT '租户编码',
  `sort` int(10) NOT NULL DEFAULT '100' COMMENT '排序',
  `alias` varchar(255) NOT NULL DEFAULT '' COMMENT '租户别名',
  `logo` varchar(255) NOT NULL DEFAULT '' COMMENT '租户LOGO',
  `pic` varchar(255) NOT NULL DEFAULT '' COMMENT '租户主图',
  `keywords` varchar(255) NOT NULL DEFAULT '' COMMENT '租户关键字',
  `description` text COMMENT '描述',
  `state` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 0 未指定  1 启用 2 停用',
  `remarks` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `creator` varchar(64) NOT NULL DEFAULT '' COMMENT '创建者',
  `updater` varchar(64) NOT NULL DEFAULT '' COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_sys_domains_deleted_at` (`deleted_at`),
  KEY `idx_sys_domains_code` (`code`),
  KEY `idx_sys_domains_state` (`state`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_domains`
--

LOCK TABLES `sys_domains` WRITE;
/*!40000 ALTER TABLE `sys_domains` DISABLE KEYS */;
INSERT INTO `sys_domains` VALUES (1,'2023-07-11 08:01:19.944','2023-07-11 08:01:20.592',NULL,'中台管理平台',0,1,'1557658909480354584',100,'','','','',NULL,1,'','',''),(2,'2023-02-17 09:36:35.479','2023-06-12 13:02:48.621',NULL,'中台运营平台',1,0,'1557658909480354583',100,'','','','',NULL,1,'','',''),(3,'2023-01-30 07:25:28.254','2023-06-12 13:01:58.709',NULL,'众缘教育',0,0,'1619959967261921280',100,'','','','',NULL,1,'','',''),(4,'2023-01-30 07:25:34.733','2023-06-12 13:08:51.232',NULL,'美墨瑞教育',3,0,'1619959994436816896',100,'','','','',NULL,1,'','',''),(5,'2023-03-15 10:18:06.921','2023-06-12 13:09:31.733',NULL,'测试租户',0,0,'1635904182743470080',100,'','','','','',1,'','','');
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
  `api_resource` varchar(255) NOT NULL DEFAULT '' COMMENT '接口资源',
  `remarks` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `creator` varchar(64) NOT NULL DEFAULT '' COMMENT '创建者',
  `updater` varchar(64) NOT NULL DEFAULT '' COMMENT '更新者',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_menu_type_name` (`name`,`type`),
  KEY `idx_sys_menus_deleted_at` (`deleted_at`),
  KEY `idx_sys_menus_parent_id` (`parent_id`)
) ENGINE=InnoDB AUTO_INCREMENT=78 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_menus`
--

LOCK TABLES `sys_menus` WRITE;
/*!40000 ALTER TABLE `sys_menus` DISABLE KEYS */;
INSERT INTO `sys_menus` VALUES (1,'2022-12-08 03:03:34.028','2023-05-25 07:34:00.808',NULL,'Dashboard',0,1,'/dashboard','LAYOUT','','',1,'ion:grid-outline','仪表盘',2,0,0,0,'','','','',''),(2,'2022-12-08 03:06:12.140','2023-06-07 09:13:11.196',NULL,'Analysis',1,2,'analysis','/dashboard/analysis/index','','dashboard:analyse',100,'ant-design:alert-twotone','分析页',2,2,1,1,'','/api.server.v1.Api/DashboardAnalyse','','',''),(3,'2022-12-08 03:50:55.084','2022-12-08 06:35:34.566',NULL,'UserCreate',10,3,'','','','system:user:create',100,'ant-design:user-add-outlined','用户新增',2,2,2,1,'','/api.server.v1.Api/CreateUser','','',''),(6,'2022-12-08 06:18:26.477','2023-06-07 09:32:17.040',NULL,'UserDelete',10,3,'','','','system:user:delete',100,'ant-design:delete-twotone','用户删除',2,2,2,1,'','/api.server.v1.Api/DeleteUser','','',''),(7,'2022-12-08 06:19:15.292','2023-06-07 09:31:10.872',NULL,'UserUpdate',10,3,'','','','system:user:update',100,'ant-design:upload-outlined','用户修改',2,2,2,1,'','/api.server.v1.Api/UpdateUser','','',''),(8,'2022-12-08 07:40:29.095','2023-06-07 09:12:25.080',NULL,'Workbench',1,2,'workbench','/dashboard/workbench/index','','dashboard:workbench',100,'','工作台',2,2,2,1,'http://localhost','/api.server.v1.Api/DashboardWorkbench','','',''),(9,'2022-12-08 07:41:41.753','2023-06-07 08:25:32.590',NULL,'System',0,1,'/system','LAYOUT','','',3,'ion:settings-outline','系统管理',2,0,0,0,'','','','',''),(10,'2022-12-08 07:45:05.125','2023-06-07 09:34:03.688',NULL,'UserManagement',9,2,'user','/core/user/index','','system:user:list',100,'','用户管理',2,2,2,1,'','/api.server.v1.Api/ListUser','','',''),(11,'2022-12-08 07:47:50.582','2023-06-07 09:33:36.044',NULL,'UserDetail',10,2,'user_detail/:id','/core/user/UserDetail','','system:user:detail',100,'','用户详情',1,2,2,1,'','/api.server.v1.Api/GetUser','','',''),(12,'2022-12-08 07:48:47.697','2023-06-14 11:54:24.403',NULL,'RoleManagement',9,2,'role','/core/role/index','','system:role:list',100,'','角色管理',2,2,2,1,'','/api.server.v1.Api/ListRole','','',''),(13,'2022-12-08 07:49:18.253','2023-06-14 12:18:37.540',NULL,'MenuManagement',17,2,'menu','/core/menu/index','','system:menu:list:tree',100,'','菜单管理',2,2,2,1,'','/api.server.v1.Api/ListMenuTree','','',''),(14,'2022-12-08 07:49:43.133','2023-06-14 12:16:24.148',NULL,'DeptManagement',9,2,'dept','/core/dept/index','','system:dept:list:tree',100,'','部门管理',2,2,2,1,'','/api.server.v1.Api/ListDeptTree','','',''),(15,'2022-12-08 07:50:54.820','2023-06-07 09:53:10.971',NULL,'ChangePassword',10,2,'change_password/:id','/core/password/index','','system:user:password:update',100,'','用户密码修改',1,2,2,1,'','/api.server.v1.Api/UpdateUserPassword','','',''),(16,'2023-01-10 02:12:34.805','2023-06-07 09:16:48.420',NULL,'PostManagement',9,2,'post','/core/post/index','','system:post:list',100,'','岗位管理',2,2,2,1,'','/api.server.v1.Api/ListPost','','',''),(17,'2023-01-11 03:22:07.673','2023-06-07 09:13:47.735',NULL,'Middle',0,1,'/middle','LAYOUT','','',2,'ant-design:setting-outlined','中台管理',2,0,0,0,'','','','',''),(18,'2023-01-11 03:24:59.502','2023-06-14 12:18:57.749',NULL,'DomainManagement',17,2,'domain','/core/domain/index','','system:domain:list:tree',100,'','租户管理',2,2,2,1,'','/api.server.v1.Api/ListDomainTree','','',''),(20,'2022-12-08 03:06:12.140','2023-05-25 07:56:11.708',NULL,'Step',0,2,'step','/demo/page/form/step/index','','step',5,'ant-design:alert-twotone','步骤演示',2,2,2,1,'http://localhost','','','',''),(21,'2023-02-14 09:36:07.747','2023-05-25 08:20:31.015',NULL,'Doc',0,2,'/doc','/sys/iframe/FrameBlank','','',4,'','文档',2,2,2,2,'https://doc.vvbin.cn','','','',''),(22,'2023-02-14 10:59:18.875','2023-06-07 09:30:15.118',NULL,'RoleCreate',12,3,'','','','system:role:create',100,'','角色新增',2,0,0,0,'','/api.server.v1.Api/CreateRole','','',''),(23,'2023-05-30 06:23:37.080','2023-06-08 02:12:17.689',NULL,'RoleMenuHandle',12,3,'','','','system:role:menu:handle',100,'','角色菜单关联',2,1,1,1,'','/api.server.v1.Api/HandleRoleMenu','','',''),(25,'2023-05-30 07:32:48.084','2023-06-08 02:12:29.635',NULL,'RoleRoleDataHandle',12,3,'','','','system:role:dataScope:handle',100,'','数据范围关联',2,0,0,0,'','/api.server.v1.Api/HandleRoleDataScope','','',''),(27,'2022-12-08 03:50:55.084','2022-12-08 06:35:34.566',NULL,'DomainCreate',18,3,'','','','system:domain:create',100,'ant-design:user-add-outlined','租户新增',2,2,2,1,'','/api.server.v1.Api/CreateDomain','','',''),(28,'2022-12-08 05:56:56.234','2022-12-08 06:35:43.458','2023-09-08 18:09:16.000','DomainEdit',18,3,'','','','system:domain:update',100,'ant-design:edit-filled','租户修改',2,2,2,1,'','/api.server.v1.Api/UpdateDomain','','',''),(29,'2022-12-08 06:18:26.477','2023-06-07 09:32:17.040',NULL,'DomainDelete',18,3,'','','','system:domain:delete',100,'ant-design:delete-twotone','租户删除',2,2,2,1,'','/api.server.v1.Api/DeleteDomain','','',''),(30,'2022-12-08 06:19:15.292','2023-06-07 09:31:10.872',NULL,'DomainUpdate',18,3,'','','','system:domain:update',100,'ant-design:upload-outlined','租户修改',2,2,2,1,'','/api.server.v1.Api/CreateDomain','','',''),(31,'2022-12-08 03:50:55.084','2022-12-08 06:35:34.566',NULL,'PostCreate',16,3,'','','','system:post:create',100,'ant-design:user-add-outlined','岗位新增',2,2,2,1,'','/api.server.v1.Api/CreatePost','','',''),(32,'2022-12-08 05:56:56.234','2022-12-08 06:35:43.458','2023-09-08 18:08:28.000','PostEdit',16,3,'','','','system:post:update',100,'ant-design:edit-filled','岗位修改',2,2,2,1,'','/api.server.v1.Api/UpdatePost','','',''),(33,'2022-12-08 06:18:26.477','2023-06-07 09:32:17.040',NULL,'PostDelete',16,3,'','','','system:post:delete',100,'ant-design:delete-twotone','岗位删除',2,2,2,1,'','/api.server.v1.Api/DeletePost','','',''),(34,'2022-12-08 06:19:15.292','2023-06-07 09:31:10.872',NULL,'PostUpdate',16,3,'','','','system:post:update',100,'ant-design:upload-outlined','岗位修改',2,2,2,1,'','/api.server.v1.Api/CreatePost','','',''),(36,'2022-12-08 05:56:56.234','2022-12-08 06:35:43.458','2023-09-08 18:07:19.000','RoleEdit',12,3,'','','','system:role:update',100,'ant-design:edit-filled','角色修改',2,2,2,1,'','/api.server.v1.Api/UpdateRole','','',''),(37,'2022-12-08 06:18:26.477','2023-06-07 09:32:17.040',NULL,'RoleDelete',12,3,'','','','system:role:delete',100,'ant-design:delete-twotone','角色删除',2,2,2,1,'','/api.server.v1.Api/DeleteRole','','',''),(38,'2022-12-08 06:19:15.292','2023-06-07 09:31:10.872',NULL,'RoleUpdate',12,3,'','','','system:role:update',100,'ant-design:upload-outlined','角色修改',2,2,2,1,'','/api.server.v1.Api/CreateRole','','',''),(39,'2022-12-08 03:50:55.084','2022-12-08 06:35:34.566',NULL,'DeptCreate',14,3,'','','','system:dept:create',100,'ant-design:user-add-outlined','部门新增',2,2,2,1,'','/api.server.v1.Api/CreateDept','','',''),(40,'2022-12-08 05:56:56.234','2022-12-08 06:35:43.458','2023-09-08 18:09:43.000','DeptEdit',14,3,'','','','system:dept:update',100,'ant-design:edit-filled','部门修改',2,2,2,1,'','/api.server.v1.Api/UpdateDept','','',''),(41,'2022-12-08 06:18:26.477','2023-06-07 09:32:17.040',NULL,'DeptDelete',14,3,'','','','system:dept:delete',100,'ant-design:delete-twotone','部门删除',2,2,2,1,'','/api.server.v1.Api/DeleteDept','','',''),(42,'2022-12-08 06:19:15.292','2023-06-07 09:31:10.872',NULL,'DeptUpdate',14,3,'','','','system:dept:update',100,'ant-design:upload-outlined','部门修改',2,2,2,1,'','/api.server.v1.Api/CreateDept','','',''),(43,'2022-12-08 03:50:55.084','2022-12-08 06:35:34.566',NULL,'MenuCreate',13,3,'','','','system:menu:create',100,'ant-design:user-add-outlined','菜单新增',2,2,2,1,'','/api.server.v1.Api/CreateMenu','','',''),(44,'2022-12-08 05:56:56.234','2022-12-08 06:35:43.458','2023-09-08 18:08:38.000','MenuEdit',13,3,'','','','system:menu:update',100,'ant-design:edit-filled','菜单修改',2,2,2,1,'','/api.server.v1.Api/UpdateMenu','','',''),(45,'2022-12-08 06:18:26.477','2023-06-07 09:32:17.040',NULL,'MenuDelete',13,3,'','','','system:menu:delete',100,'ant-design:delete-twotone','菜单删除',2,2,2,1,'','/api.server.v1.Api/DeleteMenu','','',''),(46,'2022-12-08 06:19:15.292','2023-06-07 09:31:10.872',NULL,'MenuUpdate',13,3,'','','','system:menu:update',100,'ant-design:upload-outlined','菜单修改',2,2,2,1,'','/api.server.v1.Api/CreateMenu','','',''),(47,'2023-06-08 01:55:18.460','2023-06-14 12:18:46.028',NULL,'DictManagement',17,2,'dict','/core/dict/index','','system:dict:list',100,'','字典管理',2,2,2,1,'','/api.server.v1.Api/ListDict','','',''),(48,'2022-12-08 03:50:55.084','2022-12-08 06:35:34.566',NULL,'DictCreate',47,3,'','','','system:dict:create',100,'ant-design:user-add-outlined','字典新增',2,2,2,1,'','/api.server.v1.Api/CreateDict','','',''),(49,'2022-12-08 05:56:56.234','2022-12-08 06:35:43.458',NULL,'DictEdit',47,3,'','','','system:dict:update',100,'ant-design:edit-filled','字典修改',2,2,2,1,'','/api.server.v1.Api/UpdateDict','','',''),(50,'2022-12-08 06:18:26.477','2023-06-07 09:32:17.040',NULL,'DictDelete',47,3,'','','','system:dict:delete',100,'ant-design:delete-twotone','字典删除',2,2,2,1,'','/api.server.v1.Api/DeleteDict','','',''),(51,'2022-12-08 06:19:15.292','2023-06-07 09:31:10.872',NULL,'DictUpdate',47,3,'','','','system:dict:update',100,'ant-design:upload-outlined','字典修改',2,2,2,1,'','/api.server.v1.Api/CreateDict','','',''),(52,'2023-06-08 02:01:18.401','2023-06-08 10:43:42.856',NULL,'DictDataManagement',47,2,'data/:dict_type','/core/dict/data/index','','system:dict:data:list',100,'','字典数据管理',1,2,2,1,'','/api.server.v1.Api/ListDictData','','',''),(53,'2022-12-08 03:50:55.084','2022-12-08 06:35:34.566',NULL,'DictDataCreate',52,3,'','','','system:dict:data:create',100,'ant-design:user-add-outlined','字典数据新增',2,2,2,1,'','/api.server.v1.Api/CreateDictData','','',''),(54,'2022-12-08 05:56:56.234','2022-12-08 06:35:43.458',NULL,'DictDataUpdate',52,3,'','','','system:dict:data:update',100,'ant-design:edit-filled','字典数据修改',2,2,2,1,'','/api.server.v1.Api/UpdateDictData','','',''),(55,'2022-12-08 06:18:26.477','2023-06-07 09:32:17.040',NULL,'DictDataDelete',52,3,'','','','system:dict:data:delete',100,'ant-design:delete-twotone','字典数据删除',2,2,2,1,'','/api.server.v1.Api/DeleteDictData','','',''),(56,'2022-12-08 06:19:15.292','2023-06-15 07:18:17.267',NULL,'DictDataDetail',52,3,'','','','system:dict:data:detail',100,'ant-design:upload-outlined','字典数据详情',2,2,2,1,'','/api.server.v1.Api/UpdateDictData','','',''),(57,'2023-06-13 08:33:50.857','2023-06-13 08:35:00.893',NULL,'Base',0,1,'/base','LAYOUT','','',1,'','基础权限管理',1,1,1,1,'','','','',''),(58,'2023-06-13 08:38:40.107','2023-06-15 07:16:52.955',NULL,'GetUserInfo',57,3,'','','','base:user:info',1,'','登录用户信息',2,1,1,1,'','/api.server.v1.Api/GetUserInfo','','',''),(59,'2023-06-13 08:40:00.187','2023-06-15 07:17:25.643',NULL,'ListUserRole',57,3,'','','','base:user:role:list',100,'','登录用户角色列表',2,1,1,1,'','/api.server.v1.Api/ListUserRole','','',''),(60,'2023-06-13 08:42:32.796','2023-06-15 07:17:18.947',NULL,'ListUserRoleMenuRouterTree',57,3,'','','','base:user:role:router',100,'','登录用户角色的路由列表',2,1,1,1,'','/api.server.v1.Api/ListUserRoleMenuRouterTree','','',''),(61,'2023-06-13 08:44:12.571','2023-06-15 07:17:11.426',NULL,'ListUserRolePermission',57,3,'','','','base:user:role:permission',100,'','登录用户角色权限列表',2,1,1,1,'','/api.server.v1.Api/ListUserRolePermission','','',''),(62,'2023-06-13 08:46:27.787','2023-06-13 08:48:35.260',NULL,'RoleDetail',12,3,'','','','system:role:detail',100,'','角色详情',2,1,1,1,'','/api.server.v1.Api/GetRole','','',''),(63,'2023-06-13 08:46:27.787','2023-06-13 09:02:48.140',NULL,'DeptDetail',14,3,'','','','system:dept:detail',100,'','部门详情',2,1,1,1,'','/api.server.v1.Api/GetDept','','',''),(64,'2023-06-13 08:46:27.787','2023-06-13 09:02:37.851',NULL,'PostDetail',16,3,'','','','system:post:detail',100,'','岗位详情',2,1,1,1,'','/api.server.v1.Api/GetPost','','',''),(65,'2023-06-13 08:46:27.787','2023-06-13 09:02:04.641',NULL,'MenuDetail',13,3,'','','','system:menu:detail',100,'','菜单详情',2,1,1,1,'','/api.server.v1.Api/GetMenu','','',''),(66,'2023-06-13 08:46:27.787','2023-06-13 09:02:25.775',NULL,'DomainDetail',18,3,'','','','system:domain:detail',100,'','租户详情',2,1,1,1,'','/api.server.v1.Api/GetDomain','','',''),(67,'2023-06-13 08:46:27.787','2023-06-13 08:52:41.324',NULL,'DictDetail',47,3,'','','','system:dict:detail',100,'','字典详情',2,1,1,1,'','/api.server.v1.Api/GetRole','','',''),(68,'2023-06-15 06:59:24.780','2023-06-15 06:59:24.780',NULL,'ExistUserName',10,3,'','','','system:user:name:exist',100,'','用户名称重复查询',2,1,1,1,'','/api.server.v1.Api/ExistUserName','','',''),(69,'2023-06-15 07:16:09.704','2023-06-15 07:16:33.846',NULL,'Logout',57,3,'','','','base:user:loginout',100,'','登录用户退出登录',2,1,1,1,'','/api.server.v1.Api/Logout','','',''),(70,'2023-06-15 08:06:35.974','2023-06-15 08:06:35.974',NULL,'LoginTest',57,3,'','','','base:user:test',100,'','登录用户测试',2,1,1,1,'','/api.server.v1.Api/LoginTest','','',''),(71,'2023-09-08 09:03:14.998','2023-09-08 09:17:44.009',NULL,'DomainPackageDetail',73,3,'','','','system:domain:package:detail',100,'','租户详情',2,1,1,1,'','/api.server.v1.Api/GetDomainPackage','','',''),(73,'2023-01-11 03:24:59.502','2023-09-08 09:12:31.965',NULL,'DomainPackageManagement',18,2,'package','/core/domain/package/index','','system:domain:package:list',100,'','租户套餐',1,2,2,1,'','/api.server.v1.Api/ListDomainPackage','','',''),(74,'2022-12-08 06:19:15.292','2023-06-07 09:31:10.872',NULL,'DomainPackageUpdate',73,3,'','','','system:domain:package:update',100,'ant-design:upload-outlined','套餐修改',2,2,2,1,'','/api.server.v1.Api/CreateDomainPackage','','',''),(75,'2022-12-08 05:56:56.234','2022-12-08 06:35:43.458','2023-09-08 18:10:02.000','DomainPackageEdit',73,3,'','','','system:domain:package:update',100,'ant-design:edit-filled','套餐修改',2,2,2,1,'','/api.server.v1.Api/UpdateDomainPackage','','',''),(76,'2022-12-08 06:18:26.477','2023-06-07 09:32:17.040',NULL,'DomainPackageDelete',73,3,'','','','system:domain:package:delete',100,'ant-design:delete-twotone','套餐删除',2,2,2,1,'','/api.server.v1.Api/DeleteDomainPackage','','',''),(77,'2022-12-08 03:50:55.084','2022-12-08 06:35:34.566',NULL,'DomainPackageCreate',73,3,'','','','system:domain:package:create',100,'ant-design:user-add-outlined','套餐新增',2,2,2,1,'','/api.server.v1.Api/CreateDomainPackage','','','');
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
  `domain_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '租户ID',
  `name` varchar(255) NOT NULL COMMENT '岗位名称',
  `sort` int(10) NOT NULL DEFAULT '100' COMMENT '排序',
  `state` tinyint(1) NOT NULL DEFAULT '1' COMMENT '岗位状态 0 未指定  1 启用 2 停用',
  `remarks` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `code` varchar(100) NOT NULL DEFAULT '' COMMENT '岗位编码',
  `creator` varchar(64) NOT NULL DEFAULT '' COMMENT '创建者',
  `updater` varchar(64) NOT NULL DEFAULT '' COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_sys_posts_state` (`state`),
  KEY `idx_sys_posts_deleted_at` (`deleted_at`),
  KEY `idx_domain_id_data` (`domain_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_posts`
--

LOCK TABLES `sys_posts` WRITE;
/*!40000 ALTER TABLE `sys_posts` DISABLE KEYS */;
INSERT INTO `sys_posts` VALUES (1,'2023-07-07 11:01:52.402','2023-07-07 11:01:53.452',NULL,1,'董事长',100,1,'董事长','1612708008079396865','',''),(2,'2023-07-07 11:02:05.695','2023-07-07 11:02:06.241',NULL,1,'总经理',99,1,'总经理职位','1612708008079396864','',''),(3,'2023-03-24 09:49:40.407','2023-03-24 09:49:42.009',NULL,1,'秘书长',88,1,'秘书长','1639202775373975552','','');
/*!40000 ALTER TABLE `sys_posts` ENABLE KEYS */;
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
INSERT INTO `sys_role_depts` VALUES (1,1),(2,1),(3,1),(4,1),(5,1),(5,5),(6,1),(7,1);
/*!40000 ALTER TABLE `sys_role_depts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_menus`
--

DROP TABLE IF EXISTS `sys_role_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role_menus` (
  `sys_role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `sys_menu_id` bigint(20) unsigned NOT NULL COMMENT '菜单ID',
  UNIQUE KEY `idx_role_menu_role_id_menu_id` (`sys_role_id`,`sys_menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_menus`
--

LOCK TABLES `sys_role_menus` WRITE;
/*!40000 ALTER TABLE `sys_role_menus` DISABLE KEYS */;
INSERT INTO `sys_role_menus` VALUES (2,1),(2,2),(2,3),(2,6),(2,7),(2,8),(2,9),(2,10),(2,11),(2,12),(2,13),(2,14),(2,15),(2,16),(2,17),(2,18),(2,20),(2,21),(2,22),(2,23),(2,25),(2,27),(2,28),(2,29),(2,30),(2,31),(2,32),(2,33),(2,34),(2,36),(2,37),(2,38),(2,39),(2,40),(2,41),(2,42),(2,43),(2,44),(2,45),(2,46),(2,47),(2,48),(2,49),(2,50),(2,51),(2,52),(2,53),(2,54),(2,55),(2,56),(2,58),(2,59),(2,60),(2,61),(2,62),(2,63),(2,64),(2,65),(2,66),(2,67),(2,68),(2,69),(3,3),(3,6),(3,7),(3,9),(3,10),(3,11),(3,12),(3,14),(3,15),(3,16),(3,21),(3,22),(3,23),(3,25),(3,31),(3,32),(3,33),(3,34),(3,36),(3,37),(3,38),(3,39),(3,40),(3,41),(3,42),(3,57),(3,58),(3,59),(3,60),(3,61),(3,62),(3,63),(3,64),(3,68),(3,69),(4,20),(4,21),(4,57),(4,58),(4,59),(4,60),(4,61),(5,2),(5,57),(5,58),(5,59),(5,60),(5,61),(5,69);
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
  `domain_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '租户ID',
  `name` varchar(255) NOT NULL COMMENT '角色名称',
  `parent_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '父角色ID',
  `default_router` varchar(255) NOT NULL DEFAULT '/dashboard' COMMENT '默认路由',
  `sort` int(10) NOT NULL DEFAULT '100' COMMENT '排序',
  `state` tinyint(1) NOT NULL DEFAULT '1' COMMENT '角色状态 0 未指定  1 启用 2 停用',
  `remarks` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `data_scope` tinyint(2) NOT NULL DEFAULT '1' COMMENT '数据范围（0：未指定 1：本人数据权限 2：全部数据权限 3：本部门数据权限 4：本部门及以下数据权限 5：自定部门数据权限 ）',
  `menu_check_strictly` tinyint(2) NOT NULL DEFAULT '1' COMMENT '菜单树选择项是否关联显示',
  `dept_check_strictly` tinyint(2) NOT NULL DEFAULT '1' COMMENT '部门树选择项是否关联显示',
  `creator` varchar(64) NOT NULL DEFAULT '' COMMENT '创建者',
  `updater` varchar(64) NOT NULL DEFAULT '' COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_sys_roles_deleted_at` (`deleted_at`),
  KEY `idx_domain_id_data` (`domain_id`,`name`),
  KEY `idx_sys_roles_state` (`state`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_roles`
--

LOCK TABLES `sys_roles` WRITE;
/*!40000 ALTER TABLE `sys_roles` DISABLE KEYS */;
INSERT INTO `sys_roles` VALUES (1,'2023-06-14 11:25:53.290','2023-06-14 11:25:53.290',NULL,1,'超级管理员',0,'/dashboard',100,1,'',2,1,1,'',''),(2,'2023-06-13 09:04:10.847','2023-06-16 10:55:41.445',NULL,1,'高级管理员',0,'/dashboard',100,1,'',2,1,1,'',''),(3,'2023-06-13 08:18:57.368','2023-06-15 07:39:38.927',NULL,1,'普通管理员',0,'/dashboard',0,1,'',3,1,1,'',''),(4,'2023-06-13 08:26:11.172','2023-06-14 12:06:55.799',NULL,1,'员工管理员',0,'/dashboard',0,1,'',1,1,1,'',''),(5,'2023-07-11 08:43:32.738','2023-07-11 08:43:36.373',NULL,1,'测试管理员',0,'/dashboard',100,1,'',5,1,1,'','');
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
  `sys_dict_id` bigint(20) unsigned DEFAULT NULL COMMENT '主键ID',
  PRIMARY KEY (`sys_post_id`,`sys_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_posts`
--

LOCK TABLES `sys_user_posts` WRITE;
/*!40000 ALTER TABLE `sys_user_posts` DISABLE KEYS */;
INSERT INTO `sys_user_posts` VALUES (1,2,NULL),(2,1,NULL),(2,3,NULL),(2,13,NULL),(3,1,NULL),(3,4,NULL),(3,13,NULL);
/*!40000 ALTER TABLE `sys_user_posts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_roles`
--

DROP TABLE IF EXISTS `sys_user_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_user_roles` (
  `sys_role_id` bigint(20) unsigned NOT NULL COMMENT '主键ID',
  `sys_user_id` bigint(20) unsigned NOT NULL COMMENT '主键ID',
  PRIMARY KEY (`sys_role_id`,`sys_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_roles`
--

LOCK TABLES `sys_user_roles` WRITE;
/*!40000 ALTER TABLE `sys_user_roles` DISABLE KEYS */;
INSERT INTO `sys_user_roles` VALUES (1,1),(2,2),(3,3),(3,12),(4,4),(5,13);
/*!40000 ALTER TABLE `sys_user_roles` ENABLE KEYS */;
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
  `domain_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '租户ID',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `name` varchar(255) NOT NULL COMMENT '用户名称',
  `nick_name` varchar(255) NOT NULL DEFAULT '' COMMENT '昵称',
  `real_name` varchar(100) NOT NULL DEFAULT '' COMMENT '实名',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(255) NOT NULL DEFAULT '' COMMENT '密码加盐',
  `birthday` datetime DEFAULT NULL COMMENT '生日',
  `gender` tinyint(1) NOT NULL DEFAULT '1' COMMENT '性别 0 未指定 1 男 2 女',
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `dept_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '部门ID',
  `state` tinyint(1) NOT NULL DEFAULT '1' COMMENT '用户状态 0 未指定  1 启用 2 停用',
  `last_use_role_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '最后使用角色ID',
  `last_login_at` datetime DEFAULT NULL COMMENT '最后登录时间',
  `last_login_ip` varchar(50) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `remarks` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `creator` varchar(64) NOT NULL DEFAULT '' COMMENT '创建者',
  `updater` varchar(64) NOT NULL DEFAULT '' COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_users_mobile_email` (`phone`,`email`),
  KEY `idx_sys_users_state` (`state`),
  KEY `idx_sys_users_deleted_at` (`deleted_at`),
  KEY `idx_users_name_nick_name_real_name` (`name`,`real_name`,`nick_name`),
  KEY `idx_users_phone_email` (`phone`,`email`),
  KEY `idx_domain_id_data` (`domain_id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_users`
--

LOCK TABLES `sys_users` WRITE;
/*!40000 ALTER TABLE `sys_users` DISABLE KEYS */;
INSERT INTO `sys_users` VALUES (1,'2022-11-05 06:17:21.364','2023-09-08 10:05:58.277',NULL,1,'','xiong','xiong','熊龙军','$2a$10$065qdXsShtSEw8YizPKTxeVB59Q79/lfE3hC8xnn95UpNOHThlgi2','','2023-06-02 00:00:00',1,'18584565115','',5,1,1,'2023-09-08 10:05:58','0.0.0.0','','',''),(2,'2022-12-24 02:01:13.833','2023-06-15 03:52:53.837',NULL,1,'','jayden1','jayden1','熊军','$2a$10$JePrEVFbXjYv/BsDsowslOzjLqzhwEy1CJ1YomxBIyWPmSnImhxP2','','2023-06-02 00:00:00',1,'18584565116','jayden@qq.com',1,1,2,'2023-06-13 08:55:14','0.0.0.0','','',''),(3,'2023-06-03 06:14:04.418','2023-06-15 07:39:12.829',NULL,1,'','jingong','jd','熊1军','$2a$10$065qdXsShtSEw8YizPKTxeVB59Q79/lfE3hC8xnn95UpNOHThlgi2','',NULL,1,'19584565121','737043980@qq.com',2,1,3,'2023-06-15 07:39:13','0.0.0.0','','',''),(4,'2023-06-03 06:24:42.648','2023-06-15 07:21:50.118',NULL,1,'','caiwu','caiwu','财务','$2a$10$065qdXsShtSEw8YizPKTxeVB59Q79/lfE3hC8xnn95UpNOHThlgi2','','2023-06-03 00:00:00',2,'15854856512','',7,2,4,'2023-06-15 07:21:50','0.0.0.0','','',''),(12,'2023-06-14 12:37:24.368','2023-06-15 03:55:56.738',NULL,1,'','kaifa02','kaifa02','开发2','$2a$10$065qdXsShtSEw8YizPKTxeVB59Q79/lfE3hC8xnn95UpNOHThlgi2','',NULL,1,'19584765121','',2,1,0,NULL,'','','',''),(13,'2023-06-14 12:38:38.311','2023-06-15 07:09:47.630',NULL,1,'','caiwu02','caiwu02','财务2','$2a$10$zBQLgpXutxe5YmUnPJuDs.k1jNPVBFNHaKRNrzSUtIGFo1iN1LWNO','',NULL,2,'12565845896','',7,1,5,'2023-06-15 07:09:48','0.0.0.0','','','');
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

-- Dump completed on 2023-09-09 18:21:42
