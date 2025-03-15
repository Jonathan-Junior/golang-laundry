-- DDL.sql
-- Database creation for Enigma Laundry Management System

-- Create database
CREATE DATABASE enigma_laundry;

-- Connect to the database
\c enigma_laundry

-- Create tables
CREATE TABLE customer (
    customer_id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    address VARCHAR(255) DEFAULT '',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE service (
    service_id INT PRIMARY KEY,
    service_name VARCHAR(255) NOT NULL,
    unit VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "order" (
    order_id INT PRIMARY KEY,
    customer_id INT REFERENCES customer(customer_id),
    order_date DATE NOT NULL,
    completion_date DATE,
    received_by VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE order_detail (
    order_detail_id SERIAL PRIMARY KEY,
    order_id INT REFERENCES "order"(order_id),
    service_id INT REFERENCES service(service_id),
    qty INT NOT NULL
);
