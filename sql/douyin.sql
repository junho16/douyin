/*
Navicat MySQL Data Transfer

Source Server         : MYSQL
Source Server Version : 80027
Source Host           : localhost:3306
Source Database       : douyin

Target Server Type    : MYSQL
Target Server Version : 80027
File Encoding         : 65001

Date: 2022-06-14 23:15:23
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `video_id` int DEFAULT NULL,
  `user_id` int DEFAULT NULL,
  `content` varchar(255) DEFAULT NULL,
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES ('5', '3', '，thh', '3', '2022-06-14 22:49:36', '2022-06-14 22:49:36', null);

-- ----------------------------
-- Table structure for favorites
-- ----------------------------
DROP TABLE IF EXISTS `favorites`;
CREATE TABLE `favorites` (
  `user_id` int DEFAULT NULL,
  `video_id` int DEFAULT NULL,
  `state` int DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of favorites
-- ----------------------------
INSERT INTO `favorites` VALUES ('0', '5', '0', '2022-06-14 22:24:10', '2022-06-14 22:24:10', null, '1');
INSERT INTO `favorites` VALUES ('3', '5', '0', '2022-06-14 22:54:04', '2022-06-14 22:54:04', null, '2');

-- ----------------------------
-- Table structure for followers
-- ----------------------------
DROP TABLE IF EXISTS `followers`;
CREATE TABLE `followers` (
  `host_id` int DEFAULT NULL,
  `guest_id` int DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`),
  KEY `idx_followers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of followers
-- ----------------------------
INSERT INTO `followers` VALUES ('3', '3', '2022-06-14 22:50:27', '2022-06-14 22:50:27', '2022-06-14 22:54:17', '1');

-- ----------------------------
-- Table structure for followings
-- ----------------------------
DROP TABLE IF EXISTS `followings`;
CREATE TABLE `followings` (
  `host_id` int DEFAULT NULL,
  `guest_id` int DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of followings
-- ----------------------------
INSERT INTO `followings` VALUES ('3', '3', '2022-06-14 22:50:27', '2022-06-14 22:50:27', '2022-06-14 22:54:17', '1');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `follow_count` int unsigned DEFAULT NULL,
  `follower_count` int unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('1', '2022-06-09 23:16:03', '2022-06-09 23:16:03', null, 'wjh', '3132333435363738f156cc10b7ac90757f82c743f38cc912', '0', '0');
INSERT INTO `users` VALUES ('2', '2022-06-09 23:17:25', '2022-06-09 23:17:25', null, 'wjh2', '31323334353637385511818b35f42142eca0ce8d66206647', '0', '0');
INSERT INTO `users` VALUES ('3', '2022-06-12 15:40:05', '2022-06-14 22:54:18', null, 'nchu', '31323334353637385511818b35f42142eca0ce8d66206647', '0', '0');

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `author_id` int DEFAULT NULL,
  `play_url` varchar(255) DEFAULT NULL,
  `cover_url` varchar(255) DEFAULT NULL,
  `favorite_count` int DEFAULT NULL,
  `comment_count` int DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of videos
-- ----------------------------
INSERT INTO `videos` VALUES ('1', '1', 'http://10.68.2.116:8080/data/videos/0_wx_camera_1654046478743.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', null, null, 'tsr', '2022-06-12 17:07:35', '2022-06-12 17:07:35', null);
INSERT INTO `videos` VALUES ('2', '1', 'http://10.68.2.116:8080/data/videos/0_wx_camera_1647671018956.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', null, null, 'test2', '2022-06-12 19:54:37', '2022-06-12 19:54:37', null);
INSERT INTO `videos` VALUES ('4', '3', 'http://10.68.2.116:8080/data/videos/3_wx_camera_1647945336334.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', null, null, 'tsq', '2022-06-12 23:14:30', '2022-06-12 23:14:30', null);
INSERT INTO `videos` VALUES ('5', '3', 'http://10.68.2.116:8080/data/videos/3_wx_camera_1646884539470.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', null, null, '停停停', '2022-06-12 23:56:51', '2022-06-12 23:56:51', null);
