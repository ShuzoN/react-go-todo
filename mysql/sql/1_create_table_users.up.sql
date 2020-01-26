create table if not exists users(
  id int not null auto_increment,
  name varchar(255) not null default '',
  primary key (id)
);
