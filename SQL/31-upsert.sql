/*
 * İlişkisel veritabanlarında, upsert terimi birleştirme olarak
 * adlandırılır. Buradaki fikir, tabloya yeni bir satır eklediğinizde
 * PostgreSQL satırı zaten varsa günceller , aksi takdirde yeni satırı 
 * ekler. Bu yüzden eyleme upsert (güncelleme veya ekleme kombinasyonu) 
 * diyoruz. PostgreSQL'de upsert özelliğini kullanmak için INSERT ON CONFLICT
 * kullanıyoruz.
*/

DROP TABLE IF EXISTS customers;

CREATE TABLE customers (
	customer_id serial PRIMARY KEY,
	name VARCHAR UNIQUE,
	email VARCHAR NOT NULL,
	active bool NOT NULL DEFAULT TRUE
);

INSERT INTO 
    customers (name, email)
VALUES 
    ('IBM', 'contact@ibm.com'),
    ('Microsoft', 'contact@microsoft.com'),
    ('Intel', 'contact@intel.com');

INSERT INTO customers (NAME, email)
VALUES('Microsoft','hotline@microsoft.com') 
ON CONFLICT ON CONSTRAINT customers_name_key 
DO NOTHING;

INSERT INTO customers (name, email)
VALUES('Microsoft','hotline@microsoft.com') 
ON CONFLICT (name) 
DO 
    UPDATE SET email = EXCLUDED.email || ';' || customers.email;