CREATE TABLE users
(
    user_id SERIAL PRIMARY KEY,
    first_name varchar(50),
    last_name varchar(50),
    email varchar(50),
    date_created date not null default CURRENT_DATE
);