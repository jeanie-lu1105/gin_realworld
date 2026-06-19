create table user
(
    username varchar(128)             not null
        primary key,
    password varchar(512)  default '' not null,
    email    varchar(256)  default '' not null,
    image    varchar(1024) default '' not null,
    bio      varchar(1024) default '' not null,
    constraint email
        unique (email)
)
    charset = utf8mb4;

INSERT INTO realworld.user (username, password, email, image, bio) VALUES ('ww1111', 'JDJhJDEwJHVkbTN2ejBwZFU5MTMyT3NmUy9sa3VpOFMxVExhNXZCV25qZE00STJQbVFNZkFRalRuYjVl', 'ww123@gmail.com', 'https://api.realworld.io/images/smiley-cyrus.jpeg', '');
INSERT INTO realworld.user (username, password, email, image, bio) VALUES ('xxxx1122', 'JDJhJDEwJDYwaG1sZUFELmhCR05WLk9TQlR1WU92RVhOS2hrRzJqMTRSSDFYRW1XaEp5Tms4cWtSd29H', 'xxxx123123@gmail.com', 'https://api.realworld.io/images/smiley-cyrus.jpeg', 'bbbbbbbbbbbbbbb');
