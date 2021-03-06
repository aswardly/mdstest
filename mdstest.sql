-- MySQL Script generated by MySQL Workbench
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema mdstest
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema mdstest
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `mdstest` DEFAULT CHARACTER SET utf8 ;
USE `mdstest` ;

-- -----------------------------------------------------
-- Table `mdstest`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mdstest`.`users` (
  `user_id` VARCHAR(255) NOT NULL,
  `user_name` VARCHAR(255) NULL,
  `user_password` VARCHAR(255) NULL,
  `user_status` CHAR(3) NULL,
  `last_updated` DATETIME NULL,
  PRIMARY KEY (`user_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mdstest`.`user_settings`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mdstest`.`user_settings` (
  `setting_id` INT NOT NULL AUTO_INCREMENT,
  `user_id` VARCHAR(255) NULL,
  `setting_key` VARCHAR(255) NULL,
  `setting_value` VARCHAR(255) NULL,
  `last_updated` DATETIME NULL,
  PRIMARY KEY (`setting_id`),
  INDEX `fk_user_setting_idx` (`user_id` ASC),
  CONSTRAINT `fk_user_setting`
    FOREIGN KEY (`user_id`)
    REFERENCES `mdstest`.`users` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
