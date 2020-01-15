create table user (
    user_id char(6) PRIMARY KEY not null,
    user_picture varchar(30) not null,
    name varchar(30) not null,
    password varchar(16) not null
);

create table follow (
    user_id1 char(6) not null,
    user_id2 char(6) not null,
    INDEX id_index (user_id1),
    INDEX id_index2 (user_id2)
);

create table album (
    user_id char(6) not null,
    album_id char(3) not null,
    title varchar(30) not null,
    summary TINYTEXT not null,
    content_sum varchar(3) not null,
    UNIQUE id (user_id,album_id)
);

create table album_review (
    user_id char(6) not null,
    album_id char(3) not null,
    review_id char(10) not null,
    UNIQUE id (user_id,album_id)
);

create table collection (
    user_id char(6) not null,
    review_id char(10) not null,
    INDEX id_index (user_id)
);

create table myreviews (
    user_id char(6) not null,
    review_id char(10) not null,
    INDEX id_index (user_id)
);

create table review_like (
    review_id char(10) not null,
    user_id char(6) not null,
    INDEX id_index (review_id)
);

create table comment_like (
    user_id char(6) not null,
    comment_id char(10) not null,
    INDEX id_index (user_id)
);

create table comment (
    comment_id char(10) PRIMARY KEY not null,
    content TINYTEXT not null,
    time varchar(30) not null,
    like_sum varchar(8) not null
);

create table user_com (
    user_id char(6) not null,
    review_id char(10) not null,
    comment_id char(10) not null,
    INDEX id_index (review_id),
    INDEX id_index2 (user_id)
);

create table review (
    review_id char(10) PRIMARY KEY not null,
    title varchar(50) not null,
    content TEXT not null,
    time varchar(30) not null,
    tag varchar(60) not null,
    picture varchar(50) not null,
    like_sum varchar(8) not null
);

