-- ************************************** `create database filer-loader`
CREATE DATABASE "filer-loader"
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1;

COMMENT ON DATABASE "filer-loader"
    IS 'filer-loader database';


-- ************************************** `tbl_status`

CREATE TABLE tbl_status 
(
 id integer PRIMARY KEY,
 status varchar(50),
);


-- ************************************** `tbl_group`

CREATE TABLE tbl_group
(
 id integer PRIMARY KEY,
 name varchar(45) NOT NULL ,
 imdb varchar(100) NOT NULL ,
 type varchar(50) NOT NULL ,
);


-- ************************************** `tbl_tasks`

CREATE TABLE tbl_tasks
(
 id integer PRIMARY KEY ,
 fk_group int NOT NULL ,
 link     varchar(256) NOT NULL ,
 hash     varchar(256) NOT NULL ,
 name     varchar(100) NOT NULL ,
 fk_status   int NOT NULL ,
 message  varchar(500) NOT NULL ,
 size     int NOT NULL ,
 CONSTRAINT fk_status
     FOREIGN KEY (fk_status) 
     REFERENCES tbl_status (id) ,
 CONSTRAINT fk_group
     FOREIGN KEY (fk_group) 
     REFERENCES tbl_group (id)
);

