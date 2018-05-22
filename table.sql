create table app_config
(
  id    int auto_increment
    primary key,
  value varchar(255) not null,
  `key` varchar(64)  not null,
  constraint app_config_Id_uindex
  unique (id),
  constraint app_config_key_uindex
  unique (`key`)
);

create table candy
(
  id              int auto_increment,
  all_count       bigint            not null,
  remaining_count bigint            not null,
  candy_type      int               not null
    primary key,
  token_addr      varchar(255)      null,
  candy_label     varchar(64)       not null,
  rate            float default '1' not null,
  `decimal`       int               not null,
  candy_alias     varchar(64)       null,
  average         int               not null,
  constraint all_candy_token_id_uindex
  unique (id),
  constraint all_candy_token_candy_id_uindex
  unique (candy_type)
)
  charset = utf8;

create table game
(
  id     int auto_increment
    primary key,
  link   varchar(255) charset latin1 not null,
  sort   int                         null,
  status int                         null,
  icon   varchar(255) charset latin1 null,
  name   varchar(64) charset latin1  null,
  constraint game_id_uindex
  unique (id)
)
  charset = utf8;

create table game_candy
(
  id            int auto_increment
    primary key,
  game_field_id varchar(64)     not null,
  candy_pool    int default '0' not null,
  candy_consume int default '0' not null,
  game_id       int             not null,
  candy_id      int             not null,
  constraint gamecandy_id_uindex
  unique (id)
);

create table record
(
  id            int auto_increment
    primary key,
  addr          varchar(255) not null,
  count         bigint       not null,
  candy_id      int          not null,
  create_at     varchar(64)  not null,
  game_id       int          not null,
  game_field_id varchar(64)  not null,
  constraint record_id_uindex
  unique (id)
);

create table token
(
  id        int auto_increment
    primary key,
  addr      varchar(64) not null,
  count     bigint      not null,
  candy_id  int         not null,
  update_at varchar(64) null,
  constraint tokens_id_uindex
  unique (id)
);

create table user
(
  id        int auto_increment
    primary key,
  addr      varchar(255) null,
  name      varchar(64)  null,
  create_at varchar(64)  not null,
  status    int          not null,
  constraint user_id_uindex
  unique (id),
  constraint user_addr_uindex
  unique (addr)
);


