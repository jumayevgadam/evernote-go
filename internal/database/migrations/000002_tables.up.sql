alter table users
add column email varchar(128) not null unique check (email <> '');