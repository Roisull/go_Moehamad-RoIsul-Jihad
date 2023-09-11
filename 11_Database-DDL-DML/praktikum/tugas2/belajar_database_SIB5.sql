CREATE TABLE Products (
    product_id INT PRIMARY KEY,
    product_name VARCHAR(255),
    product_type VARCHAR(255),
    product_description TEXT,
    operator VARCHAR(255),
    payment_methods VARCHAR(255)
);

ALTER TABLE Products
MODIFY COLUMN product_id INT AUTO_INCREMENT;

use alta_online_shop;

CREATE TABLE Customers(
	customer_id INT PRIMARY KEY,
    customer_name VARCHAR(100),
    customer_address VARCHAR(100),
    customer_birth DATE,
    status_user VARCHAR(20),
    gender char,
    create_at DATETIME DEFAULT NOW(),
    update_at DATETIME DEFAULT NOW()
    );
    
alter table Customers
MODIFY COLUMN customer_id INT auto_increment;

CREATE TABLE Transactions (
    transaction_id INT AUTO_INCREMENT PRIMARY KEY,
    customer_id INT,
    product_id INT,
    transaction_date DATETIME,
    total_amount DECIMAL(10,2),
    FOREIGN KEY (customer_id) REFERENCES Customers(customer_id),
    FOREIGN KEY (product_id) REFERENCES Products(product_id)
);

CREATE TABLE TransactionDetails(
	detail_id INT AUTO_INCREMENT PRIMARY KEY,
    transaction_id INT,
    quantity INT,
    price_per_unit DECIMAL(10,2),
    FOREIGN KEY (transaction_id) REFERENCES Transactions(transaction_id));
    
CREATE TABLE Kurir(
	kurir_id INT AUTO_INCREMENT PRIMARY KEY,
    kurir_name VARCHAR(100),
    create_at DATETIME DEFAULT NOW(),
    update_at DATETIME DEFAULT NOW());
    
ALTER TABLE Kurir
ADD ongkos_dasar DECIMAL(10,2);

ALTER TABLE Kurir
RENAME TO Shipping;

DROP TABLE Shipping;

CREATE TABLE Payment_methods(
	payment_method_id INT AUTO_INCREMENT PRIMARY KEY,
    payment_method_name VARCHAR(50),
    payment_method_description VARCHAR(100),
    active_status BOOLEAN,
    create_at DATETIME DEFAULT NOW(),
    update_at DATETIME DEFAULT NOW());
    
ALTER TABLE Products
CHANGE COLUMN payment_methods payment_method_id INT;

ALTER TABLE Products
DROP COLUMN payment_method_id;

CREATE TABLE Payment_method_details(
	payment_method_detail_id INT auto_increment primary key,
    user_id INT,
    payment_method_id INT,
    account_number VARCHAR(100),
    expiration_date Date,
    active_status boolean,
    create_at DATETIME DEFAULT NOW(),
    update_at DATETIME DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES Customers(customer_id),
    FOREIGN KEY (payment_method_id) REFERENCES Payment_methods(payment_method_id)
    );
    
ALTER TABLE Payment_method_details
RENAME TO Payment_method_descriptions;

CREATE TABLE Addreses(
	address_id INT auto_increment PRIMARY KEY,
    street_address VARCHAR(100),
    city VARCHAR(100),
    state_province VARCHAR(100),
    postal_code VARCHAR(100),
    country VARCHAR(100),
    create_at DATETIME DEFAULT NOW(),
    update_at DATETIME DEFAULT NOW());
    

ALTER TABLE Customers
CHANGE COLUMN customer_address address_id INT;

ALTER TABLE Customers
DROP COLUMN address_id;

ALTER TABLE Customers
ADD COLUMN address_id INT,
ADD CONSTRAINT fk_address_id
FOREIGN KEY (address_id) REFERENCES Addreses(address_id);

ALTER TABLE Products
DROP COLUMN operator;

CREATE TABLE Operators(
	operator_id INT auto_increment primary key,
    operator_name varchar(255),
    create_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp on update current_timestamp);
    
CREATE TABLE Product_types(
	product_type_id INT auto_increment primary key,
    product_type_name varchar(255),
    create_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp on update current_timestamp);
    
ALTER TABLE Products
DROP COLUMN product_type;

ALTER TABLE Products
DROP COLUMN product_description;

ALTER TABLE Products
ADD COLUMN code varchar(50);

ALTER TABLE Products
ADD COLUMN status smallint;

ALTER TABLE Products
ADD COLUMN create_at timestamp default current_timestamp;
ALTER TABLE Products
ADD COLUMN update_at timestamp default current_timestamp on update current_timestamp;

ALTER TABLE Products
ADD COLUMN product_type_id INT,
ADD CONSTRAINT fk_product_type_id
FOREIGN KEY (product_type_id) REFERENCES Product_types(product_type_id);

ALTER TABLE Products
ADD COLUMN operator_id INT,
ADD CONSTRAINT fk_operator_id
FOREIGN KEY (operator_id) REFERENCES Operators(operator_id);

CREATE TABLE Product_descriptions(
	product_description_id INT auto_increment primary key,
    description TEXT,
    create_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp on update current_timestamp);
    
RENAME TABLE Transactiondetails TO Transaction_details;

SHOW CREATE TABLE Transactions;

ALTER TABLE Transactions
DROP COLUMN product_id;

ALTER TABLE Payment_methods
DROP COLUMN active_status;

ALTER TABLE Payment_methods
ADD COLUMN status smallint;

ALTER TABLE Payment_methods
drop column create_at;
ALTER TABLE Payment_methods
drop column update_at;

ALTER TABLE Payment_methods
ADD COLUMN create_at timestamp default current_timestamp;
ALTER TABLE Payment_methods
ADD COLUMN update_at timestamp default current_timestamp on update current_timestamp;

ALTER TABLE Payment_methods
drop column payment_method_description;

ALTER TABLE Customers
drop column status_user;

ALTER TABLE Customers
drop column customer_name;

ALTER TABLE Customers
drop column customer_birth;

SHOW CREATE TABLE Customers;

ALTER TABLE Customers
DROP FOREIGN KEY fk_address_id;

ALTER TABLE Customers
DROP column address_id;

ALTER TABLE Customers
add column status smallint;

ALTER TABLE Customers
add column dob date;

show create table Transactions;

ALTER TABLE Transactions
ADD column payment_method_id int,
add constraint fk_payment_method_id
foreign key (payment_method_id) references Payment_methods(payment_method_id);

ALTER TABLE Transactions
drop column transaction_date;
ALTER TABLE Transactions
drop column total_amount;

ALTER TABLE Transactions
add column status varchar(10);
ALTER TABLE Transactions
add column total_qty int;
ALTER TABLE Transactions
add column total_price numeric(25,2);

ALTER TABLE Transactions
add column create_at timestamp default current_timestamp;

ALTER TABLE Transactions
add column update_at timestamp default current_timestamp on update current_timestamp;

drop table transaction_details;

create table Transaction_details(
	transaction_id INT,
    product_id INT,
    status varchar(10),
    qty INT,
    price numeric(25,2),
    create_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp on update current_timestamp
    );

alter table Transaction_details
add constraint fk_transaction_id
foreign key (transaction_id) references Transactions(transaction_id);

alter table Transaction_details
add constraint fk_product_id
foreign key (product_id) references Products(product_id);

drop table addreses;

drop table Payment_method_descriptions;




