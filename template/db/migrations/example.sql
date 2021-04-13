-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table example
(
    id int not null auto_increment,
    text varchar(255) not null,

    created_at datetime,
    updated_at datetime,
    deleted_at datetime,

    primary key(id)
);
-- initialize data [optional]
insert into role(text) values ('text');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table role;
