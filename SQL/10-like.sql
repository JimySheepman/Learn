/*
 * LIKE gibi anlamında kullanılır. Benzetmeler yapmaya yarar
 * '%'-> dizi döner ve '_' -> karakter döner  kullanılır
*/

--- 'Jen%' Jen ile başlayan tüm customerler
SELECT
    first_name,
    last_name
FROM
    customer
WHERE
    first_name LIKE 'Jen%';

--- foo kelimesinde arama yapma örnekleri
SELECT
	'foo' LIKE 'foo', -- true
	'foo' LIKE 'f%', -- true
	'foo' LIKE '_o_', -- true
	'bar' LIKE 'b_'; -- false

--- içerisinde 'er' geçenleri bulma
SELECT
    first_name,
    last_name
FROM
    customer
WHERE
    first_name LIKE '%er%'
ORDER BY
    first_name;

--- herhangi bir tek karakter ile başlayan
--- içinde 'her' geçenleri bulma
SELECT
    first_name,
    last_name
FROM
    customer
WHERE
    first_name LIKE '_her%'
ORDER BY
    first_name;