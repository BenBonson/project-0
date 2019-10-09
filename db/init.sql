create table bankaccount (
    id serial primary key,
    name varchar unique,
    funds FLOAT 
);

insert into bankaccount (name, funds) values ('Bulbasaur', 12.3123);
insert into bankaccount (name, funds) values ('Ivysaur', 42);
insert into bankaccount (name, funds) values ('Venasaur', 32.23132);