-- Create the users table
CREATE TABLE users (
                       user_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                       username VARCHAR(50) NOT NULL,
                       email VARCHAR(60) NOT NULL,
                       password VARCHAR(255) NOT NULL
);

-- Create the products table with image URL
CREATE TABLE products (
                          product_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                          name VARCHAR(50) NOT NULL,
                          image_url TEXT, -- URL to the image
                          description TEXT NOT NULL,
                          price DECIMAL(10, 2) NOT NULL,
                          stock_quantity INT NOT NULL
);

-- Create the carts table
CREATE TABLE carts (
                       cart_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                       user_id INT REFERENCES users(user_id)
);

-- Create the cart_items table
CREATE TABLE cart_items (
                            cart_item_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                            cart_id INT REFERENCES carts(cart_id),
                            product_id INT REFERENCES products(product_id),
                            quantity INT NOT NULL CHECK (quantity > 0)
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_user_id ON carts(user_id);
CREATE INDEX idx_cart_items_product_id ON cart_items(product_id);
