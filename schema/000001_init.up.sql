
create table users (
 id            serial not null unique,
 name          varchar(255) not null,
 surname       varchar(255) not null,
 email         varchar(255) not null unique,
 password      varchar(255) not null,
 registered_at timestamp not null default now()
);

create table refresh_tokens(
 id            serial not null unique,
 user_id       int references users (id) on delete cascade not null,
 token         varchar(255) not null,
 expires_at    timestamp not null
);

INSERT INTO users (name, surname, email, password) VALUES ('Taras', 'Tarkovskyi', 'taras6317492@gmail.com', 'wer12345');