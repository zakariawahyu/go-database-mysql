-- MySQL dump 10.13  Distrib 8.0.27, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: golang_database
-- ------------------------------------------------------
-- Server version	5.7.35

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `comments`
--

DROP TABLE IF EXISTS `comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comments` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(100) NOT NULL,
  `comment` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comments`
--

LOCK TABLES `comments` WRITE;
/*!40000 ALTER TABLE `comments` DISABLE KEYS */;
INSERT INTO `comments` VALUES (1,'zakaria@gmail.com','test comment'),(2,'zakaria@gmail.com','test comment'),(3,'zakaria@gmail.com','test comment'),(4,'zakaria@gmail.com','test comment'),(5,'zakaria0@gmail.com','Ini comment ke 0'),(6,'zakaria1@gmail.com','Ini comment ke 1'),(7,'zakaria2@gmail.com','Ini comment ke 2'),(8,'zakaria3@gmail.com','Ini comment ke 3'),(9,'zakaria4@gmail.com','Ini comment ke 4'),(10,'zakaria5@gmail.com','Ini comment ke 5'),(11,'zakaria6@gmail.com','Ini comment ke 6'),(12,'zakaria7@gmail.com','Ini comment ke 7'),(13,'zakaria8@gmail.com','Ini comment ke 8'),(14,'zakaria9@gmail.com','Ini comment ke 9'),(15,'zakaria0@gmail.com','Ini comment ke 0'),(16,'zakaria1@gmail.com','Ini comment ke 1'),(17,'zakaria2@gmail.com','Ini comment ke 2'),(18,'zakaria3@gmail.com','Ini comment ke 3'),(19,'zakaria4@gmail.com','Ini comment ke 4'),(20,'zakaria5@gmail.com','Ini comment ke 5'),(21,'zakaria6@gmail.com','Ini comment ke 6'),(22,'zakaria7@gmail.com','Ini comment ke 7'),(23,'zakaria8@gmail.com','Ini comment ke 8'),(24,'zakaria9@gmail.com','Ini comment ke 9'),(25,'repository@test.com','Test repository'),(26,'repository@test.com','Test repository'),(27,'repository@test.com','Test repository'),(28,'repository@test.com','Test repository');
/*!40000 ALTER TABLE `comments` ENABLE KEYS */;
UNLOCK TABLES;
