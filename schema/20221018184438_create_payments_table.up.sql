CREATE TABLE payments
(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    type int not null,
    comment varchar(255) not null,
    sum float not null,
    created_at timestamp default(now()) not null,
    deleted_at timestamp
);