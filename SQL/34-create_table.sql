/*
 * NOT NULL –> bir sütundaki değerlerin olmamasını sağlar NULL.
 * UNIQUE  -> aynı tablodaki satırlar arasında benzersiz bir sütundaki değerleri sağlar.
 * PRIMARY KEY -> bir tablodaki satırları benzersiz şekilde tanımlayan birincil 
 *                anahtar sütunu. Bir tablonun bir ve yalnızca bir birincil anahtarı 
 *                olabilir. Birincil anahtar kısıtlaması, bir tablonun birincil 
 *                anahtarını tanımlamanıza olanak tanır.
 * CHECK -> bir CHECKkısıtlama, verilerin bir boolean ifadesini karşılamasını sağlar.
 * FOREIGN KEY -> bir tablodaki bir sütundaki veya bir sütun grubundaki değerlerin başka 
 *                bir tablodaki bir sütunda veya sütun grubunda var olmasını sağlar. 
 *                Birincil anahtarın aksine, bir tablonun birçok yabancı anahtarı olabilir.
 * UNLOGGED -> sözcük, yeni tablonun günlüğe kaydedilmemiş bir tablo olarak oluşturulmasına
 *             izin verir.
 * IF NOT EXISTS -> eğer yoksa demek.
 * CREATE SEQUENCE -> dizileri ve bir dizi sayı oluşturmak için bir dizi nesnesinin nasıl 
 *                    kullanılacağı hakkında bilgi edineceksiniz. Tanım olarak, bir dizi 
 *                    sıralı bir tamsayı listesidir. Sıralamadaki sayıların sırası önemlidir. 
 *                    Örneğin {1,2,3,4,5}ve {5,4,3,2,1}tamamen farklı dizilerdir. Bir dizi, 
 *                    belirli bir belirtime dayalı olarak bir tamsayı dizisi oluşturan, 
 *                    kullanıcı tanımlı şemaya bağlı bir nesnedir.
 *  GENERATED AS IDENTITY-> kimlik sütunu oluşturmak için kullanılır.
 *  ALTER TABLE -> bir tablo yapısını değiştirmek için kullanılır. Actionlar;
 *                 Add a column
 *                 Drop a column
 *                 Change the data type of a column
 *                 Rename a column
 *                 Set a default value for the column.
 *                 Add a constraint to a column.
 *                 Rename a table
 * TRUNCATE TABLE -> büyük tablolardan tüm verileri hızlı bir şekilde silmek için kullanılır.
 * temporary table-> Geçici bir tablo, adından da anlaşılacağı gibi, bir veritabanı oturumu 
 *                   süresince var olan kısa ömürlü bir tablodur. PostgreSQL , bir oturumun
 *                   veya bir işlemin sonunda geçici tabloları otomatik olarak bırakır .
*/

CREATE UNLOGGED TABLE new_table_name
AS query;

CREATE TABLE IF NOT EXISTS new_table_name
AS query;


CREATE SEQUENCE [ IF NOT EXISTS ] sequence_name
    [ AS { SMALLINT | INT | BIGINT } ]
    [ INCREMENT [ BY ] increment ]
    [ MINVALUE minvalue | NO MINVALUE ] 
    [ MAXVALUE maxvalue | NO MAXVALUE ]
    [ START [ WITH ] start ] 
    [ CACHE cache ] 
    [ [ NO ] CYCLE ]
    [ OWNED BY { table_name.column_name | NONE } ]

CREATE TABLE color (
    color_id INT GENERATED ALWAYS AS IDENTITY,
    color_name VARCHAR NOT NULL
);

-- ALTER TABLE
ALTER TABLE table_name 
ADD COLUMN column_name datatype column_constraint;

ALTER TABLE table_name 
DROP COLUMN column_name;

ALTER TABLE table_name 
RENAME COLUMN column_name 
TO new_column_name;

ALTER TABLE table_name 
ALTER COLUMN column_name 
[SET DEFAULT value | DROP DEFAULT];

ALTER TABLE table_name 
ALTER COLUMN column_name 
[SET NOT NULL| DROP NOT NULL];

ALTER TABLE table_name 
ALTER COLUMN column_name 
[SET NOT NULL| DROP NOT NULL];

ALTER TABLE table_name 
ADD CHECK expression;

ALTER TABLE table_name 
ADD CONSTRAINT constraint_name constraint_definition;

ALTER TABLE table_name 
RENAME TO new_table_name;


TRUNCATE TABLE invoices;

TRUNCATE TABLE table_name 
RESTART IDENTITY;

TRUNCATE TABLE table_name 
CASCADE;

CREATE TEMPORARY TABLE temp_table_name(
    column_list
);

CREATE TEMP TABLE temp_table(
    column_list
);


-- tablo kopyalamaya yarar
CREATE TABLE new_table AS 
TABLE existing_table;

CREATE TABLE new_table AS 
TABLE existing_table 
WITH NO DATA;

-- CREATE DOMAIN statement
CREATE DOMAIN contact_name AS 
   VARCHAR NOT NULL CHECK (value !~ '\s');