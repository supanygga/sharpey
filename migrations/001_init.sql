create table users(
    id uuid primary key,
    name text not null unique,
    is_bot bool,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp
);

create table messages(
    id uuid primary key,
    channel uuid references users(id) not null,
    sender uuid references users(id) not null,
    content text not null,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp
)

---- create above / drop below ----

drop table messages;

drop table users;
