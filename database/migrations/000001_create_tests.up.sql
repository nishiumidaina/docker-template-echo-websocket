CREATE TABLE `tests` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `tests` VARCHAR(191) NULL,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;