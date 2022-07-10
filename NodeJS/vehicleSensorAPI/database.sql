CREATE DATABASE "vehicle_sensor"
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

DROP TABLE IF EXISTS vehicles;
DROP TABLE IF EXISTS log_location;
DROP TABLE IF EXISTS devices;
DROP TABLE IF EXISTS devices_type;
DROP TABLE IF EXISTS log_temperature;


CREATE TABLE IF NOT EXISTS vehicles
(
    id SERIAL PRIMARY KEY NOT NULL,
    vehicle_plate varchar(20)  NOT NULL,
    current_status int2 NOT NULL,
    is_active bool NOT NULL
);


CREATE TABLE IF NOT EXISTS devices
(
    id SERIAL PRIMARY KEY NOT NULL,
    vehicle_id int8 NOT NULL,
    device_type_id int8  NOT NULL,
    device_name varchar(75) NOT NULL,
    is_online bool NOT NULL,
    is_active bool NOT NULL
);

CREATE TABLE IF NOT EXISTS  devices_type
(
    id SERIAL PRIMARY KEY NOT NULL,
    device_name varchar(75) NOT NULL,
    device_description varchar(255) NOT NULL,
    is_active bool NOT NULL
);


CREATE TABLE IF NOT EXISTS  log_temperature
(
    id SERIAL PRIMARY KEY NOT NULL,
    vehicle_id int8, 
    device_id  int8,
    read_data varchar(50),
    created_at timestamp
);


CREATE TABLE IF NOT EXISTS  log_location
(
    id SERIAL PRIMARY KEY NOT NULL,
    vehicle_id int8, 
    device_id  int8,
    latitude varchar(50),
    longtitude varchar(50),
    created_at timestamp
);
