/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : 127.0.0.1:3306
 Source Schema         : oh-admin

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 03/04/2023 16:55:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for oh_admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `oh_admin_menu`;
CREATE TABLE `oh_admin_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `icon` varchar(50) DEFAULT NULL COMMENT '图标',
  `url` varchar(100) DEFAULT NULL COMMENT 'URL地址',
  `pid` int(11) NOT NULL DEFAULT '0' COMMENT '上级ID',
  `type` tinyint(2) NOT NULL DEFAULT '0' COMMENT '类型：1模块 2导航 3菜单 4节点',
  `permission` varchar(255) DEFAULT '' COMMENT '权限标识',
  `is_show` tinyint(1) DEFAULT '1' COMMENT '是否显示：1显示 2不显示',
  `sort` int(11) DEFAULT NULL COMMENT '显示顺序',
  `remark` varchar(255) DEFAULT NULL COMMENT '菜单备注',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态1-在用 2-禁用',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of oh_admin_menu
-- ----------------------------
BEGIN;
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (1, '系统管理', 'layui-icon-component', '#', 0, 0, 'sys:sysconfig', 1, 1, '', 1, '2023-03-15 07:20:52', '2023-03-30 14:58:07');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (2, '权限管理', 'layui-icon-component', NULL, 1, 0, '', 1, 2, NULL, 1, '2023-03-15 07:21:28', '2023-03-24 07:19:18');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (3, '用户管理', 'layui-icon-component', 'user/index', 2, 0, '', 1, 3, NULL, 1, '2023-03-15 07:23:02', '2023-03-24 07:19:19');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (4, '角色管理', 'layui-icon-component', 'role/index', 2, 0, '', 1, 3, NULL, 1, '2023-03-15 07:23:02', '2023-03-27 02:26:41');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (5, '菜单管理', 'layui-icon-component', 'menu/index', 2, 0, '', 1, 3, NULL, 1, '2023-03-15 07:23:02', '2023-03-27 03:28:15');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (6, '运营管理', '', '', 0, 0, '', 1, 1, '', 1, '2023-03-30 15:00:28', '2023-04-03 08:54:27');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (7, '添加用户', 'layui-icon-rate', 'user/add', 3, 1, 'sys:user:add', 1, 1, ' 添加用户', 1, '2023-03-28 15:29:02', '2023-04-03 08:55:18');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (16, '友链管理', 'layui-icon-link', '#', 6, 0, 'sys:link', 1, 2, '', 1, '2023-03-30 14:59:53', '2023-03-30 14:59:53');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (17, '友链列表', 'layui-icon-util', 'link/index', 16, 0, 'sys:link:index', 1, 1, '', 1, '2023-03-30 15:00:28', '2023-04-03 08:55:21');
COMMIT;

-- ----------------------------
-- Table structure for oh_admin_role
-- ----------------------------
DROP TABLE IF EXISTS `oh_admin_role`;
CREATE TABLE `oh_admin_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '角色名称',
  `code` varchar(50) NOT NULL DEFAULT '' COMMENT '角色编码',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态 1-启用2-禁用',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of oh_admin_role
-- ----------------------------
BEGIN;
INSERT INTO `oh_admin_role` (`id`, `name`, `code`, `status`, `sort`, `created_time`, `updated_time`) VALUES (1, '超级管理员', 'super', 1, 3, '2023-03-27 02:01:56', '2023-03-30 16:29:13');
INSERT INTO `oh_admin_role` (`id`, `name`, `code`, `status`, `sort`, `created_time`, `updated_time`) VALUES (2, '管理员', 'admin', 1, 0, '2023-03-27 14:46:24', '2023-03-27 18:02:14');
INSERT INTO `oh_admin_role` (`id`, `name`, `code`, `status`, `sort`, `created_time`, `updated_time`) VALUES (3, '测试', 'test', 1, 0, '2023-03-27 14:57:33', '2023-03-30 16:53:50');
COMMIT;

