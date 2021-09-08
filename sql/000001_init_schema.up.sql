CREATE TABLE IF NOT EXISTS `users` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NULL,
  `email` VARCHAR(255) NULL,
  `address` VARCHAR(255) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `brands` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `products` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `brand_id` INT UNSIGNED NOT NULL,
  `sku` VARCHAR(128) NULL,
  `name` VARCHAR(255) NULL,
  `qty` INT UNSIGNED NULL,
  `price` BIGINT UNSIGNED NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_products_brands_idx` (`brand_id` ASC),
  CONSTRAINT `fk_products_brands`
    FOREIGN KEY (`brand_id`)
    REFERENCES `brands` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `transactions` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` INT UNSIGNED NOT NULL,
  `date` DATETIME NULL,
  `grand_total` BIGINT UNSIGNED NULL,
  `discount` INT UNSIGNED NULL,
  `reason` TEXT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_transaction_users1_idx` (`user_id` ASC),
  CONSTRAINT `fk_transaction_users1`
    FOREIGN KEY (`user_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `transaction_detail` (
  `transaction_id` INT UNSIGNED NOT NULL,
  `product_id` INT UNSIGNED NOT NULL,
  `price` BIGINT UNSIGNED NULL,
  `qty` INT UNSIGNED NULL,
  `sub_total` BIGINT UNSIGNED NULL,
  INDEX `fk_transaction_detail_transaction1_idx` (`transaction_id` ASC),
  INDEX `fk_transaction_detail_products1_idx` (`product_id` ASC),
  CONSTRAINT `fk_transaction_detail_transaction1`
    FOREIGN KEY (`transaction_id`)
    REFERENCES `transactions` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_transaction_detail_products1`
    FOREIGN KEY (`product_id`)
    REFERENCES `products` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
