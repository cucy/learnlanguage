CREATE DATABASE /*!32312 IF NOT EXISTS*/ `db_apiserver` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `db_apiserver`;
DROP TABLE IF EXISTS `tb_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_users` (
    `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `createdAt` TIMESTAMP NULL DEFAULT NULL,
    `updatedAt` TIMESTAMP NULL DEFAULT NULL,
    `deletedAt` TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`),
    KEY `idx_tb_users_deletedAt` (`deletedAt`)
)   AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8;
LOCK TABLES `tb_users` WRITE;
INSERT INTO `tb_users` VALUES (0,'admin','$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG','2018-05-27 16:25:33','2018-05-27 16:25:33',NULL);

UNLOCK TABLES;

