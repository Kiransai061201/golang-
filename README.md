# golang-



docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres

docker exec -it some-postgres psql -U postgres

docker stop some-postgres
docker rm some-postgres

CREATE DATABASE testdb;

\c testdb

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age INT NOT NULL,
    gender VARCHAR(10),
    email VARCHAR(100),
    mobile VARCHAR(15),
    address TEXT
);



CREATE USER kiran WITH PASSWORD 'kiran0612';

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO kiran;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO kiran;
\q
