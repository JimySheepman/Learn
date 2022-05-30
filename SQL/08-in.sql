/*
 * IN içerip içermediğini kontrol etmemize sağlar
 * içerisine sorgu alabilir
*/

SELECT
    cutomer_id,
    rental_id,
    return_date
FROM
    rental
WHERE
    cutomer_id IN(1,2)
ORDER BY
    return_date DESC;


SELECT
    cutomer_id,
    first_name,
    last_name
FROM
    customer
WHERE
    cutomer_id IN(
        SELECT cutomer_id
        FROM rental
        WHERE CAST (return_date as DATE) = '2005-05-27'
    )
ORDER BY
    cutomer_id;