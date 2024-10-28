CREATE TABLE `users` (
                         `id` varchar(24) COLLATE utf8mb4_unicode_ci NOT NULL,
                         `avatar` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
                         `nickname` varchar(24) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
                         `phone` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
                         `password` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
                         `status` tinyint(4) NOT NULL DEFAULT '0',
                         `sex` tinyint(4) NOT NULL DEFAULT '0',
                         `created_at` timestamp NOT NULL,
                         `updated_at` timestamp NOT NULL,
                         `deleted_at` timestamp NULL DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci