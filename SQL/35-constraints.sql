/*
 * Primary Key ->  Bir tablodaki bir satırı benzersiz bir şekilde tanımlamak
 *                 için kullanılan bir sütun veya bir sütun grubudur.
 * Foreign Key ->  Bir tablodaki başka bir tablonun birincil anahtarına 
 *                 başvuran bir sütun veya sütun grubudur.
 * CHECK -> Bir CHECKsütundaki değerlerin belirli bir gereksinimi karşılaması
 *          gerekip gerekmediğini belirtmenize olanak tanıyan bir tür kısıtlamadır.
 * UNIQUE -> Verilerin benzersizliğini doğru şekilde koruyan kısıtlamayı sağlar.
 * Not-Null ->  Veritabanı teorisinde NULL, bilinmeyeni veya eksik bilgiyi temsil eder. 
 *              NULL, boş bir dize veya sıfır sayısı ile aynı değildir. Bir tabloya bir 
 *              kişinin e-posta adresini eklemeniz gerektiğini varsayalım. Onun e-posta 
 *              adresini talep edebilirsiniz. Ancak, kişinin bir e-posta adresi olup 
 *              olmadığını bilmiyorsanız, e-posta adresi sütununa NULL değerini 
 *              ekleyebilirsiniz. Bu durumda NULL, kayıt sırasında e-posta adresinin
 *              bilinmediğini gösterir.
 */
DROP TABLE IF EXISTS contacts;

DROP TABLE IF EXISTS customers;

CREATE TABLE customers(
    customer_id INT GENERATED ALWAYS AS IDENTITY,
    customer_name VARCHAR(255) NOT NULL,
    PRIMARY KEY(customer_id)
);

CREATE TABLE contacts(
    contact_id INT GENERATED ALWAYS AS IDENTITY,
    customer_id INT,
    contact_name VARCHAR(255) NOT NULL,
    phone VARCHAR(15),
    email VARCHAR(100),
    PRIMARY KEY(contact_id),
    CONSTRAINT fk_customer FOREIGN KEY(customer_id) REFERENCES customers(customer_id) ON DELETE
    SET
        NULL
);

INSERT INTO
    customers(customer_name)
VALUES
    ('BlueBird Inc'),
    ('Dolphin LLC');

INSERT INTO
    contacts(customer_id, contact_name, phone, email)
VALUES
    (
        1,
        'John Doe',
        '(408)-111-1234',
        'john.doe@bluebird.dev'
    ),
    (
        1,
        'Jane Doe',
        '(408)-111-1235',
        'jane.doe@bluebird.dev'
    ),
    (
        2,
        'David Wright',
        '(408)-222-1234',
        'david.wright@dolphin.dev'
    );

DROP TABLE IF EXISTS employees;

CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR (50),
    last_name VARCHAR (50),
    birth_date DATE CHECK (birth_date > '1900-01-01'),
    joined_date DATE CHECK (joined_date > birth_date),
    salary numeric CHECK(salary > 0)
);

CREATE TABLE invoices(
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    qty numeric NOT NULL CHECK(qty > 0),
    net_price numeric CHECK(net_price > 0)
);