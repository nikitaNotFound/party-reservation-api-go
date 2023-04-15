CREATE DATABASE IF NOT EXISTS users_service_database;

\c users_service_database;

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  phone_number VARCHAR(20) NOT NULL,
  age INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS organizers (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS parties (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  start_date TIMESTAMP NOT NULL,
  end_date TIMESTAMP NOT NULL,
  description TEXT NOT NULL,
  cancelled_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS registrations (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL REFERENCES users(id),
  party_id INTEGER NOT NULL REFERENCES parties(id)
);

CREATE TABLE IF NOT EXISTS comments (
  id SERIAL PRIMARY KEY,
  message TEXT NOT NULL,
  user_id INTEGER NOT NULL REFERENCES users(id),
  party_id INTEGER NOT NULL REFERENCES parties(id)
);
