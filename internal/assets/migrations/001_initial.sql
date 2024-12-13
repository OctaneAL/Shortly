-- +migrate Up

create table url_mappings
(
    id          bigserial primary key not null,
    original_url text                 not null,
    short_code   varchar(10)          not null unique,
    created_at   timestamp without time zone default now()
);

create index url_mappings_short_code_index on url_mappings (short_code);
create index url_mappings_original_url_index on url_mappings (original_url);

-- +migrate Down

drop index url_mappings_short_code_index;
drop index url_mappings_original_url_index;
drop table url_mappings;
