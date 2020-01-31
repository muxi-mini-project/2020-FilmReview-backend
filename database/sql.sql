create table user (
    user_id char(6) PRIMARY KEY not null,
    user_picture varchar(50) not null,
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
    content_sum varchar(3) not null
);

create table album_review (
    user_id char(6) not null,
    album_id char(3) not null,
    review_id int not null
);

create table collection (
    user_id char(6) not null,
    review_id int not null,
    INDEX id_index (user_id)
);

create table review_like (
    review_id int not null,
    user_id char(6) not null,
    INDEX id_index (review_id)
);

create table comment_like (
    user_id char(6) not null,
    comment_id int not null,
    review_id int not null,
    INDEX id_index (user_id)
);

create table comment (
    user_id char(6) not null,
    name varchar(30) not null,
    user_picture varchar(50) not null,
    review_id int not null,
    comment_id int  not null,
    content TINYTEXT not null,
    time timestamp default current_timestamp,
    like_sum int not null
);

create table user_review (
    user_id char(6) not null,
    name varchar(30) not null,
    user_picture varchar(30) not null,
    review_id int not null,
    title varchar(50) not null,
    content TEXT not null,
    time timestamp default current_timestamp,
    tag varchar(60) not null,
    picture varchar(50) not null,
    comment_sum int not null,
    like_sum int not null default 0
)
