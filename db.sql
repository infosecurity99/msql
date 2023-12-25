CREATE TABLE countries (
id SERIAL  primary key  not   null ,
name VARCHAR(60)  not null,  
 currency   VARCHAR(10)   not  null
)

CREATE TABLE  cities(
    id SERIAL ,
    name VARCHAR(30),
    population  int, 
    countries_id int references countries(id)
)
