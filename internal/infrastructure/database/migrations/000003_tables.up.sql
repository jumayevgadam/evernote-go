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