-- ----------------------------
-- Table structure for oh_admin_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `oh_admin_role_menu`;
CREATE TABLE `oh_admin_role_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `menu_id` int(11) NOT NULL DEFAULT '0' COMMENT '菜单id',
  `role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色id',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_role` (`role_id`),
  KEY `idx_menu` (`menu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of oh_admin_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (40, 1, 1, '2023-03-30 15:00:35', '2023-03-30 15:00:35');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (41, 2, 1, '2023-03-30 15:00:35', '2023-03-30 15:00:35');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (42, 3, 1, '2023-03-30 15:00:35', '2023-03-30 15:00:35');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (43, 4, 1, '2023-03-30 15:00:35', '2023-03-30 15:00:35');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (44, 5, 1, '2023-03-30 15:00:35', '2023-03-30 15:00:35');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (45, 11, 1, '2023-03-30 15:00:35', '2023-03-30 15:00:35');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (46, 12, 1, '2023-03-30 15:00:35', '2023-03-30 15:00:35');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (47, 6, 1, '2023-03-30 15:00:35', '2023-03-30 15:00:35');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (48, 13, 1, '2023-03-30 15:00:35', '2023-03-30 15:00:35');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (49, 14, 1, '2023-03-30 15:00:35', '2023-03-30 15:00:35');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (50, 15, 1, '2023-03-30 15:00:35', '2023-03-30 15:00:35');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (51, 16, 1, '2023-03-30 15:00:35', '2023-03-30 15:00:35');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (52, 17, 1, '2023-03-30 15:00:35', '2023-03-30 15:00:35');
COMMIT;

-- ----------------------------
-- Table structure for oh_admin_user
-- ----------------------------
DROP TABLE IF EXISTS `oh_admin_user`;
CREATE TABLE `oh_admin_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `realname` varchar(50) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `username` varchar(100) NOT NULL DEFAULT '' COMMENT '登录用户名',
  `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别:1男 2女 3保密',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `mobile` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `email` varchar(100) NOT NULL DEFAULT '' COMMENT '邮箱地址',
  `address` varchar(255) DEFAULT '' COMMENT '地址',
  `password` varchar(150) NOT NULL DEFAULT '' COMMENT '登录密码',
  `salt` varchar(30) NOT NULL DEFAULT '' COMMENT '盐加密',
  `intro` varchar(255) DEFAULT '' COMMENT '备注',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态：1正常 2禁用',
  `login_num` int(11) NOT NULL DEFAULT '0' COMMENT '登录次数',
  `login_ip` varchar(20) NOT NULL DEFAULT '' COMMENT '最近登录ip',
  `login_time` int(11) NOT NULL DEFAULT '0' COMMENT '最近登录时间',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of oh_admin_user
-- ----------------------------
BEGIN;
INSERT INTO `oh_admin_user` (`id`, `realname`, `username`, `gender`, `avatar`, `mobile`, `email`, `address`, `password`, `salt`, `intro`, `status`, `login_num`, `login_ip`, `login_time`, `created_time`, `updated_time`) VALUES (1, '管理员', 'admin', 2, '', '123456', 'sd_mwq@163.com', '北京市', '52af3ce8a82f62707789fe00899ed3f0', '123456', 'test', 1, 15, '127.0.0.1', 1680511743, '2023-03-16 09:38:04', '2023-04-03 16:49:03');
COMMIT;

-- ----------------------------
-- Table structure for oh_admin_user_role
-- ----------------------------
DROP TABLE IF EXISTS `oh_admin_user_role`;
CREATE TABLE `oh_admin_user_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
  `role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色id',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_role` (`user_id`,`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of oh_admin_user_role
-- ----------------------------
BEGIN;
INSERT INTO `oh_admin_user_role` (`id`, `user_id`, `role_id`, `created_time`, `updated_time`) VALUES (1, 1, 1, '2023-03-27 02:02:14', NULL);
COMMIT;

-- ----------------------------
-- Table structure for oh_link
-- ----------------------------
DROP TABLE IF EXISTS `oh_link`;
CREATE TABLE `oh_link` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '友链名称',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '友链地址',
  `image` varchar(255) DEFAULT '' COMMENT 'logo',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1-启用 2-禁用',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of oh_link
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
