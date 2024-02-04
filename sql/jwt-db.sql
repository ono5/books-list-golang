-- Create TABLES
create table users (
	id serial primary key,
	email text not null unique,
	password text not null
);

-- Insert Data
insert into users (email, password) values ('test@example.com', '12345');
insert into users (email, password) values ('test123@example.com', 'abcd');

-- Show data
select * from users;

