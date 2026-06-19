create table article_comment
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

INSERT INTO realworld.article_comment (id, author_username, body, created_at, updated_at, article_id) VALUES (2, 'xxxx1122', 'xxxxxxxxxxxxxxxx', '2023-11-12 02:47:23', '2023-11-12 02:47:23', 267);
INSERT INTO realworld.article_comment (id, author_username, body, created_at, updated_at, article_id) VALUES (3, 'xxxx1122', 'aaaaaaaaaaaa', '2023-11-12 02:47:26', '2023-11-12 02:47:26', 267);
