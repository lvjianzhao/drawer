/*
 Navicat Premium Data Transfer

 Source Server         : local_mysql
 Source Server Type    : MySQL
 Source Server Version : 50735
 Source Host           : localhost:3306
 Source Schema         : init_drawer

 Target Server Type    : MySQL
 Target Server Version : 50735
 File Encoding         : 65001

 Date: 11/07/2023 16:52:25
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 879 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (858, 'p', '1', '/api/*', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (853, 'p', '1', '/api/addApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (855, 'p', '1', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (856, 'p', '1', '/api/editApi', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (854, 'p', '1', '/api/getApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (857, 'p', '1', '/api/getElTreeApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (832, 'p', '1', '/base/captcha', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (833, 'p', '1', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (852, 'p', '1', '/casbin/editCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (860, 'p', '1', '/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (861, 'p', '1', '/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (862, 'p', '1', '/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (863, 'p', '1', '/customer/*', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (864, 'p', '1', '/customer/all', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (868, 'p', '1', '/customer/env', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (866, 'p', '1', '/customer/env', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (865, 'p', '1', '/customer/env', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (867, 'p', '1', '/customer/env', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (874, 'p', '1', '/customer/env/:customerId', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (872, 'p', '1', '/customer/event', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (869, 'p', '1', '/customer/event', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (870, 'p', '1', '/customer/event', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (871, 'p', '1', '/customer/event', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (873, 'p', '1', '/customer/summary', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (859, 'p', '1', '/jwt/joinInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (875, 'p', '1', '/license', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (876, 'p', '1', '/license', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (877, 'p', '1', '/license', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (878, 'p', '1', '/license/:licenseId', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (848, 'p', '1', '/menu/addMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (850, 'p', '1', '/menu/deleteMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (849, 'p', '1', '/menu/editMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (851, 'p', '1', '/menu/getElTreeMenus', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (847, 'p', '1', '/menu/getMenus', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (843, 'p', '1', '/role/addRole', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (844, 'p', '1', '/role/deleteRole', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (845, 'p', '1', '/role/editRole', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (846, 'p', '1', '/role/editRoleMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (842, 'p', '1', '/role/getRoles', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (837, 'p', '1', '/user/addUser', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (836, 'p', '1', '/user/deleteUser', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (838, 'p', '1', '/user/editUser', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (834, 'p', '1', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (841, 'p', '1', '/user/getUserInfo', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (835, 'p', '1', '/user/getUsers', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (839, 'p', '1', '/user/modifyPass', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (840, 'p', '1', '/user/switchActive', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (225, 'p', '2', '/api/addApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (227, 'p', '2', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (228, 'p', '2', '/api/editApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (226, 'p', '2', '/api/getApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (229, 'p', '2', '/api/getElTreeApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (206, 'p', '2', '/base/captcha', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (207, 'p', '2', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (224, 'p', '2', '/casbin/editCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (231, 'p', '2', '/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (232, 'p', '2', '/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (233, 'p', '2', '/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (234, 'p', '2', '/customer/*', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (230, 'p', '2', '/jwt/joinInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (220, 'p', '2', '/menu/addMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (222, 'p', '2', '/menu/deleteMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (221, 'p', '2', '/menu/editMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (223, 'p', '2', '/menu/getElTreeMenus', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (219, 'p', '2', '/menu/getMenus', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (215, 'p', '2', '/role/addRole', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (216, 'p', '2', '/role/deleteRole', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (217, 'p', '2', '/role/editRole', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (218, 'p', '2', '/role/editRoleMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (214, 'p', '2', '/role/getRoles', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (211, 'p', '2', '/user/addUser', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (210, 'p', '2', '/user/deleteUser', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (212, 'p', '2', '/user/editUser', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (208, 'p', '2', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (209, 'p', '2', '/user/getUsers', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (213, 'p', '2', '/user/modifyPass', 'POST', '', '', '');

-- ----------------------------
-- Table structure for customer
-- ----------------------------
DROP TABLE IF EXISTS `customer`;
CREATE TABLE `customer`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `name` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '客户名称',
  `service_start_date` date NOT NULL DEFAULT '2023-01-01' COMMENT '维保起始日期',
  `service_end_date` date NOT NULL DEFAULT '2023-01-01' COMMENT '维保起始日期',
  `service_status` bigint(20) UNSIGNED NOT NULL COMMENT '维保状态: (0: 已过期,1: 维保中)',
  `description` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '描述信息',
  `business_name` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '商务名称',
  `type` bigint(20) UNSIGNED NOT NULL DEFAULT 1 COMMENT '客户类型(1:POC 2:正式)',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE,
  UNIQUE INDEX `name_2`(`name`) USING BTREE,
  INDEX `idx_customer_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 87 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of customer
-- ----------------------------

-- ----------------------------
-- Table structure for customer_env
-- ----------------------------
DROP TABLE IF EXISTS `customer_env`;
CREATE TABLE `customer_env`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `env` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '环境',
  `vpn` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'VPN信息',
  `customer_id` bigint(20) UNSIGNED NOT NULL COMMENT '客户ID',
  `deploy_report` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '部署报告',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_customerId_env`(`customer_id`, `env`) USING BTREE,
  INDEX `idx_customer_env_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_customer_env_customer_id`(`customer_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of customer_env
-- ----------------------------

-- ----------------------------
-- Table structure for customer_event
-- ----------------------------
DROP TABLE IF EXISTS `customer_event`;
CREATE TABLE `customer_event`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `customer_id` bigint(20) UNSIGNED NOT NULL COMMENT '客户ID',
  `event_time` datetime(0) NOT NULL COMMENT '事件发生日期',
  `event_content` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '事件内容',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_customer_event_customer_id`(`customer_id`) USING BTREE,
  INDEX `idx_customer_event_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 46 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of customer_event
-- ----------------------------

-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklists`;
CREATE TABLE `jwt_blacklists`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT 'jwt',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_jwt_blacklists_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jwt_blacklists
-- ----------------------------
INSERT INTO `jwt_blacklists` VALUES (1, '2023-03-11 10:11:53', '2023-03-11 10:11:53', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVJZCI6MSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjc4NjE1ODU1LCJuYmYiOjE2Nzg1Mjk0NTV9.-Ww9hn9RFISVSFpRidVSUvkIjsW2FZGJum710fhnMRA');
INSERT INTO `jwt_blacklists` VALUES (2, '2023-03-13 08:19:48', '2023-03-13 08:19:48', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVJZCI6MSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjc4NzYyMDgzLCJuYmYiOjE2Nzg2NzU2ODN9.0xZoylhrvlMh7Sn-KTpNKci2pE3DNlfHt1upB5-_jyQ');
INSERT INTO `jwt_blacklists` VALUES (3, '2023-06-22 06:39:25', '2023-06-22 06:39:25', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVJZCI6MSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjg3NDgzMjc4LCJuYmYiOjE2ODczOTY4Nzh9._ySAj85zZBXt_p2l3qN0YzjjC8InFJ-MlLYVQjmLNpM');
INSERT INTO `jwt_blacklists` VALUES (4, '2023-06-22 14:11:12', '2023-06-22 14:11:12', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVJZCI6MSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjg3NTI5NDE2LCJuYmYiOjE2ODc0NDMwMTZ9.9ZAco5evqILkqPkoBCkRmTDVvNaV4_R8rOf3dLtlEaU');
INSERT INTO `jwt_blacklists` VALUES (5, '2023-06-22 15:21:16', '2023-06-22 15:21:16', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVJZCI6MSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjg3NTAyMzczLCJuYmYiOjE2ODc0MTU5NzN9.MpUi9KblujSRhgfF6oy41_d40NeWfyFDYjThkkRqhpI');
INSERT INTO `jwt_blacklists` VALUES (6, '2023-06-22 15:24:45', '2023-06-22 15:24:45', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVJZCI6MSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjg3NTMzNjkwLCJuYmYiOjE2ODc0NDcyOTB9.QbvVL1s6n9Np-k_pvG03FXGEnc1oniNhDdJ66Ye41Uc');
INSERT INTO `jwt_blacklists` VALUES (7, '2023-06-23 01:39:01', '2023-06-23 01:39:01', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVJZCI6MSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjg3NTcwNjk3LCJuYmYiOjE2ODc0ODQyOTd9.4Igs05axltGpqYozcLTYW5EDphdWdZy1betMFDZWAz8');
INSERT INTO `jwt_blacklists` VALUES (8, '2023-06-23 01:40:25', '2023-06-23 01:40:25', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVJZCI6MSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjg3NTcwNzY1LCJuYmYiOjE2ODc0ODQzNjV9.6Ni0Q7OgemT80DQnty1-qkgpulvwi7Gff0wGe8QKCOU');
INSERT INTO `jwt_blacklists` VALUES (9, '2023-06-23 02:43:08', '2023-06-23 02:43:08', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVJZCI6MSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjg3NTczMjQ3LCJuYmYiOjE2ODc0ODY4NDd9.ucEDHhiTtKNs8Js17VMBSvHQp1sHUIlCeHC2J-_zjQU');
INSERT INTO `jwt_blacklists` VALUES (10, '2023-06-23 14:42:58', '2023-06-23 14:42:58', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVJZCI6MSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjg3NTc0NjY4LCJuYmYiOjE2ODc0ODgyNjh9.gRVIwHi7NbDPtnoEiSALkgqc8RyUAN-m-NCpzh9PzdM');
INSERT INTO `jwt_blacklists` VALUES (11, '2023-06-23 14:48:54', '2023-06-23 14:48:54', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVJZCI6MSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjg3NjE3Nzg2LCJuYmYiOjE2ODc1MzEzODZ9.RkaAim68FnyTTKxqXNDq7c55ihi_yZRxwjHN_qH7Z20');
INSERT INTO `jwt_blacklists` VALUES (12, '2023-06-28 08:33:06', '2023-06-28 08:33:06', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVJZCI6MSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjg3OTQ5NTYzLCJuYmYiOjE2ODc4NjMxNjN9.M_E06uVDrUHHePqbIWlctsT-mGy_clIcD3wvl360GgY');
INSERT INTO `jwt_blacklists` VALUES (13, '2023-07-02 06:18:12', '2023-07-02 06:18:12', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVJZCI6MSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjg4Mjk3NjYyLCJuYmYiOjE2ODgyMTEyNjJ9.EZot1HRJg4Coxf-Kgkaw-VhaniTeu497bYR3kPGgs0g');
INSERT INTO `jwt_blacklists` VALUES (14, '2023-07-02 10:51:12', '2023-07-02 10:51:12', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVJZCI6MSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjg4MzY1MTY3LCJuYmYiOjE2ODgyNzg3Njd9.v_QWPlyzAktAUcy21UkjS-BRlx9FPQm-39JVXUEycDY');
INSERT INTO `jwt_blacklists` VALUES (15, '2023-07-02 11:07:00', '2023-07-02 11:07:00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVJZCI6MSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjg4MzgyMzI3LCJuYmYiOjE2ODgyOTU5Mjd9.VHYR9DvIiA20CL_F1_-FMMsyFwz9VSn43tQuuFBNiQI');
INSERT INTO `jwt_blacklists` VALUES (16, '2023-07-02 11:38:48', '2023-07-02 11:38:48', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVJZCI6MSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjg4Mzg0MjE3LCJuYmYiOjE2ODgyOTc4MTd9.O97_PCT5eJmEUbzTWrILfPAFh6A4ipDtou_7nPt6tiM');
INSERT INTO `jwt_blacklists` VALUES (17, '2023-07-02 12:56:10', '2023-07-02 12:56:10', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MiwidXNlcm5hbWUiOiJ6aGFuZ3NhbiIsInJvbGVJZCI6MiwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjg4Mzg4ODQ5LCJuYmYiOjE2ODgzMDI0NDl9.JHViWPme1i_bfP02k74HwonqHYXCWC0igm2rm3VlnOA');
INSERT INTO `jwt_blacklists` VALUES (18, '2023-07-02 12:56:30', '2023-07-02 12:56:30', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MiwidXNlcm5hbWUiOiJ6aGFuZ3NhbiIsInJvbGVJZCI6MiwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjg4Mzg4OTg0LCJuYmYiOjE2ODgzMDI1ODR9.ddmu9zkEt0DlFLC30RlRXElh2B4diq9IbW8T8l85lPE');
INSERT INTO `jwt_blacklists` VALUES (19, '2023-07-02 12:57:55', '2023-07-02 12:57:55', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MiwidXNlcm5hbWUiOiJ6aGFuZ3NhbiIsInJvbGVJZCI6MiwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjg4Mzg5MDMzLCJuYmYiOjE2ODgzMDI2MzN9.Vg4VhCg8xsEVYvlpJ2seF3K-j7OkPd-1gGARjzZNKk4');
INSERT INTO `jwt_blacklists` VALUES (20, '2023-07-02 14:00:11', '2023-07-02 14:00:11', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MiwidXNlcm5hbWUiOiJ6aGFuZ3NhbiIsInJvbGVJZCI6MiwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjg4Mzg5MTA1LCJuYmYiOjE2ODgzMDI3MDV9.rHtUNOFbdEo6IbN6-u0om41z1FGcjvIlNx5tC9HtW_I');
INSERT INTO `jwt_blacklists` VALUES (21, '2023-07-03 05:54:40', '2023-07-03 05:54:40', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVJZCI6MSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjg4NDM1MTYyLCJuYmYiOjE2ODgzNDg3NjJ9.qfxjAPjiB7pxzJuBavzLHDhwglxo-Q04S0ru-1_b92w');
INSERT INTO `jwt_blacklists` VALUES (22, '2023-07-06 02:50:22', '2023-07-06 02:50:22', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVJZCI6MSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZFpsIiwiZXhwIjoxNjg4NjQ5MjIyLCJuYmYiOjE2ODg1NjI4MjJ9.rQJ7VOOuxj7gRGnmjAmTkqbHIK24VMTBDcVpkjfMZ2w');

-- ----------------------------
-- Table structure for license
-- ----------------------------
DROP TABLE IF EXISTS `license`;
CREATE TABLE `license`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `customer_id` bigint(20) UNSIGNED NOT NULL COMMENT '客户ID',
  `env` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '环境',
  `expired_date` datetime(0) NOT NULL COMMENT 'license有效期',
  `license` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT 'license',
  `manager_apis` bigint(20) UNSIGNED NOT NULL COMMENT '管理平台API数',
  `manager_users` bigint(20) UNSIGNED NOT NULL COMMENT '管理平台用户数',
  `manager_projects` bigint(20) UNSIGNED NOT NULL COMMENT '管理平台项目数',
  `manager_integrations` bigint(20) UNSIGNED NOT NULL COMMENT '管理平台集成平台限制',
  `manager_organizations` bigint(20) UNSIGNED NOT NULL COMMENT '管理平台支持组织数',
  `manager_tests` bigint(20) UNSIGNED NOT NULL COMMENT '管理平台测试限制',
  `manager_products` bigint(20) UNSIGNED NOT NULL COMMENT '管理平台发布限制',
  `studio_users` bigint(20) UNSIGNED NOT NULL COMMENT '集成平台用户数',
  `studio_comp_black_list` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '集成组件黑名单',
  `permission_menu` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '权限设置',
  `check_list` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '权限设置复选框',
  `active` bigint(20) UNSIGNED NULL DEFAULT 1 COMMENT 'license是否生效(1:生效中,2:已作废)',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_license_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_license_customer_id`(`customer_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 69 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of license
-- ----------------------------

-- ----------------------------
-- Table structure for role_menus
-- ----------------------------
DROP TABLE IF EXISTS `role_menus`;
CREATE TABLE `role_menus`  (
  `menu_model_id` bigint(20) UNSIGNED NOT NULL,
  `role_model_id` bigint(20) UNSIGNED NOT NULL,
  PRIMARY KEY (`menu_model_id`, `role_model_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_menus
-- ----------------------------
INSERT INTO `role_menus` VALUES (1, 1);
INSERT INTO `role_menus` VALUES (1, 2);
INSERT INTO `role_menus` VALUES (2, 1);
INSERT INTO `role_menus` VALUES (2, 2);
INSERT INTO `role_menus` VALUES (3, 1);
INSERT INTO `role_menus` VALUES (3, 2);
INSERT INTO `role_menus` VALUES (4, 1);
INSERT INTO `role_menus` VALUES (4, 2);
INSERT INTO `role_menus` VALUES (5, 1);
INSERT INTO `role_menus` VALUES (5, 2);
INSERT INTO `role_menus` VALUES (10, 1);
INSERT INTO `role_menus` VALUES (10, 2);
INSERT INTO `role_menus` VALUES (11, 1);
INSERT INTO `role_menus` VALUES (11, 2);
INSERT INTO `role_menus` VALUES (12, 1);
INSERT INTO `role_menus` VALUES (12, 2);
INSERT INTO `role_menus` VALUES (13, 1);
INSERT INTO `role_menus` VALUES (15, 1);
INSERT INTO `role_menus` VALUES (16, 1);
INSERT INTO `role_menus` VALUES (17, 1);
INSERT INTO `role_menus` VALUES (18, 1);
INSERT INTO `role_menus` VALUES (20, 1);

-- ----------------------------
-- Table structure for sys_api
-- ----------------------------
DROP TABLE IF EXISTS `sys_api`;
CREATE TABLE `sys_api`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `path` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'api路径',
  `description` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'api中文描述',
  `api_group` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'api组',
  `method` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_api_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 72 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_api
-- ----------------------------
INSERT INTO `sys_api` VALUES (1, '2023-03-10 06:24:36', '2023-03-10 06:24:36', NULL, '/base/captcha', '获取验证码（必选）', 'base', 'POST');
INSERT INTO `sys_api` VALUES (2, '2023-03-08 06:36:24', '2023-03-09 08:50:20', NULL, '/base/login', '登录（必选）', 'base', 'POST');
INSERT INTO `sys_api` VALUES (3, '2023-03-08 08:56:13', '2023-03-10 07:11:53', NULL, '/user/getUserInfo', '获取用户信息（必选）', 'user', 'GET');
INSERT INTO `sys_api` VALUES (4, '2023-03-08 08:56:54', '2023-03-08 08:56:54', NULL, '/user/getUsers', '获取所有用户', 'user', 'POST');
INSERT INTO `sys_api` VALUES (5, '2023-03-10 06:41:32', '2023-03-10 06:41:32', NULL, '/user/deleteUser', '删除用户', 'user', 'POST');
INSERT INTO `sys_api` VALUES (6, '2023-03-10 06:42:24', '2023-03-10 06:42:24', NULL, '/user/addUser', '添加用户', 'user', 'POST');
INSERT INTO `sys_api` VALUES (7, '2023-03-10 06:47:18', '2023-03-10 06:47:18', NULL, '/user/editUser', '编辑用户', 'user', 'POST');
INSERT INTO `sys_api` VALUES (8, '2023-03-10 06:47:59', '2023-03-10 06:47:59', NULL, '/user/modifyPass', '修改用户密码', 'user', 'POST');
INSERT INTO `sys_api` VALUES (9, '2023-03-10 06:48:43', '2023-03-10 06:48:43', NULL, '/user/switchActive', '切换用户状态', 'user', 'POST');
INSERT INTO `sys_api` VALUES (10, '2023-03-10 06:58:30', '2023-03-10 06:58:30', NULL, '/role/getRoles', '获取所有角色', 'role', 'POST');
INSERT INTO `sys_api` VALUES (11, '2023-03-10 06:59:08', '2023-03-10 06:59:08', NULL, '/role/addRole', '添加角色', 'role', 'POST');
INSERT INTO `sys_api` VALUES (12, '2023-03-10 06:59:54', '2023-03-10 06:59:54', NULL, '/role/deleteRole', '删除角色', 'role', 'POST');
INSERT INTO `sys_api` VALUES (13, '2023-03-10 07:00:14', '2023-03-10 07:00:53', NULL, '/role/editRole', '编辑角色', 'role', 'POST');
INSERT INTO `sys_api` VALUES (14, '2023-03-10 07:01:44', '2023-03-10 07:01:44', NULL, '/role/editRoleMenu', '编辑角色菜单', 'role', 'POST');
INSERT INTO `sys_api` VALUES (15, '2023-03-10 07:14:44', '2023-03-10 07:14:44', NULL, '/menu/getMenus', '获取所有菜单', 'menu', 'GET');
INSERT INTO `sys_api` VALUES (16, '2023-03-10 07:15:25', '2023-03-10 07:15:25', NULL, '/menu/addMenu', '添加菜单', 'menu', 'POST');
INSERT INTO `sys_api` VALUES (17, '2023-03-10 07:15:50', '2023-03-10 07:15:50', NULL, '/menu/editMenu', '编辑菜单', 'menu', 'POST');
INSERT INTO `sys_api` VALUES (18, '2023-03-10 07:16:18', '2023-03-10 07:16:18', NULL, '/menu/deleteMenu', '删除菜单', 'menu', 'POST');
INSERT INTO `sys_api` VALUES (19, '2023-03-10 07:17:13', '2023-03-10 07:17:13', NULL, '/menu/getElTreeMenus', '获取所有菜单（el-tree结构）', 'menu', 'POST');
INSERT INTO `sys_api` VALUES (20, '2023-03-10 07:21:37', '2023-03-10 07:21:37', NULL, '/casbin/editCasbin', '编辑casbin规则', 'casbin', 'POST');
INSERT INTO `sys_api` VALUES (21, '2023-03-10 07:23:21', '2023-03-10 07:33:01', NULL, '/api/addApi', '添加api', 'api', 'POST');
INSERT INTO `sys_api` VALUES (22, '2023-03-10 07:24:00', '2023-03-10 07:24:00', NULL, '/api/getApis', '获取所有api', 'api', 'POST');
INSERT INTO `sys_api` VALUES (23, '2023-03-10 07:24:33', '2023-03-10 07:24:33', NULL, '/api/deleteApi', '删除api', 'api', 'POST');
INSERT INTO `sys_api` VALUES (24, '2023-03-10 07:26:15', '2023-03-10 07:26:15', NULL, '/api/editApi', '编辑api', 'api', 'PUT');
INSERT INTO `sys_api` VALUES (25, '2023-03-10 07:34:08', '2023-03-10 07:35:04', NULL, '/api/getElTreeApis', '获取所有api（el-tree结构）', 'api', 'POST');
INSERT INTO `sys_api` VALUES (27, '2023-03-11 13:05:40', '2023-03-11 13:05:40', NULL, '/jwt/joinInBlacklist', '拉黑token', 'jwt', 'POST');
INSERT INTO `sys_api` VALUES (28, '2023-06-20 10:07:08', '2023-06-21 08:50:59', NULL, '/customer', '获取客户名称', 'customer', 'GET');
INSERT INTO `sys_api` VALUES (29, '2023-06-21 02:38:44', '2023-06-22 10:15:56', NULL, '/customer', '添加客户名称a', 'customer', 'POST');
INSERT INTO `sys_api` VALUES (30, '2023-06-21 08:50:48', '2023-06-21 08:50:48', NULL, '/customer', '修改客户名称', 'customer', 'PUT');
INSERT INTO `sys_api` VALUES (31, '2023-06-21 08:51:30', '2023-06-21 10:23:29', NULL, '/customer/*', '删除客户名称', 'customer', 'DELETE');
INSERT INTO `sys_api` VALUES (32, '2023-06-22 02:22:08', '2023-06-22 02:22:08', NULL, '/api/*', 'api相关的所有get操作', 'api', 'GET');
INSERT INTO `sys_api` VALUES (56, '2023-06-25 10:26:22', '2023-06-25 10:27:31', NULL, '/customer/all', '获取所有客户名称', 'customer', 'GET');
INSERT INTO `sys_api` VALUES (57, '2023-06-26 07:44:39', '2023-06-26 07:44:39', NULL, '/customer/env', '对客户环境信息进行增删改查', 'customer', 'POST');
INSERT INTO `sys_api` VALUES (58, '2023-06-26 07:44:39', '2023-06-26 07:44:39', NULL, '/customer/env', '对客户环境信息进行增删改查', 'customer', 'GET');
INSERT INTO `sys_api` VALUES (59, '2023-06-26 07:44:39', '2023-06-26 07:44:39', NULL, '/customer/env', '对客户环境信息进行增删改查', 'customer', 'PUT');
INSERT INTO `sys_api` VALUES (60, '2023-06-26 07:44:39', '2023-06-26 07:44:39', NULL, '/customer/env', '对客户环境信息进行增删改查', 'customer', 'DELETE');
INSERT INTO `sys_api` VALUES (61, '2023-06-28 08:37:00', '2023-06-28 08:37:00', NULL, '/user/getUserInfo', '获取用户信息（必选）', 'user', 'POST');
INSERT INTO `sys_api` VALUES (62, '2023-07-01 06:23:37', '2023-07-01 06:23:37', NULL, '/license', 'license管理', 'license', 'GET');
INSERT INTO `sys_api` VALUES (63, '2023-07-01 06:23:37', '2023-07-01 06:23:37', NULL, '/license', 'license管理', 'license', 'POST');
INSERT INTO `sys_api` VALUES (64, '2023-07-01 06:23:37', '2023-07-01 06:23:37', NULL, '/license', 'license管理', 'license', 'PUT');
INSERT INTO `sys_api` VALUES (65, '2023-07-01 06:23:37', '2023-07-04 05:30:02', NULL, '/license/:licenseId', 'license管理', 'license', 'DELETE');
INSERT INTO `sys_api` VALUES (66, '2023-07-06 07:35:35', '2023-07-06 07:35:35', NULL, '/customer/event', '对客户事件进行增删改查', 'customer', 'GET');
INSERT INTO `sys_api` VALUES (67, '2023-07-06 07:35:35', '2023-07-06 07:35:35', NULL, '/customer/event', '对客户事件进行增删改查', 'customer', 'POST');
INSERT INTO `sys_api` VALUES (68, '2023-07-06 07:35:35', '2023-07-06 07:35:35', NULL, '/customer/event', '对客户事件进行增删改查', 'customer', 'PUT');
INSERT INTO `sys_api` VALUES (69, '2023-07-06 07:35:35', '2023-07-06 07:35:35', NULL, '/customer/event', '对客户事件进行增删改查', 'customer', 'DELETE');
INSERT INTO `sys_api` VALUES (70, '2023-07-07 21:43:28', '2023-07-07 21:43:28', NULL, '/customer/summary', '获取仪表盘所需的汇总数据', 'customer', 'GET');
INSERT INTO `sys_api` VALUES (71, '2023-07-08 09:40:47', '2023-07-08 09:40:47', NULL, '/customer/env/:customerId', '根据客户id获取其所有的环境类型', 'customer', 'GET');

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `pid` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `path` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `redirect` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `component` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `meta` json NULL,
  `sort` bigint(20) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `path`(`path`) USING BTREE,
  UNIQUE INDEX `path_2`(`path`) USING BTREE,
  INDEX `idx_sys_menu_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1, NULL, '2023-06-22 14:05:48', NULL, 0, 'Setting', '/setting', '/setting/user', 'Layout', '{\"title\": \"系统管理\", \"hidden\": true}', 0);
INSERT INTO `sys_menu` VALUES (2, NULL, NULL, NULL, 1, 'User', 'user', NULL, 'setting/user/index.vue', '{\"title\": \"用户管理\"}', 0);
INSERT INTO `sys_menu` VALUES (3, NULL, NULL, NULL, 1, 'Role', 'role', NULL, 'setting/role/index.vue', '{\"title\": \"角色管理\"}', 0);
INSERT INTO `sys_menu` VALUES (4, NULL, NULL, NULL, 1, 'Menu', 'menu', NULL, 'setting/menu/index.vue', '{\"title\": \"菜单管理\"}', 0);
INSERT INTO `sys_menu` VALUES (5, '2023-03-07 01:50:48', '2023-06-22 14:07:00', NULL, 1, 'Api', 'api', '', 'setting/api/index.vue', '{\"title\": \"API管理\"}', 0);
INSERT INTO `sys_menu` VALUES (7, NULL, NULL, NULL, 6, 'Cenu1', 'cenu1', '/cenu/cenu1/cenu1-1', 'cenu/cenu1/index.vue', '{\"title\": \"cenu1\"}', 0);
INSERT INTO `sys_menu` VALUES (8, NULL, '2023-03-11 14:06:08', NULL, 7, 'Cenu1-1', 'cenu1-1', '', 'cenu/cenu1/cenu1-1/index.vue', '{\"title\": \"cenu1-1\"}', 0);
INSERT INTO `sys_menu` VALUES (9, '2023-03-13 06:14:27', '2023-03-13 06:14:27', NULL, 7, 'Cenu1-2', 'cenu1-2', '', 'cenu/cenu1/cenu1-2/index.vue', '{\"title\": \"cenu1-2\"}', 0);
INSERT INTO `sys_menu` VALUES (10, '2023-06-20 01:21:08', '2023-06-22 15:04:17', NULL, 0, 'customer', '/customer', '/customer/list', 'Layout', '{\"affix\": true, \"title\": \"客户管理\", \"svgIcon\": \"customer-group\"}', 0);
INSERT INTO `sys_menu` VALUES (11, '2023-06-20 01:34:15', '2023-07-05 11:35:54', NULL, 10, 'CustomerList', '/customer/list', '', 'customer/list/index.vue', '{\"title\": \"客户列表\", \"svgIcon\": \"customer-list\"}', 2);
INSERT INTO `sys_menu` VALUES (12, '2023-06-20 01:37:34', '2023-07-07 00:26:48', NULL, 0, 'test', '/test', '/test/index', 'Layout', '{\"title\": \"测试页面\", \"hidden\": true, \"svgIcon\": \"bug\"}', 1);
INSERT INTO `sys_menu` VALUES (13, '2023-06-26 06:04:41', '2023-07-08 10:44:11', NULL, 10, 'CustomerInfo', '/customer/env', '', 'customer/env/index.vue', '{\"title\": \"环境管理\", \"svgIcon\": \"customer-info\"}', 3);
INSERT INTO `sys_menu` VALUES (15, '2023-06-30 05:36:11', '2023-07-07 09:31:04', NULL, 0, 'license', '/license', '/license/list', 'Layout', '{\"affix\": true, \"title\": \"license\", \"svgIcon\": \"license\"}', 1);
INSERT INTO `sys_menu` VALUES (16, '2023-06-30 06:14:56', '2023-07-03 06:46:55', NULL, 15, 'CreateLicense', '/license/generate', '', 'license/createLicense/index.vue', '{\"affix\": true, \"title\": \"创建license\", \"hidden\": true, \"svgIcon\": \"license\"}', 0);
INSERT INTO `sys_menu` VALUES (17, '2023-06-30 07:11:01', '2023-07-07 09:31:24', NULL, 15, 'licenseList', '/license/list', '', 'license/list/index.vue', '{\"affix\": true, \"title\": \"license管理\", \"svgIcon\": \"customer-list\"}', 1);
INSERT INTO `sys_menu` VALUES (18, '2023-07-05 09:19:41', '2023-07-05 09:19:41', NULL, 12, 'Test', '/test/index', '', 'test/index.vue', '{\"title\": \"测试\", \"svgIcon\": \"bug\"}', 0);
INSERT INTO `sys_menu` VALUES (20, '2023-07-06 02:43:31', '2023-07-06 11:50:08', NULL, 10, 'Event', '/customer/event', '', 'customer/event/index.vue', '{\"title\": \"事件管理\", \"svgIcon\": \"customer-list\"}', 5);

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `role_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `role_name`(`role_name`) USING BTREE,
  UNIQUE INDEX `role_name_2`(`role_name`) USING BTREE,
  INDEX `idx_sys_role_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, NULL, '2023-07-07 00:25:59', NULL, 'admin');
INSERT INTO `sys_role` VALUES (2, '2023-06-22 01:51:03', '2023-06-22 01:51:12', NULL, 'devops');

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `username` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '用户名',
  `password` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '密码',
  `phone` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '手机号',
  `email` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '邮箱',
  `active` tinyint(1) NULL DEFAULT NULL,
  `role_model_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE,
  UNIQUE INDEX `username_2`(`username`) USING BTREE,
  INDEX `idx_sys_user_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_sys_user_username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, '2023-02-20 12:51:58', '2023-07-11 16:49:05', NULL, 'admin', 'e10adc3949ba59abbe56e057f20f883e', '11111111111', 'xxxx@xxxxx', 1, 1);

SET FOREIGN_KEY_CHECKS = 1;
