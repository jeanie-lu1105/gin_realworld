create table popular_tags
(
    name varchar(255) not null
        primary key
)
    charset = utf8mb4;

INSERT INTO realworld.popular_tags (name) VALUES ('codebaseShow');
INSERT INTO realworld.popular_tags (name) VALUES ('implementations');
INSERT INTO realworld.popular_tags (name) VALUES ('introduction');
INSERT INTO realworld.popular_tags (name) VALUES ('welcome');
