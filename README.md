# eCommerce Project

This is an eCommerce project build using go programming. It provides APIs to retrieve and create products.

## Prerequisites

Before running the project, ensure that you have the following dependencies installed on you machine.

- Docker 
- Docker Compose

## Getting Started

Follow the steps below to run the project:

1. Clone the repository to your local machine:
   ```bash 
   git clone <repository.url>

2. Navigate to the project directory:
   ```bash
   cd ecommerce
 
3. Update the database configuration:
   Open `products/resources/yml/db.yaml` file and pdate the following properties.
   - `host` (should be the hostname or IP address of your database container)
   - `port` (should match the exposed port of your database container)
   - `user` (the database username)
   - `password` (the database password)
   - `name` (the name of the database)

4. Build and run the project using Docker compose:
   ```bash
   docker-compose up --build

the command will build the docker images and start the container for the application
and the database

5. Once the containers are running. you can access the application at `http://localhost:8000`

## API Endpoints
- GET/products - Retrieve all products
- POST/products - Create a new product