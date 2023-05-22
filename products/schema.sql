-- Create the database if it does not exist
CREATE DATABASE IF NOT EXISTS ecommerce;

-- Create the products table
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price INTEGER NOT NULL
    );

-- Insert dummy data
INSERT INTO products (name, price)VALUES('Product 1', 10),('Product 2', 20),('Product 3', 30);

-- Add additional queries or statements here
