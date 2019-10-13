create table bankaccount (
    name varchar unique,
    funds FLOAT,
    cfunds FLOAT
);

insert into bankaccount (name, funds, cfunds) values ('Lokier FellHeart', 123123, 123.32);
insert into bankaccount (name, funds, cfunds) values ('Brave Sir Robin', 42, 98);
insert into bankaccount (name, funds, cfunds) values ('Bob', 32.23132, 653);