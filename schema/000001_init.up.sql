CREATE TABLE users
(
    id serial not null unique,
    email varchar(255) not null unique,
    password_hash varchar(255) not null,
    username varchar(255) not null,
    type int not null,
    balance float default(0.00) not null,
    points int default (0) not null,
    last_online timestamp default(now()) not null,
    created_at timestamp default(now()) not null,
    deleted_at timestamp
);

CREATE TABLE options
(
    id serial not null unique,
    parent_id int references options(id) on delete cascade,
    title varchar(255) not null,
    price float not null,
    created_at timestamp default(now()) not null,
    deleted_at timestamp
);

CREATE TABLE tasks
(
    id serial not null unique,
    customer_id int references users(id) on delete cascade not null,
    status int not null,
    amount int not null,
    delivery_date timestamp not null,
    link varchar(255) not null,
    description text not null,
    created_at timestamp default(now()) not null,
    deleted_at timestamp
);

CREATE TABLE task_options
(
    id serial not null unique,
    task_id int references tasks(id) on delete cascade not null,
    option_id int references options(id) on delete cascade not null
);

CREATE TABLE offers
(
    id serial not null unique,
    performer_id int references users(id) on delete cascade not null,
    task_id int references tasks(id) on delete cascade not null,
    status int default(0) not null,
    created_at timestamp default(now()) not null,
    deleted_at timestamp
);

CREATE TABLE orders
(
    id serial not null unique,
    offer_id int references offers(id) on delete cascade not null,
    status int default(0) not null,
    canceled_user_id int references users(id) on delete cascade,
    return_comment text,
    surrender_comment text,
    cancel_comment text,
    created_at timestamp default(now()) not null,
    deleted_at timestamp
);

CREATE TABLE events
(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    message text not null,
    link text not null,
    viewed_at timestamp,
    created_at timestamp default(now()) not null,
    deleted_at timestamp
);