create table bankaccount (
    name varchar unique,
    funds FLOAT,
    cfunds FLOAT
);

insert into bankaccount (name, funds, cfunds) values ('LokierFellheart', 123123, 123.32);
insert into bankaccount (name, funds, cfunds) values ('BraveSirRobin', 42, 98);
insert into bankaccount (name, funds, cfunds) values ('Bob', 32.23132, 653);