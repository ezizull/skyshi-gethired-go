CREATE TABLE IF NOT EXISTS `activities` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`title` longtext,
	`email` longtext,
	`created_at` datetime(3) DEFAULT NULL,
	`updated_at` datetime(3) DEFAULT NULL,
	`deleted_at` timestamp NULL DEFAULT NULL,
	PRIMARY KEY (`id`)
  ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;