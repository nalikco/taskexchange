CREATE TABLE post_categories
(
    id serial not null unique,
    title varchar(255) not null,
    created_at timestamp default(now()) not null,
    deleted_at timestamp
);

CREATE TABLE posts
(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    status int not null,
    title varchar(255) not null,
    short text not null,
    text text not null,
    created_at timestamp default(now()) not null,
    deleted_at timestamp
);

CREATE TABLE post_category
(
    id serial not null unique,
    category_id int references post_categories(id) on delete cascade not null,
    post_id int references posts(id) on delete cascade not null
);
