-- users table
create table if not exists users (
    id serial primary key,
    username varchar(128) not null unique check (username <> ''),
    password text,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

-- notebooks table
create table if not exists notebooks (
    id serial primary key,
    user_id int references users (id) on delete cascade,
    name varchar(127) not null check (name <> ''),
    description text,
    is_shared boolean default false,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

-- notes table
create table if not exists notes (
    id serial primary key,
    notebook_id int references notebooks (id) on delete cascade,
    title varchar(128) not null check (title <> ''),
    content text,
    is_archived boolean default false,
    is_deleted boolean default false,
    created_at timestamp default current_timestamp
);

-- tags table
create table if not exists tags (
    id serial primary key,
    user_id int references users (id) on delete cascade,
    name varchar(127) not null check (name <> ''),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

-- note_tags table 
create table if not exists note_tags (
    note_id int references notes (id),
    tag_id int references tags (id)
);

-- attachments table
create table if not exists attachments (
    id serial primary key,
    note_id int references notes (id) on delete cascade,
    file_name text,
    file_path text,
    created_at timestamp default current_timestamp
);

create type permission_enum as enum ('read', 'write');

-- shared_notebooks table
create table if not exists shared_notebooks (
    id serial primary key,
    notebook_id int references notebooks (id) on delete cascade,
    shared_user_id int references users (id) on delete cascade,
    permissions permission_enum not null default 'read',
    shared_at timestamp default current_timestamp
);

create table if not exists tasks (
    id serial primary key,
    note_id int references notes (id) on delete cascade, 
    title varchar(128) not null check (title <> ''),     
    description text,                                   
    is_completed boolean default false,               
    due_date timestamp,                                
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);
