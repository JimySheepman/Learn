/*
 * Bir veritabanı transaction, bir veya daha fazla işlemden oluşan
 * tek bir iş birimidir. Klasik bir işlem örneği, bir hesaptan 
 * diğerine banka havalesidir. Eksiksiz bir işlem, gönderen ve alıcı
 * hesapları arasında bir denge sağlamalıdır. Bu, gönderen hesap 
 * X tutarı transfer ederse, alıcının ne eksik X ne de fazla miktarda
 * alacağı anlamına gelir. Bir PostgreSQL transaction Atomicity, Consistency,
 * Isolation ve Durability. Bu özellikler genellikle ACID olarak adlandırılır:

 * Atomicity, transaction ya hep ya hiç şeklinde tamamlanmasını garanti eder.
 * Consistency, veritabanına yazılan verilerde yapılan değişikliğin geçerli olmasını ve önceden tanımlanmış kurallara uymasını sağlar.
 * İzolasyon, işlem bütünlüğünün diğer işlemlere nasıl görüneceğini belirler.
 * Durability, taahhüt edilen işlemlerin veri tabanında kalıcı olarak saklanmasını sağlar. 
 */

DROP TABLE IF EXISTS accounts;

CREATE TABLE accounts (
    id INT GENERATED BY DEFAULT AS IDENTITY,
    name VARCHAR(100) NOT NULL,
    balance DEC(15,2) NOT NULL,
    PRIMARY KEY(id)
);

--Begin a transaction

-- BEGIN WORK; or BEGIN TRANSACTION; or BEGIN;
INSERT INTO accounts(name,balance)
VALUES('Bob',10000);


BEGIN;

INSERT INTO accounts(name,balance)
VALUES('Alice',10000);


SELECT 
    id,
    name,
    balance
FROM 
    accounts;

--Commit a transaction

-- COMMIT WORK; or COMMIT TRANSACTION; or COMMIT;
SELECT 
    id,
    name,
    balance
FROM 
    accounts;


-- start a transaction
BEGIN;

-- insert a new row into the accounts table
INSERT INTO accounts(name,balance)
VALUES('Alice',10000);

-- commit the change (or roll it back later)
COMMIT;

-- Bank account transfer example

-- start a transaction
BEGIN;

-- deduct 1000 from account 1
UPDATE accounts 
SET balance = balance - 1000
WHERE id = 1;

-- add 1000 to account 2
UPDATE accounts
SET balance = balance + 1000
WHERE id = 2; 

-- select the data from accounts
SELECT id, name, balance
FROM accounts;

-- commit the transaction
COMMIT;

--Rolling back a transaction

-- ROLLBACK WORK; or ROLLBACK TRANSACTION; or ROLLBACK;

-- begin the transaction
BEGIN;

-- deduct the amount from the account 1
UPDATE accounts 
SET balance = balance - 1500
WHERE id = 1;

-- add the amount from the account 3 (instead of 2)
UPDATE accounts
SET balance = balance + 1500
WHERE id = 3; 

-- roll back the transaction
ROLLBACK;