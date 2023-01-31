-- MySQL dump 10.13  Distrib 8.0.26, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: app_database
-- ------------------------------------------------------
-- Server version	8.0.26

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `accounts`
--

DROP TABLE IF EXISTS `accounts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `accounts` (
  `account_id` int NOT NULL AUTO_INCREMENT,
  `password` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `firstname` varchar(255) NOT NULL,
  `lastname` varchar(255) NOT NULL,
  `phone` varchar(10) NOT NULL,
  `role` varchar(255) NOT NULL,
  `status` varchar(45) NOT NULL,
  `reneval` datetime DEFAULT NULL,
  PRIMARY KEY (`account_id`,`email`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `accounts`
--

LOCK TABLES `accounts` WRITE;
/*!40000 ALTER TABLE `accounts` DISABLE KEYS */;
INSERT INTO `accounts` VALUES (31,'$2a$14$a7FAyCGF/iC0H2nU9NUGF.2.lOkPFyZZWN7qRUUwnxz08euamlWOa','daemon@gmail.com','daemon','targelian','0984437173','admin','inactive',NULL),(41,'$2a$14$/okZpJhBQMTDk7Q36rGuNupEtUXLJDrhlU0VZySD3yRHOKSJpnnUm','admin@gmail.com','MasterAdmin','MasterAdmin','0984437173','admin','inactive','2022-12-05 00:00:00');
/*!40000 ALTER TABLE `accounts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `bucketimage`
--

DROP TABLE IF EXISTS `bucketimage`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `bucketimage` (
  `image_id` int NOT NULL AUTO_INCREMENT,
  `product_id` int NOT NULL,
  `url_path` varchar(255) NOT NULL,
  `image_name` varchar(255) NOT NULL,
  `upload_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`image_id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bucketimage`
--

LOCK TABLES `bucketimage` WRITE;
/*!40000 ALTER TABLE `bucketimage` DISABLE KEYS */;
INSERT INTO `bucketimage` VALUES (15,1,'https://storage.googleapis.com/image_services_golang/product1.png','product1.png','2022-01-27 13:38:36'),(16,1,'https://storage.googleapis.com/image_services_golang/product1.png','product1.png','2022-01-27 13:38:36'),(20,1,'https://google.cloud.arnon.png','arnon.png','2022-01-27 13:38:36'),(21,1,'https://storage.googleapis.com/image_services_golang/product2.png','product2.png','2022-01-27 13:38:36'),(22,1,'https://storage.googleapis.com/image_services_golang/product13.png','product13.png','2022-01-27 13:39:02');
/*!40000 ALTER TABLE `bucketimage` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `product_id` int NOT NULL AUTO_INCREMENT,
  `product_name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `product_type` varchar(255) NOT NULL,
  `bgColor` varchar(255) NOT NULL,
  `price` float(10,2) NOT NULL,
  PRIMARY KEY (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (33,'new name','destest','shoes','#fff',55.09),(34,'Nike all Evans','destest','shoes','#fff',55.09),(35,'Nike all Evans','destest','shoes','#fff',55.09),(36,'Nike all Evans','destest','shoes','#fff',55.09),(39,'vans all start','destest','shoes','#fff',55.09);
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `idusers` int NOT NULL AUTO_INCREMENT,
  `account_id` int NOT NULL,
  `usr_idx` int NOT NULL,
  `username` varchar(45) NOT NULL,
  `image_url` varchar(255) NOT NULL,
  `pin` varchar(255) NOT NULL,
  PRIMARY KEY (`idusers`,`account_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,41,1,'arnon','http://localhost:8080/public/user1_default.jpg','$2a$14$zz.0zpb1smzT9ryOkfnS8eWY3YvJizGjmabdMZgda/HK4SHZwizIa'),(2,41,2,'Sunio','http://localhost:8080/public/user2_default.png','$2a$14$zz.0zpb1smzT9ryOkfnS8eWY3YvJizGjmabdMZgda/HK4SHZwizIa'),(3,41,3,'Sunligh','http://localhost:8080/public/user3_default.jpg','$2a$14$zz.0zpb1smzT9ryOkfnS8eWY3YvJizGjmabdMZgda/HK4SHZwizIa'),(4,41,4,'StremAdmin','http://localhost:8080/public/user4_default.jpg','$2a$14$zz.0zpb1smzT9ryOkfnS8eWY3YvJizGjmabdMZgda/HK4SHZwizIa'),(5,41,5,'Wishdom','http://localhost:8080/public/user5_default.png','$2a$14$zz.0zpb1smzT9ryOkfnS8eWY3YvJizGjmabdMZgda/HK4SHZwizIa');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-01-31 22:50:27
