/*
 Navicat Premium Data Transfer

 Source Server         : 姜姜大人
 Source Server Type    : MySQL
 Source Server Version : 80017
 Source Host           : localhost:3306
 Source Schema         : homework

 Target Server Type    : MySQL
 Target Server Version : 80017
 File Encoding         : 65001

 Date: 08/04/2020 21:24:40
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for announcement
-- ----------------------------
DROP TABLE IF EXISTS `announcement`;
CREATE TABLE `announcement` (
  `announcementID` int(11) NOT NULL AUTO_INCREMENT COMMENT '公告ID',
  `content` text COMMENT '公告内容',
  `name` varchar(255) DEFAULT NULL COMMENT '公告名称',
  `time` datetime DEFAULT NULL COMMENT '发布时间',
  `releaseName` int(11) DEFAULT NULL COMMENT '发布人ID',
  PRIMARY KEY (`announcementID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for class
-- ----------------------------
DROP TABLE IF EXISTS `class`;
CREATE TABLE `class` (
  `className` varchar(255) DEFAULT NULL COMMENT '班级名称',
  `classID` int(11) NOT NULL AUTO_INCREMENT COMMENT '班级ID',
  `createID` int(11) NOT NULL COMMENT '创建者ID',
  `invitation` int(255) NOT NULL COMMENT '邀请码',
  `number` int(11) DEFAULT NULL COMMENT '人数',
  `startTime` datetime NOT NULL COMMENT '创建时间',
  `membersID` varchar(255) DEFAULT NULL COMMENT '成员ID',
  PRIMARY KEY (`classID`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of class
-- ----------------------------
BEGIN;
INSERT INTO `class` VALUES ('dsada', 9, 8, 8603, 0, '2020-03-25 08:32:17', '');
INSERT INTO `class` VALUES ('222', 10, 8, 4604, 0, '2020-03-25 08:35:09', '');
INSERT INTO `class` VALUES ('222', 11, 8, 4421, 0, '2020-03-25 08:35:34', '');
INSERT INTO `class` VALUES ('22', 12, 8, 6092, 0, '2020-03-25 08:40:43', '');
INSERT INTO `class` VALUES ('421', 13, 8, 117, 0, '2020-03-25 08:43:11', '');
INSERT INTO `class` VALUES ('1111', 14, 8, 721, 0, '2020-03-25 08:43:15', '');
INSERT INTO `class` VALUES ('412421', 15, 8, 8433, 0, '2020-03-25 08:43:18', '');
INSERT INTO `class` VALUES ('123', 19, 8, 9652, 0, '2020-03-26 06:01:31', '');
INSERT INTO `class` VALUES ('123', 20, 8, 7626, 0, '2020-03-26 06:01:35', '');
INSERT INTO `class` VALUES ('123', 21, 8, 4687, 0, '2020-03-26 11:35:17', '');
INSERT INTO `class` VALUES ('fff', 22, 8, 8719, 0, '2020-03-26 12:28:22', '');
INSERT INTO `class` VALUES ('zzz', 23, 8, 6271, 0, '2020-03-26 12:31:46', '');
INSERT INTO `class` VALUES ('hhh', 24, 8, 2726, 0, '2020-03-26 12:49:18', '');
INSERT INTO `class` VALUES ('15173135646', 27, 8, 4723, 0, '2020-03-26 13:03:20', '');
INSERT INTO `class` VALUES ('15173135646', 28, 8, 874, 0, '2020-03-26 13:07:40', '');
INSERT INTO `class` VALUES ('333', 29, 8, 1578, 1, '2020-03-26 13:08:09', '7');
INSERT INTO `class` VALUES ('qqq', 30, 8, 5193, 1, '2020-03-26 13:17:59', '7');
INSERT INTO `class` VALUES ('aaa', 31, 8, 9841, 2, '2020-03-26 21:22:23', '7,11');
COMMIT;

-- ----------------------------
-- Table structure for homework
-- ----------------------------
DROP TABLE IF EXISTS `homework`;
CREATE TABLE `homework` (
  `StartTime` datetime NOT NULL COMMENT '开始时间',
  `EndTime` datetime NOT NULL COMMENT '结束时间',
  `HomeworkName` varchar(255) NOT NULL COMMENT '作业名称',
  `HomeworkID` int(11) NOT NULL COMMENT '作业号',
  `classID` int(11) DEFAULT NULL COMMENT '班级ID',
  PRIMARY KEY (`HomeworkID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for title_bank
-- ----------------------------
DROP TABLE IF EXISTS `title_bank`;
CREATE TABLE `title_bank` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '题ID',
  `topic_describe` text COMMENT '描述',
  `category` int(10) DEFAULT NULL COMMENT '题目类型',
  `answer` varchar(255) DEFAULT NULL COMMENT '题目答案',
  `annotation` varchar(255) DEFAULT NULL COMMENT '题目注释',
  `knowledge` int(20) DEFAULT NULL COMMENT '知识点',
  `createID` int(11) DEFAULT NULL COMMENT '创建者ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of title_bank
-- ----------------------------
BEGIN;
INSERT INTO `title_bank` VALUES (1, '{\"options\":[\"<p>选项1答案</p>\",\"<p>选项2答案</p>\"],\"multiSelect\":false}', 0, '<p>答案：1</p>', '<p>这里是解释</p>', 1001, 8);
INSERT INTO `title_bank` VALUES (2, '{\"problem\":\"<p>问题2</p>\",\"options\":[\"<p>选项1</p>\",\"<p>选项2<span class=\\\"mathquill-rendered-math\\\" style=\\\"font-size:20px;\\\"><span class=\\\"textarea\\\"><textarea data-cke-editable=\\\"1\\\" contenteditable=\\\"false\\\"></textarea></span><big mathquill-command-id=\\\"4\\\">∑</big><span mathquill-command-id=\\\"6\\\">2</span><span mathquill-command-id=\\\"7\\\" class=\\\"\\\">log</span><sub class=\\\"non-leaf\\\" mathquill-command-id=\\\"9\\\" mathquill-block-id=\\\"10\\\"><span mathquill-command-id=\\\"12\\\">1</span><span mathquill-command-id=\\\"15\\\" class=\\\"non-italicized-function\\\">ln</span><span mathquill-command-id=\\\"13\\\">1</span><span mathquill-command-id=\\\"14\\\">1</span><span mathquill-command-id=\\\"17\\\">2</span></sub></span><span>&nbsp;这是什么意思</span></p>\"],\"multiSelect\":false}', 0, '<p>选项1</p>', '<p>没有解释</p>', 1001, 8);
COMMIT;

-- ----------------------------
-- Table structure for tree
-- ----------------------------
DROP TABLE IF EXISTS `tree`;
CREATE TABLE `tree` (
  `itselfID` int(11) DEFAULT NULL COMMENT '本身ID',
  `describe` varchar(225) DEFAULT NULL COMMENT '本身描述',
  `fatherID` int(20) DEFAULT NULL COMMENT '父ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `phonenumber` bigint(20) NOT NULL COMMENT '电话号码',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `id` int(20) NOT NULL COMMENT 'ID',
  `position` tinyint(255) NOT NULL COMMENT '身份',
  `userName` char(255) DEFAULT NULL COMMENT '用户姓名',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (15173135646, '123', 8, 1, '张二');
INSERT INTO `user` VALUES (15898538046, '123', 9, 0, '张三');
INSERT INTO `user` VALUES (15873295831, '123', 10, 0, '张三');
INSERT INTO `user` VALUES (15873295833, '123', 11, 0, '姜一');
INSERT INTO `user` VALUES (15873295834, '123', 12, 0, '姜二');
INSERT INTO `user` VALUES (15873295835, '123', 13, 0, '姜三');
INSERT INTO `user` VALUES (15173135641, '123', 14, 1, '姜四');
INSERT INTO `user` VALUES (15173135642, '123', 15, 1, '张四');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
