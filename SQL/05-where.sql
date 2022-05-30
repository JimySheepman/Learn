/*
 * WHERE koşul ifadesi sağlar ve filtreleme yapmamıza 
 * olanak saplar.
*/

SELECT 
    last_name,
    first_name
FROM
    customer
WHERE
    first_name = 'jamie';


/*
 * WHERE koşul ifadesi kullanılan logic operatorler; 
 *  '=' eşitlik anlamı katar
 *  '>' daha büyüktür
 *  '<' daha küçüktür
 *  '>=' eşit yada daha büyüktür
 *  '<=' eşit yada daha küçüktür
 *  '<>' or '!=' eşit değil
 *  'AND' mantıksal ve
 *  'OR' mantıksal veya
 *  'IN' listedeki herkangi bir değerle eşleşirse TRUE döner
 *  'BETWEEN' verilen araklıkta değer varsa TRUE döner
 *  'LIKE' verilen pettern ile uyuşur ise TRUE döner
 *  'IS NULL' bir değer NULL ise TRUE döner
 *  'NOT' tüm operatörlerin tersini almak için kullanılır
*/
--- AND
SELECT 
    last_name,
    first_name
FROM
    customer
WHERE
    first_name = 'jamie' AND
    last_name = 'Rice';
--- OR
SELECT 
    last_name,
    first_name
FROM
    customer
WHERE
    first_name = 'jamie' OR
    last_name = 'Rice';
--- IN
SELECT 
    last_name,
    first_name
FROM
    customer
WHERE
    first_name IN ('Ann','Anne','Annie');
--- LIKE
SELECT 
    last_name,
    first_name
FROM
    customer
WHERE
    first_name LIKE 'Ann%'
--- BETWEEN
SELECT
	first_name,
	LENGTH(first_name) name_length
FROM
	customer
WHERE 
	first_name LIKE 'A%' AND
	LENGTH(first_name) BETWEEN 3 AND 5
ORDER BY
	name_length;
--- <>
SELECT 
    last_name,
    first_name
FROM
    customer
WHERE
    first_name LIKE 'Bra%' AND
    last_name <> 'Motley';