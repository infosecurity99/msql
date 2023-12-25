CREATE TABLE countries (
id uuid  primary key  not   null ,
name VARCHAR(60)  not null,  
 currency   VARCHAR(10)   not  null
);

CREATE TABLE  cities(
    id uuid primary key not null,
    name VARCHAR(30),
    population  int, 
    countries_id uuid references countries(id)
);

