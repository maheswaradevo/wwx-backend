CREATE TABLE `projects` (
  `id` int NOT NULL AUTO_INCREMENT,
  `project_name` varchar(40) NOT NULL,
  `client_name` varchar(30) NOT NULL,
  `deadline` timestamp NOT NULL,
  `status` varchar(20) NOT NULL,
  `budget` int DEFAULT NULL,
  `proposal_link` varchar(255) DEFAULT NULL,
  `assign` varchar(20) DEFAULT NULL,
  `user_id` int DEFAULT NULL,
  `resource_link` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `maintenance` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `projects_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
);