-- tweetsテーブルの作成
CREATE TABLE `tweets`
(
	`id` bigint(20) NOT NULL AUTO_INCREMENT,
	`text` varchar(255) NOT NULL,
	`created_at` datetime NOT NULL,
	`updated_at` datetime NOT NULL,
	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;