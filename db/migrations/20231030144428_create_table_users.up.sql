create table users
(
    id           varchar(100) not null,
    name         varchar(100) not null,
    password     varchar(100) not null,
    email        varchar(255) not null,
    roles        varchar(100) null,
    address      varchar(255) null,
    house_number varchar(100) null,
    token        varchar(100) null,
    phone_number varchar(50)  null,
    city         varchar(100) null,
    created_at   bigint       not null,
    updated_at   bigint       not null,
    primary key (id)
) engine = InnoDB;
