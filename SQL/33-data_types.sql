/*
 * Boolean->
 *           true -> 1,yes,y,t,true
 *           false -> 0,no,false,f
 *           null -> space
 * Character-> 
 *             CHAR(n) && VARCHAR(n)-> n yerine byte'ı belirtilir.
 *             TEXT-> ise sınırsız uzunluktaki karakter dizesidir.
 * Numeric-> 
 *           integers->
 *                      SMALLINT && SMALLSERIAL -> 2 byte
 *                      INT-> 4 byte
 *                      SERIAL -> AUTO_INCREMENT sağlar ve int il aynıdır.
 *                      BIGSERIAL-> 8 byte
 *           Floating-point number->
 *                                  float(n)-> max 8-byte
 *                                  real || float8 -> 4 byte
 *                                  numeric || numeric(p,s)-> kesir kısmını belirler.
 * Temporal data->
 *                DATE-> sadece tarih saklar
 *                TIME-> günün saati değerleri saklar
 *                TIMESTAMP-> hem tarih hem saat saklar
 *                TIMESTAMPTZ-> saat dilimine duyarlı bir timestamp'tır.
 *                INTERVAL->
 * Arrays-> bir tamsayı dizisini vb. dizi sütunlarında saklayabilirsiniz . Dizi, örneğin 
 *          haftanın günlerini, yılın aylarını saklamak gibi bazı durumlarda kullanışlıdır.
 * JSON-> JSON && JSONB deoplamayı sağlar. JSONB indexlemeyi destekler.
 * UUID->  RFC 4122 standartlarında veri üretir.
 * Special data types-> 
 *                     box-> dikdörtgen bir kutu
 *                     line-> bir dizi nokta
 *                     point-> geometrik bir sayı çifti
 *                     lseg-> bir çizgi parçası.
 *                     polygon-> kapalı bir geometrik şekil
 *                     inet-> bir IP4 adresi
 *                     macaddr-> bir MAC adresi
 *                     hstore -> key/value çiftlerini tek bir değerde depolamak için.
 */
-- JSON
CREATE TABLE orders (
    id serial NOT NULL PRIMARY KEY,
    info json NOT NULL
);

INSERT INTO
    orders (info)
VALUES
(
        '{ "customer": "Lily Bush", "items": {"product": "Diaper","qty": 24}}'
    ),
    (
        '{ "customer": "Josh William", "items": {"product": "Toy Car","qty": 1}}'
    ),
    (
        '{ "customer": "Mary Clark", "items": {"product": "Toy Train","qty": 2}}'
    );

SELECT
    info -> 'customer' AS customer
FROM
    orders;

SELECT
    info ->> 'customer' AS customer,
    info -> 'items' ->> 'product' AS product
FROM
    orders
WHERE
    CAST (info -> 'items' ->> 'qty' AS INTEGER) = 2 -- hstore
    CREATE EXTENSION hstore;

CREATE TABLE books (
    id serial primary key,
    title VARCHAR (255),
    attr hstore
);

INSERT INTO
    books (title, attr)
VALUES
    (
        'PostgreSQL Tutorial',
        '"paperback" => "243",
	   "publisher" => "postgresqltutorial.com",
	   "language"  => "English",
	   "ISBN-13"   => "978-1449370000",
		 "weight"    => "11.2 ounces"'
    );

INSERT INTO
    books (title, attr)
VALUES
    (
        'PostgreSQL Cheat Sheet',
        '
"paperback" => "5",
"publisher" => "postgresqltutorial.com",
"language"  => "English",
"ISBN-13"   => "978-1449370001",
"weight"    => "1 ounces"'
    );

SELECT
    attr
FROM
    books;

SELECT
    title,
    attr -> 'publisher' as publisher,
    attr
FROM
    books
WHERE
    attr ? 'publisher';

-- Array
CREATE TABLE contacts (
    id serial PRIMARY KEY,
    name VARCHAR (100),
    phones TEXT []
);

INSERT INTO
    contacts (name, phones)
VALUES
(
        'John Doe',
        ARRAY [ '(408)-589-5846','(408)-589-5555' ]
    );

INSERT INTO
    contacts (name, phones)
VALUES
('Lily Bush', '{"(408)-589-5841"}'),
    (
        'William Gate',
        '{"(408)-589-5842","(408)-589-58423"}'
    );

SELECT
    name,
    phones
FROM
    contacts;

-- UUID ekleneceği zaman string olarak sonradanda eklenebilri.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SELECT
    uuid_generate_v1();

SELECT
    uuid_generate_v4();

CREATE TABLE contacts (
    contact_id uuid DEFAULT uuid_generate_v4 (),
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    phone VARCHAR,
    PRIMARY KEY (contact_id)
);

INSERT INTO
    contacts (
        first_name,
        last_name,
        email,
        phone
    )
