create table cliente (
	id_cliente int primary key,
    dni int,
    name varchar(30),
    lastname varchar(30),
	birthdate date,
    province varchar(30),
    city varchar(30)
);

create table planes(
	id_plan smallint primary key,
    velocidad int,
    precio float,
    descuento float
);
alter table planes
add cliente_id int;

alter table planes
add constraint
foreign key (cliente_id) references cliente(id_cliente);






