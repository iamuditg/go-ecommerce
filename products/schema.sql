
CREATE DATABASE IF NOT EXISTS ecommerce;


CREATE TABLE IF NOT EXISTS categories(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    image VARCHAR(255)
);


CREATE TABLE IF NOT EXISTS sellers (
    id serial PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    city VARCHAR(100),
    country VARCHAR(100),
    image VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS brands (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    image VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    description TEXT,
    rating DECIMAL(3,1),
    category_id INT,
    seller_id INT,
    brand_id INT,
    image VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories(id),
    FOREIGN KEY (seller_id) REFERENCES sellers(id),
    FOREIGN KEY (brand_id) REFERENCES brands(id)
);



INSERT INTO categories (name, image) VALUES
('Electronics', 'electronics.jpg'),
('Clothing', 'clothing.jpg'),
('Home Appliances', 'home_appliances.jpg');


INSERT INTO sellers (name, city, country, image) VALUES
                                                     ('Seller A', 'New York', 'USA', 'seller_a.jpg'),
                                                     ('Seller B', 'London', 'UK', 'seller_b.jpg'),
                                                     ('Seller C', 'Tokyo', 'Japan', 'seller_c.jpg');


INSERT INTO brands (name, description, image) VALUES
                                                  ('Brand X', 'Description of Brand X', 'brand_x.jpg'),
                                                  ('Brand Y', 'Description of Brand Y', 'brand_y.jpg'),
                                                  ('Brand Z', 'Description of Brand Z', 'brand_z.jpg');


INSERT INTO products (name, price, description, rating, category_id, seller_id, brand_id, image) VALUES
                                                                                                     ('Product 1', 10.99, 'Description of Product 1', 4.5, 1, 1, 1, 'product_1.jpg'),
                                                                                                     ('Product 2', 19.99, 'Description of Product 2', 3.8, 2, 2, 2, 'product_2.jpg'),
                                                                                                     ('Product 3', 14.99, 'Description of Product 3', 4.2, 1, 3, 3, 'product_3.jpg');

