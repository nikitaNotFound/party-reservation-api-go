CREATE DATABASE organizers_service_database;

\c organizers_service_database;

CREATE TABLE organizers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    is_verified BOOLEAN,
    phone VARCHAR(20)
);

CREATE TABLE parties (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    description TEXT,
    cancelled_at TIMESTAMP NULL,
    is_draft BOOLEAN,
    is_approved BOOLEAN,
    is_published BOOLEAN
);
