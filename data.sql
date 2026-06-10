create table if not exists realworld.article
(
    id              bigint auto_increment
    primary key,
    author_username varchar(255)                            not null,
    title           varchar(4096)                           not null,
    slug            varchar(4096)                           not null,
    body            text                                    not null,
    description     varchar(2048) default ''                not null,
    tag_list        varchar(1024) default '[]'              not null,
    created_at      datetime      default CURRENT_TIMESTAMP not null,
    updated_at      datetime      default CURRENT_TIMESTAMP not null
    )
    charset = utf8mb4;

create table if not exists article_comment
(
    id              bigint auto_increment
    primary key,
    author_username varchar(255)                       not null,
    body            text                               not null,
    created_at      datetime default CURRENT_TIMESTAMP not null,
    updated_at      datetime default CURRENT_TIMESTAMP not null,
    article_id      bigint   default 0                 not null
    )
    charset = utf8mb4;

create table if not exists popular_tags
(
    name varchar(255) not null
    primary key
    )
    charset = utf8mb4;

create table if not exists user
(
    username varchar(128)  not null primary key,
    password varchar(512)  default '' not null,
    email    varchar(256)  default '' not null,
    image    varchar(1024) default '' not null,
    bio      varchar(1024) default '' not null,
    constraint email
    unique (email)
    )
    charset = utf8mb4;

