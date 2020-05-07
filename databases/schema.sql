-- SCHEMA
CREATE TABLE roles
(
    id tinyint PRIMARY KEY,
    name varchar(20) NOT NULL,
    created_at datetime NOT NULL,
    updated_at datetime DEFAULT NULL
);

CREATE TABLE users 
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    email varchar(50) NOT NULL,
    password varchar(100) NOT NULL,
    status char(1) NOT NULL DEFAULT 'I' COMMENT '- (A)ctive\n- (I)nactive\n- (B)anned',
    role_id tinyint(1) NOT NULL DEFAULT 2,
    created_at datetime NOT NULL,
    updated_at datetime DEFAULT NULL,
    CONSTRAINT `role_id` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
);

-- INITIAL DATA
INSERT INTO `roles` VALUES (1,'admin','2020-05-07 18:00:00',NULL);
INSERT INTO `users` VALUES (1,'markus.liang@gmail.com','m123','A',1,'2020-05-07 18:00:00',NULL);
