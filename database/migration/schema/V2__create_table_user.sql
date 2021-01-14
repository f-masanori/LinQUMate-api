CREATE TABLE `user` (
  `id` int UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  `firebase_uid` varchar(255) NOT NULL,
  `email` varchar(255),
  `name` varchar(255),
  `icon_url` varchar(255),
  `faculty` varchar(255),
  `grade` varchar(255),
  `department` varchar(255),
  `ctime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `utime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY `firebase_uid`
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO user 
(id, firebase_uid, email, name, icon_url,faculty,grade,department) VALUES 
(1, "0123456789", "test1@test.com", "user1", "url 1", "fa 1","","");
