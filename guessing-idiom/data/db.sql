--- use idiom;

CREATE TABLE `stages` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `level` int(11) DEFAULT NULL COMMENT '关卡',
  `contents` TEXT DEFAULT NULL COMMENT '内容json',
  `status` tinyint(1) DEFAULT '1' COMMENT '1 正常，2 下线',
  `created_at` int(11) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ;