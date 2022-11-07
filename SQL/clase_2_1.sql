-- 1 Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.
select concat(e.nombre, e.apellido) Empleado, e.puesto, d.localidad
from empleado e
inner join departamento d
on  e.depto_nro = d.depto_nro;

-- 2 Visualizar los departamentos con más de cinco empleados.
select d.nombre_depto
from empleado e
inner join departamento d
on d.depto_nro = e.depto_nro
group by e.depto_nro
having count(e.depto_nro) > 1;

-- 3 Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que 
-- ‘Mito Barchuk’.
select e.nombre, e.apellido, e.salario, d.nombre_depto, e.puesto
from empleado e
inner join departamento d
on d.depto_nro = e.depto_nro
where e.puesto = (select puesto from empleado where nombre like 'mito' and apellido like 'Barchuk');

-- 4 Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.
select e.*, d.nombre_depto
from empleado e
inner join departamento d
on e.depto_nro = d.depto_nro
where d.nombre_depto like 'Contabilidad';

-- 5 Mostrar el nombre del empleado que tiene el salario más bajo.
select nombre, salario
from empleado
where salario = (select min(salario) from empleado);

-- 6 Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.
select e.*
from empleado e
inner join departamento d
on e.depto_nro = d.depto_nro
where d.nombre_depto like 'Ventas'
and
salario = (select max(salario) from empleado);


-- ----------------------------------- Ejercicio 2 //*
-- 1 Listar los datos de los autores.
select * from autor;

-- 2 Listar nombre y edad de los estudiantes
select nombre, edad from estudiante;

-- 3 Qué estudiantes pertenecen a la carrera informática?
select * from estudiante 
where carrera like 'informatica';

-- 4 ¿Qué autores son de nacionalidad francesa o italiana?
select * from autor
where nacionalidad like 'francesa' or nacionalidad like 'italiana';

-- 5 ¿Qué libros no son del área de internet?
select * from libro where area not like 'internet';

-- 6 Listar los libros de la editorial Salamandra.
select * from libro where editorial like 'Salamandra';

-- 7 Listar los datos de los estudiantes cuya edad es mayor al promedio.
select * from estudiante
where edad > avg(edad);

-- 8 Listar los nombres de los estudiantes cuyo apellido comience con la letra G.
select nombre, apellido
from estudiante 
where apellido like 'G%';

-- 9 Listar los autores del libro “El Universo: Guía de viaje”. (Se debe listar solamente los nombres).
select a.nombre
from autor a
inner join libro l
on a.idLibro = l. idLibro
where l.titulo like 'El Universo: Guía de viaje';

-- 10 ¿Qué libros se prestaron al lector “Filippo Galli”?
select l.nombre
from prestamo p
inner join libro l
on l.idLibro = p.idLibro
inner join estudiante e
on p.idLector = e.idLector;

-- 11 Listar el nombre del estudiante de menor edad.
select nombre
from estudiante
order by edad desc
limit 1;

select nombre
from estudiante
where edad = (select MIN(edad) from estudiante);

-- 12 Listar nombres de los estudiantes a los que se prestaron libros de Base de Datos.
select e.nombre
from prestamo p
inner join estudiante e
on p.idLector = e.idLector
inner join libro l
on l.idLibro = p.idLibro
where l.area like 'Base de Datos';

-- 13 Listar los libros que pertenecen a la autora J.K. Rowling.
select l.* 
from autor a
inner join libro l
on l.idLibro = a.idLibro;


select l.titulo 
from libro l where l.idLibro in (
	select la.idLibro
	from autor a
	inner join libroautor la
	on a.idAutor = la.idAutor
	where a.nombre like 'J.K. Rowling');


-- 14 Listar títulos de los libros que debían devolverse el 16/07/2021.
select l.titulo
from libro l
inner join prestamo p
where fechaDevolucion = '16/07/2021';


