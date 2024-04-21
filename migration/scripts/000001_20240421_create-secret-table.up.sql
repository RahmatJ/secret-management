create table if not exists secret_management (
    id varchar(16) primary key unique,
    user_id varchar(50) not null,
    api_key varchar(100) not null,
    expireddate varchar(15) not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create index if not exists secret_management_user_id_created_at
    on secret_management (user_id, created_at)