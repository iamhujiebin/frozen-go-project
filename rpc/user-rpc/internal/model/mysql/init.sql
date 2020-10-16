CREATE TABLE `users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `access_token` varchar(256) NOT NULL COMMENT '用户校验token',
  `avatar` varchar(45) DEFAULT NULL COMMENT '用户头像',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1 COMMENT='用户表'

CREATE TABLE `Frozen`.`guests` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `guest_id` varchar(256) NOT NULL COMMENT '游客id',
  `platform` varchar(45) NOT NULL COMMENT '平台：ios/android/web',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='游客表'