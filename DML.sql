-- DML.sql
-- SQL queries used in Enigma Laundry Management System

-- CUSTOMER QUERIES

-- Check if customer ID exists
SELECT EXISTS(SELECT 1 FROM customer WHERE customer_id = $1);

-- Create new customer
INSERT INTO customer (customer_id, name, phone, address) VALUES ($1, $2, $3, $4);

-- View list of customers
SELECT customer_id, name, phone FROM customer ORDER BY customer_id;

-- View customer details by ID
SELECT customer_id, name, phone, address, created_at, updated_at 
FROM customer 
WHERE customer_id = $1;

-- Update customer
UPDATE customer 
SET name = $1, phone = $2, address = $3, updated_at = CURRENT_TIMESTAMP 
WHERE customer_id = $4;

-- Check if customer is used in orders
SELECT EXISTS(SELECT 1 FROM "order" WHERE customer_id = $1);

-- Delete customer
DELETE FROM customer WHERE customer_id = $1;

-- SERVICE QUERIES

-- Check if service ID exists
SELECT EXISTS(SELECT 1 FROM service WHERE service_id = $1);

-- Create new service
INSERT INTO service (service_id, service_name, unit, price) VALUES ($1, $2, $3, $4);

-- View list of services
SELECT service_id, service_name, unit, price FROM service ORDER BY service_id;

-- View service details by ID
SELECT service_id, service_name, unit, price, created_at, updated_at 
FROM service 
WHERE service_id = $1;

-- Update service
UPDATE service 
SET service_name = $1, unit = $2, price = $3, updated_at = CURRENT_TIMESTAMP 
WHERE service_id = $4;

-- Check if service is used in order details
SELECT EXISTS(SELECT 1 FROM order_detail WHERE service_id = $1);

-- Delete service
DELETE FROM service WHERE service_id = $1;

-- ORDER QUERIES

-- Check if order ID exists
SELECT EXISTS(SELECT 1 FROM "order" WHERE order_id = $1);

-- Create new order
INSERT INTO "order" (order_id, customer_id, order_date, received_by) 
VALUES ($1, $2, $3, $4);

-- Add order detail
INSERT INTO order_detail (order_id, service_id, qty) 
VALUES ($1, $2, $3);

-- Complete order
UPDATE "order" 
SET completion_date = $1, updated_at = CURRENT_TIMESTAMP 
WHERE order_id = $2;

-- View list of orders
SELECT o.order_id, c.name, o.order_date, o.completion_date, o.received_by
FROM "order" o
JOIN customer c ON o.customer_id = c.customer_id
ORDER BY o.order_id;

-- View order details
SELECT o.order_id, o.customer_id, c.name, o.order_date, o.completion_date, 
       o.received_by, o.created_at, o.updated_at
FROM "order" o
JOIN customer c ON o.customer_id = c.customer_id
WHERE o.order_id = $1;

-- View order items
SELECT od.order_detail_id, s.service_name, s.unit, s.price, od.qty, 
       (s.price * od.qty) as subtotal
FROM order_detail od
JOIN service s ON od.service_id = s.service_id
WHERE od.order_id = $1
ORDER BY od.order_detail_id;
