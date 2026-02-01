
-- +migrate Up
Create Table `permissions` (
    `id` INT PRIMARY  Key AUTO_INCREMENT,
    `title` varchar(191) NOT NULL UNIQUE,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);