CREATE DATABASE IF NOT EXISTS verifymyage;

use verifymyage;

CREATE TABLE IF NOT EXISTS users(
    id       varchar(40) NOT NULL PRIMARY KEY,
    name     VARCHAR(120) NOT NULL ,
    age      INTEGER NOT NULL ,
    email    VARCHAR(255) DEFAULT NULL,
    password VARCHAR(255) DEFAULT NULL,
    address  VARCHAR(100) DEFAULT NULL,
    status   VARCHAR(7)   DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=INNODB;