CREATE DATABASE empresa;
CREATE TABLE departamento
(
depto_nro VARCHAR(10) PRIMARY KEY NOT NULL,
nombre_depto VARCHAR(30) NOT NULL,
localidad VARCHAR(30) NOT NULL
);
CREATE TABLE empleado
(
cod_emp VARCHAR(10) PRIMARY KEY NOT NULL,
nombre VARCHAR(30) NOT NULL,
apellido VARCHAR(30) NOT NULL,
puesto VARCHAR(30) NOT NULL,
fecha_alta DATE NOT NULL,
salario FLOAT NOT NULL,
comision FLOAT NOT NULL,
depto_nro VARCHAR(10) NOT NULL,
FOREIGN KEY (depto_nro) REFERENCES departamento(depto_nro)
);

insert into departamento (depto_nro, nombre_depto, localidad) values ('D-000-1', 'Software', 'Los Tigres');
insert into departamento (depto_nro, nombre_depto, localidad) values ('D-000-2', 'Sistemas', 'Guadalupe');
insert into departamento (depto_nro, nombre_depto, localidad) values ('D-000-3', 'Contabilidad', 'La Roca');
insert into departamento (depto_nro, nombre_depto, localidad) values ('D-000-4', 'Ventas', 'La Plata');

-- insert into empleado (cod_emp, nombre, apellido, puesto, fecha_alta, salario, comision, depto_nro) values ('E-0001', 'Camey', 'Bellerby', 'Account Representative II', '2022/2/19', '3040.70', '79.56', 'D-000-4');
insert into empleado (cod_emp, nombre, apellido, puesto, fecha_alta, salario, comision, depto_nro) values ('E-0002', 'Becki', 'Helkin', 'Data Coordiator', '2022/7/12', '2973.75', '57.92', 'D-000-2');
insert into empleado (cod_emp, nombre, apellido, puesto, fecha_alta, salario, comision, depto_nro) values ('E-0003', 'Herbie', 'Abadam', 'Sales Associate', '2022/7/12', '4308.62', '75.58', 'D-000-3');
insert into empleado (cod_emp, nombre, apellido, puesto, fecha_alta, salario, comision, depto_nro) values ('E-0004', 'Chaddy', 'Ruckledge', 'Chief Design Engineer', '2022/3/24', '9584.71', '55.76', 'D-000-4');
insert into empleado (cod_emp, nombre, apellido, puesto, fecha_alta, salario, comision, depto_nro) values ('E-0005', 'Bethany', 'Risbrough', 'Chief Design Engineer', '2022/1/9', '8431.07', '92.94', 'D-000-4');
insert into empleado (cod_emp, nombre, apellido, puesto, fecha_alta, salario, comision, depto_nro) values ('E-0006', 'Austen', 'Dunkerly', 'Computer Systems Analyst IV', '2022/5/5', '4477.16', '6.61', 'D-000-3');
insert into empleado (cod_emp, nombre, apellido, puesto, fecha_alta, salario, comision, depto_nro) values ('E-0007', 'Mito', 'Barchuk', 'Chief Design Engineer', '2022/7/20', '8779.16', '64.83', 'D-000-1');

