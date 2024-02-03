-- Create Table
create table books (
	id serial,
	title varchar,
	author varchar,
	year varchar
);

-- Insert Data
insert into books (title, author, year) values('Golang is great', 'Mr. Great', '2012');
insert into books (title, author, year) values('C++ is greatest', 'Mr. C++', '2015');
insert into books (title, author, year) values('Goroutines', '"Mr. Goroutine', '2011');
insert into books (title, author, year) values('Golang good parts', 'Mr. Good', '2014');
insert into books (title, author, year) values('Golang concurrency', 'Mr. Currency', '2013');

select * from books;
