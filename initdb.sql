-- 评论表
create table comments(
  id varchar(64) primary key not null,
  video_id varchar(64),
  author_id int,
  content text,
  time datetime default CURRENT_TIMESTAMP
);
-- session
create table sessions(
  session_id varchar(256) primary key not null,
  ttl TinyText,
  login_name text
);

create table video_del_rec(video_id varchar(64) primary key not null);

-- 用户
create table users(
    id int primary key auto_increment not null,
    login_name VARCHAR(64) NOT NULL UNIQUE,
    pwd  text
);

-- 视频资源
create table video_info(
    id varchar(64) primary key not null,
    author_id  int,
    name text,
    display_ctime text,
    createe_time datetime default CURRENT_TIMESTAMP
);


