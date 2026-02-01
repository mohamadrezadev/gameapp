-- +migrate Up
ALTER TABLE `users` ADD COLUMN `role` Enum('user','admin') NOT NULL ;

-- +migrate Down
ALTER TABLE `users` DROP COLUMN `role`;