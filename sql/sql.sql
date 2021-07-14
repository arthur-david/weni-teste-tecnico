CREATE DATABASE IF NOT EXISTS to_do_list;
USE to_do_list;

DROP TABLE IF EXISTS abertos;
DROP TABLE IF EXISTS fechados;

CREATE TABLE abertos(
    id int auto_increment primary key,
    title varchar(255) not null,
    checked boolean not null
) ENGINE=INNODB;

CREATE TABLE fechados(
    id int auto_increment primary key,
    title varchar(255) not null,
    checked boolean not null
) ENGINE=INNODB;