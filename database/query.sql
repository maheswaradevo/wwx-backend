CREATE DATABASE wwx;

use wwx;

create table users(
	id int not null auto_increment,
	username varchar(255) not null,
	password varchar (255) not null,
	primary key(id)
);

insert into users (username, password, role) values('devopunda','$2a$12$orxEIgKw2V7O5UKrAKtDDO/crKpA7jGQeBlBARiU9kzxslJdqSzuu
', 'client');

select * from users;

update users set password='$2a$12$.yf9vlgyKKC0.CvL6pdM0.N5ZPDyViWFGnUXsB/VR3Ce.x0DRGXc2
' where id = 1;

alter table users 
add column role varchar(255) not null;

update users set role = 'admin' where id=1;

update users set password = '$2a$12$eYfFCXK.pvnt1kIM/L.ENOIHDTgx3OhjwT4yNReMqQp5IjtZL8wCa
' where id = 3; 

create table projects (
	id int not null auto_increment,
	project_name varchar(40) not null,
	client_name varchar(30) not null,
	deadline timestamp not null,
	status varchar(20) not null,
	budget int,
	proposal_link varchar(40),
	assign varchar(20),
	created_at timestamp,
	user_id int,
	primary key(id),
	foreign key(user_id) references users(id)
);

select * from projects p ;

insert into users (username, password, role) values('maheswaradevo','$2a$12$eHq.mvNj7kKlx9j6xq3EHeCtj6zPSpzoVXr4m26QJ8nLy6jG9Pcoy
', Staff);

alter table projects 
add column resource_link varchar(255);


select * from projects;

SELECT * FROM projects WHERE project_name LIKE 'Project%';

update projects set resource_link = "" where id = 2;

alter table projects 
add column user_id int;

alter table projects 
add foreign key (user_id) references users(id);

update projects set budget = 0 where id = 13 ;

update projects set user_id = 3 where id = 14;

SELECT p.id, p.project_name, p.client_name, p.deadline, p.status, p.budget, p.proposal_link, p.assign, p.resource_link, p.user_id FROM projects p INNER JOIN users u ON u.id = p.user_id WHERE u.id = 1;

alter table projects 
add column created_at timestamp default current_timestamp;

drop table projects;

update projects set resource_link = 'link' where id = 1;

alter table projects 
drop column maintenance;

select * from projects;
SELECT * FROM projects WHERE project_name LIKE 'web%%';

alter table projects 
add column maintenance boolean not null default 0;

alter table projects 
modify column proposal_link varchar(255);

select * from projects p where user_id = 3;

insert into projects(project_name, client_name, deadline, status, budget, proposal_link, assign, user_id, resource_link, created_at, maintenance)
values ('Web Jualan', 'Bagus', '2023-06-07', 'process', 500, 'link', 'Made', '3', 'link', current_timestamp, 0),
		('Shopify', 'Karen', '2023-10-07', 'pending', 200, 'link', 'Putu', '3', 'link', current_timestamp, 1),
		('Web Pribadi', 'Wayan', '2023-09-10', 'complete', 600, 'link', 'Indira', '3', 'link', current_timestamp, 0),
		('Web Skripsi', 'Kurniawan', '2023-07-24', 'process', 150, 'link', 'Bianca', '3', 'link', current_timestamp, 1);
		
	
	
update projects set client_name = 'Karen' where id = 18;
	
	
SELECT * FROM projects WHERE maintenance = 1 and user_id != 1;	
	
	
	
	
	
	
	
