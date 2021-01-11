-- docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=Pass1234 -d mysql
create database ecommerce;
use ecommerce;

CREATE TABLE variant (
                          variant_id INT(6)  PRIMARY KEY,
                          variant_name VARCHAR(30) ,
                          mrp FLOAT(6) NOT NULL,
                          discount_price FLOAT(6),
                          size VARCHAR(20),
                          color VARCHAR(50),
                          product_id INT(6) NOT NULL
);

CREATE TABLE product (
                         product_id INT(6)  PRIMARY KEY,
                         product_name VARCHAR(30) NOT NULL,
                         description VARCHAR(100),
                         url VARCHAR(100)
);

CREATE TABLE product_category (
                         product_id INT(6) ,
                         category_id INT(6)
);



CREATE TABLE category (
                         category_id INT(6)  PRIMARY KEY,
                         category_name VARCHAR(30) NOT NULL,
                         parent_category_id INT(6)
);




