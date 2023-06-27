drop table if exists stocks;
create table if not exists stocks (
  id serial,
  name varchar(200),
  description text,
  image_id varchar(32)
);