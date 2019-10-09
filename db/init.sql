create table bankaccount (
    id integer primary key,
    name varchar unique,
    funds FLOAT 
);

insert into bankaccount values (1, 'Bulbasaur', 12.3123);
insert into bankaccount values (2, 'Ivysaur', 42);
insert into bankaccount values (3, 'Venasaur', 32.23132);