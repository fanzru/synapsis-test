INSERT INTO categories (name) VALUES
('Electronics'),
('Books'),
('Clothing'),
('Home & Kitchen'),
('Sports & Outdoors');

INSERT INTO product (name, description, price, stock, category_id) VALUES
('Smartphone', 'A high-end smartphone with 128GB storage', 699.99, 50, (SELECT id FROM categories WHERE name = 'Electronics')),
('Laptop', 'A powerful laptop with 16GB RAM and 512GB SSD', 1299.99, 30, (SELECT id FROM categories WHERE name = 'Electronics')),
('Novel', 'A best-selling novel by a famous author', 19.99, 100, (SELECT id FROM categories WHERE name = 'Books')),
('T-shirt', 'A comfortable cotton t-shirt', 9.99, 200, (SELECT id FROM categories WHERE name = 'Clothing')),
('Blender', 'A high-performance blender for smoothies', 89.99, 40, (SELECT id FROM categories WHERE name = 'Home & Kitchen')),
('Basketball', 'Official size and weight basketball', 29.99, 60, (SELECT id FROM categories WHERE name = 'Sports & Outdoors'));