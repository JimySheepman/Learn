/*
 * EXISTS bir alt sorgudaki satırların varlığını 
 * test eden bir boole operatörüdür.
*/

SELECT 
    column1
FROM 
    table_1
WHERE 
    EXISTS( SELECT 
                1 
            FROM 
                table_2 
            WHERE 
                column_2 = table_1.column_1);


SELECT first_name,
    last_name
FROM customer c
WHERE EXISTS
    (SELECT 1
    FROM payment p
    WHERE p.customer_id = c.customer_id
    AND amount > 11 )
ORDER BY first_name,
        last_name;

SELECT first_name,
    last_name
FROM customer c
WHERE NOT EXISTS
    (SELECT 1
    FROM payment p
    WHERE p.customer_id = c.customer_id
    AND amount > 11 )
ORDER BY first_name,
        last_name;


SELECT
	first_name,
	last_name
FROM
	customer
WHERE
	EXISTS( SELECT NULL )
ORDER BY
	first_name,
	last_name;
