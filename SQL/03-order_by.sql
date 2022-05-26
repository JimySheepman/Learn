/*
 * artan azalan şekilde sıralmamızı sağlar. 
 * ASC artan düzende sıralar. Varsayılan değeri budur.
 * DESC azalan düzende sıralar.
*/

SELECT
    first_name,
    last_name
FROM
    customer
ORDER BY
    first_name ASC;

SELECT
    first_name,
    last_name
FROM
    customer
ORDER BY
    first_name DESC;

/*
 * Aşağıdaki örnekte ilk artan sonra ise azalan şekilde sıralar. 
 * ayna ada sahip insanların soy adları azalan sırada sıralanmıştır.
*/

SELECT
    first_name,
    last_name
FROM
    customer
ORDER BY
    first_name ASC,
    last_name DESC;

/*
 * LENGTH -> işlevi bir dizi alır ve onun uzunluğunu döner.
 * Aşağıdaki ifade ilk isimleri ve uzunluklarını seçer. 
 * Satırları ilk isimlerin uzunluklarına göre sıralar
*/

SELECT
    first_name,
    LENGTH(first_name) len
FROM
    customer
ORDER BY
    len DESC;

/*
 * NULLS FIRST, NULLS LAST -> null değerlerin nerede 
 * sıralanacağını belirler.
*/

SELECT num
FROM sort_demo
ORDER BY num NULLS FIRST;

SELECT num
FROM sort_demo
ORDER BY num NULLS LAST;