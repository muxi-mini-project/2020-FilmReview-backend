DROP DATABASE IF EXISTS `miniproject`;

CREATE DATABASE `miniproject`;

USE `miniproject`  ;

CREATE TABLE `user` (
	`user_id`   char(6) NOT NULL ,
	`user_picture` varchar(50)  ,
	`name`     varchar(30)  NOT NULL,
	`password`  varchar(16)  NOT NULL,

	PRIMARY KEY (user_id)

)   ENGINE=InnoDB     DEFAULT CHARSET=UTF8;

CREATE TABLE `follow` (
	`user_id1`  char(6)   NOT NULL,
	`user_id2`  char(6)	  NOT NULL,

	KEY (user_id1),
	KEY	(user_id2)
)  	ENGINE=InnoDB   DEFAULT CHARSET=UTF8;

CREATE TABLE `album` (
	`user_id`	char(6)	NOT NULL,
	`album_id`	char(3) NOT NULL,
	`title`		char(30) NOT NULL,
	`summary` TINYTEXT NOT NULL,

	KEY (user_id)

)  	ENGINE=InnoDB   DEFAULT CHARSET=UTF8;

CREATE TABLE `album_id` (
	`user_id`	char(6) NOT NULL,
	`album_id`	char(3) NOT NULL,
	`review_id`	char(10) NOT NULL,

	KEY (user_id)
)  	ENGINE=InnoDB	  DEFAULT CHARSET=UTF8;

CREATE TABLE `collection` (
	`user_id`	char(6)	NOT NULL,
	`review_id`	char(10)	NOT NULL,

	KEY (user_id)
)  	ENGINE=InnoDB   DEFAULT CHARSET=UTF8;


CREATE TABLE `myreviews` (
	`user_id`	char(6)		NOT NULL,
	`review_id`	char(10)	NOT NULL
)  	ENGINE=InnoDB   DEFAULT CHARSET=UTF8;

CREATE TABLE `review_like` (
	`review_id`	char(10)	NOT NULL,
	`user_id`	char(6)		NOT NULL,

	KEY (review_id),
	KEY	(user_id)
)  	ENGINE=InnoDB   DEFAULT CHARSET=UTF8;

CREATE TABLE `comment_like` (
	`user_id`	char(6)	NOT NULL,
	`comment_id`	char(10) NOT NULL,

	KEY (comment_id),
	KEY (user_id)
)  	ENGINE=InnoDB   DEFAULT CHARSET=UTF8;

CREATE TABLE `comment` (
	`comment_id`	char(10)	NOT NULL,
	`content`	TINYTEXT    NOT NULL,
	`time`		TIMESTAMP	NOT NULL,

	KEY	(comment_id)
)  	ENGINE=InnoDB   DEFAULT CHARSET=UTF8;

CREATE TABLE `user_com` (
	`user_id`	char(6)	NOT NULL,
	`review_id`	char(10)	NOT NULL,
	`comment_id` char(10)	NOT NULL,

	KEY	(review_id)
)  	ENGINE=InnoDB   DEFAULT CHARSET=UTF8;

CREATE TABLE `review` (
	`review_id`	char(10)	PRIMARY KEY NOT NULL,
	`title`		varchar(50)	NOT NULL,
	`content`	TEXT NOT NULL,
	`time`		TIMESTAMP NOT NULL,
	`tag`	varchar(60)	NOT NULL,
	`picture`	varchar(50)	NOT NULL,
	`like_sum`	varchar(8)	NOT NULL,

	KEY (`review_id`)
)  	ENGINE=InnoDB   DEFAULT CHARSET=UTF8;




