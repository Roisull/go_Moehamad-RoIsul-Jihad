use alta_online_shop;

INSERT INTO operators (operator_name, create_at, update_at)
VALUES
    ('Operator 1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Operator 2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Operator 3', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Operator 4', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Operator 5', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
    
    select * from Operators;
    
INSERT INTO product_types (product_type_name, create_at, update_at)
VALUES
    ('Product Type A', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Product Type B', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Product Type C', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
    select * from product_types;
    
INSERT INTO products (product_name, code, status, create_at, update_at, product_type_id, operator_id)
VALUES
    ('Product 1', 'P001', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 1, 3),
    ('Product 2', 'P002', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 1, 3);
    select * from products;
    
INSERT INTO products (product_name, code, status, create_at, update_at, product_type_id, operator_id)
VALUES
    ('Product 3', 'P003', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 2, 1),
    ('Product 4', 'P004', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 2, 1),
    ('Product 5', 'P005', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 2, 1);
	select * from products;
    
    INSERT INTO products (product_name, code, status, create_at, update_at, product_type_id, operator_id)
VALUES
    ('Product 6', 'P006', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 3, 4),
    ('Product 7', 'P007', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 3, 4),
    ('Product 8', 'P008', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 3, 4);
    select * from products;
    
INSERT INTO product_descriptions (description, create_at, update_at)
VALUES
    ('Deskripsi Produk 1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Deskripsi Produk 2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Deskripsi Produk 3', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Deskripsi Produk 4', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Deskripsi Produk 5', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Deskripsi Produk 6', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Deskripsi Produk 7', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Deskripsi Produk 8', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
    select * from product_descriptions;
    
INSERT INTO payment_methods(payment_method_name, status, create_at, update_at)
values
	('cash', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('cash', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('credit card', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
    select * from payment_methods;
    
INSERT INTO customers(gender, create_at, update_at, status, dob)
values
('F', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 1, '2023-09-07'),
('M', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 1, '2023-09-07'),
('F', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 1, '2023-09-07'),
('F', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 1, '2023-09-07'),
('M', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 1, '2023-09-07');
select * from customers;
    
INSERT INTO transactions (customer_id, payment_method_id, status, total_qty, total_price, create_at, update_at)
SELECT customer_id, 1, 'completed', 3, 150.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP 
FROM customers WHERE customer_id = 1;

INSERT INTO transactions (customer_id, payment_method_id, status, total_qty, total_price, create_at, update_at)
SELECT customer_id, 2, 'completed', 2, 100.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM customers WHERE customer_id = 2;

INSERT INTO transactions (customer_id, payment_method_id, status, total_qty, total_price, create_at, update_at)
SELECT customer_id, 3, 'completed', 1, 50.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM customers WHERE customer_id = 3;
select * from transactions;
select * from transaction_details;


INSERT INTO transaction_details (transaction_id, product_id, status, qty, price, create_at, update_at)
VALUES
    (1, 1, 'completed', 2, 50.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), 
    (1, 2, 'completed', 1, 30.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), 
    (1, 3, 'completed', 3, 80.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP); 
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price, create_at, update_at)
VALUES
    (2, 4, 'completed', 2, 70.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), 
    (2, 5, 'completed', 1, 40.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), 
    (2, 6, 'completed', 3, 90.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP); 
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price, create_at, update_at)
VALUES
    (3, 7, 'completed', 2, 60.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), 
    (3, 8, 'completed', 1, 35.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), 
    (3, 8, 'completed', 3, 75.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP); 
select * from transaction_details;
    
    SELECT * FROM products WHERE product_id IN (7, 8, 8);
    
SELECT customer_id, gender
FROM customers
WHERE gender IN ('M');

SELECT * FROM products
WHERE product_id = 3;

select * from customers;

ALTER TABLE Customers
ADD COLUMN customer_name VARCHAR(255);

UPDATE customers
SET customer_name = 
    CASE 
        WHEN customer_id = 1 THEN 'Rona'
        WHEN customer_id = 2 THEN 'Dono'
        WHEN customer_id = 3 THEN 'Susili'
        WHEN customer_id = 4 THEN 'Baraya'
        WHEN customer_id = 5 THEN 'Baqiha'
    END
WHERE customer_id IN (1, 2, 3, 4, 5);

SELECT * FROM customers
WHERE create_at >= DATE_SUB(CURRENT_DATE(), INTERVAL 7 DAY)
  AND create_at <= CURRENT_DATE()
  AND customer_name LIKE '%a%';

SELECT COUNT(*) FROM customers
WHERE gender = 'F';

SELECT *
FROM customers
ORDER BY customer_name ASC;

SELECT * FROM products LIMIT 5;