VALUES
    (
        'John',
        'Smith',
        'john.smith@example.com',
        '408-237-2345'
    ),
    (
        'Jane',
        'Smith',
        'jane.smith@example.com',
        '408-237-2344'
    ),
    (
        'Alex',
        'Smith',
        'alex.smith@example.com',
        '408-237-2343'
    );

SELECT
    *
FROM
    contacts;

-- TIME
CREATE TABLE shifts (
    id serial PRIMARY KEY,
    shift_name VARCHAR NOT NULL,
    start_at TIME NOT NULL,
    end_at TIME NOT NULL
);

INSERT INTO
    shifts(shift_name, start_at, end_at)
VALUES
('Morning', '08:00:00', '12:00:00'),
    ('Afternoon', '13:00:00', '17:00:00'),
    ('Night', '18:00:00', '22:00:00');

SELECT
    *
FROM
    shifts;

SELECT
    CURRENT_TIME;

SELECT
    CURRENT_TIME(5);

SELECT
    LOCALTIME;

SELECT
    LOCALTIME(0);

SELECT
    LOCALTIME AT TIME ZONE 'UTC-7';

-- Interval
SET
    intervalstyle = 'sql_standard';

SELECT
    INTERVAL '6 years 5 months 4 days 3 hours 2 minutes 1 second';

SET
    intervalstyle = 'postgres';

SELECT
    INTERVAL '6 years 5 months 4 days 3 hours 2 minutes 1 second';

SET
    intervalstyle = 'postgres_verbose';

SELECT
    INTERVAL '6 years 5 months 4 days 3 hours 2 minutes 1 second';

SET
    intervalstyle = 'iso_8601';

SELECT
    INTERVAL '6 years 5 months 4 days 3 hours 2 minutes 1 second';

-- Timestamp
CREATE TABLE timestamp_demo (ts TIMESTAMP, tstz TIMESTAMPTZ);

SET
    timezone = 'America/Los_Angeles';

SHOW TIMEZONE;

INSERT INTO
    timestamp_demo (ts, tstz)
VALUES
(
        '2016-06-22 19:10:25-07',
        '2016-06-22 19:10:25-07'
    );

SET
    timezone = 'America/New_York';

SELECT
    ts,
    tstz
FROM
    timestamp_demo;

SELECT
    NOW();

SELECT
    CURRENT_TIMESTAMP;

SELECT
    TIMEOFDAY();

SHOW TIMEZONE;

-- Date
DROP TABLE IF EXISTS documents;

CREATE TABLE documents (
    document_id serial PRIMARY KEY,
    header_text VARCHAR (255) NOT NULL,
    posting_date DATE NOT NULL DEFAULT CURRENT_DATE
);

INSERT INTO
    documents (header_text)
VALUES
('Billing to customer XYZ');

SELECT
    *
FROM
    documents;

DROP TABLE IF EXISTS employees;

CREATE TABLE employees (
    employee_id serial PRIMARY KEY,
    first_name VARCHAR (255),
    last_name VARCHAR (355),
    birth_date DATE NOT NULL,
    hire_date DATE NOT NULL
);

INSERT INTO
    employees (first_name, last_name, birth_date, hire_date)
VALUES
    ('Shannon', 'Freeman', '1980-01-01', '2005-01-01'),
    ('Sheila', 'Wells', '1978-02-05', '2003-01-01'),
    ('Ethel', 'Webb', '1975-01-01', '2001-01-01');

SELECT
    NOW() :: date;

SELECT
    CURRENT_DATE;

SELECT
    TO_CHAR(NOW() :: DATE, 'dd/mm/yyyy');

SELECT
    TO_CHAR(NOW() :: DATE, 'Mon dd, yyyy');

SELECT
    first_name,
    last_name,
    now() - hire_date as diff
FROM
    employees;

SELECT
    employee_id,
    first_name,
    last_name,
    AGE(birth_date)
FROM
    employees;

SELECT
    employee_id,
    first_name,
    last_name,
    EXTRACT (
        YEAR
        FROM
            birth_date
    ) AS YEAR,
    EXTRACT (
        MONTH
        FROM
            birth_date
    ) AS MONTH,
    EXTRACT (
        DAY
        FROM
            birth_date
    ) AS DAY
FROM
    employees;

-- string
CREATE TABLE character_tests (
    id serial PRIMARY KEY,
    x CHAR (1),
    y VARCHAR (10),
    z TEXT
);

INSERT INTO
    character_tests (x, y, z)
VALUES
(
        'Yes',
        'This is a test for varchar',
        'This is a very long text for the PostgreSQL text column'
    );

-- Boolean
CREATE TABLE stock_availability (
    product_id INT PRIMARY KEY,
    available BOOLEAN NOT NULL
);

INSERT INTO
    stock_availability (product_id, available)
VALUES
    (100, TRUE),
    (200, FALSE),
    (300, 't'),
    (400, '1'),
    (500, 'y'),
    (600, 'yes'),
    (700, 'no'),
    (800, '0');