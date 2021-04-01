# Host: localhost  (Version: 5.7.26)
# Date: 2021-04-02 06:24:16
# Generator: MySQL-Front 5.3  (Build 4.234)

/*!40101 SET NAMES utf8 */;

#
# Structure for table "gc_friends"
#

DROP TABLE IF EXISTS `gc_friends`;
CREATE TABLE `gc_friends` (
  `uid` int(11) NOT NULL AUTO_INCREMENT,
  `fid` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#
# Data for table "gc_friends"
#

INSERT INTO `gc_friends` VALUES (1,2),(2,1);

#
# Structure for table "gc_users"
#

DROP TABLE IF EXISTS `gc_users`;
CREATE TABLE `gc_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(255) DEFAULT NULL,
  `nickname` varchar(64) DEFAULT NULL,
  `password` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

#
# Data for table "gc_users"
#

/*!40000 ALTER TABLE `gc_users` DISABLE KEYS */;
INSERT INTO `gc_users` VALUES (1,'1711861430@qq.com','奋斗者之书',NULL),(2,'2256684847@qq.com','霜冷血饮寒',NULL);
/*!40000 ALTER TABLE `gc_users` ENABLE KEYS */;
