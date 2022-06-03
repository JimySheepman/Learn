/*
 * JOIN iki veya daha fazla tablonun birleştirir.
*/

CREATE TABLE basket_a (
    a INT PRIMARY KEY,
    fruit_a VARCHAR (100) NOT NULL
);

CREATE TABLE basket_b (
    b INT PRIMARY KEY,
    fruit_b VARCHAR (100) NOT NULL
);

INSERT INTO basket_a (a, fruit_a)
VALUES
    (1, 'Apple'),
    (2, 'Orange'),
    (3, 'Banana'),
    (4, 'Cucumber');

INSERT INTO basket_b (b, fruit_b)
VALUES
    (1, 'Orange'),
    (2, 'Apple'),
    (3, 'Watermelon'),
    (4, 'Pear');


/*
 * INNER JOIN iki veya daha fazla tablonun birleştirir.
 * 2 kümenin kesişim noktasıdır. yani 2 kümedeki ortak olan
 * elamanları döndürür. (A n B)
*/
SELECT
    a,
    fruit_a,
    b,
    fruit_b
FROM
    basket_a
INNER JOIN basket_b
    ON fruit_a = fruit_b;

/*
 * LEFT JOIN iki veya daha fazla tablonun birleştirir.
 * A kümesini göndürür ve karşılık gelmeten B kümesi elamanları
 * için NULL döner. ((A\B) + (A n B))
 * wehere koşulu eklenirse  LEFT OUTER JOIN olur. (A\B)
*/

SELECT
    a,
    fruit_a,
    b,
    fruit_b
FROM
    basket_a
LEFT JOIN basket_b 
    ON fruit_a = fruit_b;


SELECT
    a,
    fruit_a,
    b,
    fruit_b
FROM
    basket_a
LEFT JOIN basket_b 
    ON fruit_a = fruit_b
WHERE b IS NULL;

/*
 * RIGHT JOIN iki veya daha fazla tablonun birleştirir.
 * B kümesini göndürür ve karşılık gelmeten A kümesi elamanları
 * için NULL döner. ((B\A) + (A n B))
 * wehere koşulu eklenirse  RIGHT OUTER JOIN olur. (B\A)
*/

SELECT
    a,
    fruit_a,
    b,
    fruit_b
FROM
    basket_a
RIGHT JOIN basket_b 
    ON fruit_a = fruit_b;

SELECT
    a,
    fruit_a,
    b,
    fruit_b
FROM
    basket_a
RIGHT JOIN basket_b 
    ON fruit_a = fruit_b
WHERE a IS NULL;


/*
 * FULL JOIN iki veya daha fazla tablonun birleştirir.
 * A ve B kümesinin tüm elamanlarını birleştiri. (A U B)
 * wehere koşulu eklenirse  FULL OUTER JOIN olur. ((A U B) - (A n B))
*/

SELECT
    a,
    fruit_a,
    b,
    fruit_b
FROM
    basket_a
FULL OUTER JOIN basket_b 
    ON fruit_a = fruit_b;

SELECT
    a,
    fruit_a,
    b,
    fruit_b
FROM
    basket_a
FULL JOIN basket_b 
    ON fruit_a = fruit_b
WHERE a IS NULL OR b IS NULL;

/*
 * CROSS JOIN iki veya daha fazla tablonun birleştirir.
 * kartezyan çarpımını üretir. yan tümleci yoktur
*/

DROP TABLE IF EXISTS T1;
CREATE TABLE T1 (label CHAR(1) PRIMARY KEY);

DROP TABLE IF EXISTS T2;
CREATE TABLE T2 (score INT PRIMARY KEY);

INSERT INTO T1 (label)
VALUES
	('A'),
	('B');

INSERT INTO T2 (score)
VALUES
	(1),
	(2),
	(3);

SELECT *
FROM T1
CROSS JOIN T2;

/*
 * NATURAL JOIN irleştirilmiş tablolardaki aynı sütun
 * adlarına dayalı olarak örtük bir birleşim oluşturan bir birleşimdir.
*/

DROP TABLE IF EXISTS categories;
CREATE TABLE categories (
	category_id serial PRIMARY KEY,
	category_name VARCHAR (255) NOT NULL
);

DROP TABLE IF EXISTS products;
CREATE TABLE products (
	product_id serial PRIMARY KEY,
	product_name VARCHAR (255) NOT NULL,
	category_id INT NOT NULL,
	FOREIGN KEY (category_id) REFERENCES categories (category_id)
);

INSERT INTO categories (category_name)
VALUES
	('Smart Phone'),
	('Laptop'),
	('Tablet');

INSERT INTO products (product_name, category_id)
VALUES
	('iPhone', 1),
	('Samsung Galaxy', 1),
	('HP Elite', 2),
	('Lenovo Thinkpad', 2),
	('iPad', 3),
	('Kindle Fire', 3);

SELECT * FROM products
NATURAL JOIN categories;