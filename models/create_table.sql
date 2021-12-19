CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `username` VARCHAR(64) COLLATE utf8mb4_general_ci NOT NULL,
  `password` VARCHAR(64) COLLATE utf8mb4_general_ci NOT NULL,
  `email` VARCHAR(64) COLLATE utf8mb4_general_ci,
  `gender` TINYINT(4) NOT NULL DEFAULT '0',
  `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`) USING BTREE,
  UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;
DROP TABLE IF EXISTS `community`;
CREATE TABLE `community` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `community_id` int(10) unsigned NOT NULL,
  `community_name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
  `introduction` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_community_id` (`community_id`),
  UNIQUE KEY `idx_community_name` (`community_name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;