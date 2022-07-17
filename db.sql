-- noinspection SqlNoDataSourceInspectionForFile

create  database api_todo;
create table todos (id serial primary key, name varchar, description text, done bool);