/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80023
 Source Host           : localhost:3306
 Source Schema         : quick_chat

 Target Server Type    : MySQL
 Target Server Version : 80023
 File Encoding         : 65001

 Date: 18/09/2022 15:51:54
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for chat_record
-- ----------------------------
DROP TABLE IF EXISTS `chat_record`;
CREATE TABLE `chat_record` (
  `id` varchar(255) NOT NULL COMMENT '聊天记录ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `user_relation_id` varchar(36) NOT NULL COMMENT '关系id',
  `record_type` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '1文本类型 2语音',
  `record` varchar(255) DEFAULT NULL COMMENT '聊天内容',
  `extend` varchar(100) NOT NULL DEFAULT '' COMMENT '扩展字段',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='聊天记录表';

-- ----------------------------
-- Table structure for community
-- ----------------------------
DROP TABLE IF EXISTS `community`;
CREATE TABLE `community` (
  `id` varchar(255) NOT NULL COMMENT '群聊ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `name` varchar(30) DEFAULT NULL COMMENT '群名称',
  `owner_id` bigint DEFAULT NULL COMMENT '群主ID',
  `icon` varchar(250) DEFAULT NULL COMMENT '群logo',
  `memo` varchar(120) DEFAULT NULL COMMENT '群描述',
  `extend` varchar(100) NOT NULL DEFAULT '' COMMENT '扩展字段',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='群聊表';

-- ----------------------------
-- Table structure for user_base
-- ----------------------------
DROP TABLE IF EXISTS `user_base`;
CREATE TABLE `user_base` (
  `id` varchar(255) NOT NULL COMMENT '用户ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `nick_name` varchar(32) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `user_id` varchar(32) NOT NULL DEFAULT '' COMMENT '用户id，必须唯一',
  `password` varchar(100) NOT NULL DEFAULT '' COMMENT '用户密码',
  `user_role` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '1正常用户 2封禁用户 3管理员',
  `gender` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '用户性别 1男 2女',
  `signature` varchar(255) NOT NULL DEFAULT '' COMMENT '用户个人签名',
  `mobile` varchar(16) NOT NULL DEFAULT '' COMMENT '手机号码(唯一)',
  `face` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `extend1` varchar(100) NOT NULL DEFAULT '' COMMENT '扩展字段1',
  `extend2` varchar(100) NOT NULL DEFAULT '' COMMENT '扩展字段2',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户基础信息表';

-- ----------------------------
-- Table structure for user_relation
-- ----------------------------
DROP TABLE IF EXISTS `user_relation`;
CREATE TABLE `user_relation` (
  `id` varchar(255) NOT NULL COMMENT '关系ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `user_id` varchar(32) NOT NULL COMMENT '用户id',
  `friend_id` varchar(32) NOT NULL COMMENT '好友id',
  `relation_type` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '1一对一好友 2群聊好友',
  `memo` varchar(120) DEFAULT NULL COMMENT '描述',
  `extend` varchar(100) NOT NULL DEFAULT '' COMMENT '扩展字段',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户关系表';

SET FOREIGN_KEY_CHECKS = 1;
