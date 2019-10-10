create table bankaccount (
    id serial primary key,
    name varchar unique,
    funds FLOAT 
);

insert into bankaccount (name, funds) values ('Lokier FellHeart', 12.3123);
insert into bankaccount (name, funds) values ('Brave Sir Robin', 42);
insert into bankaccount (name, funds) values ('Bob', 32.23132);