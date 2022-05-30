/*
 * LIMIT satır sayısını sınırlayan ifadedir. 
 * row_count -> satır sayısı
 * row_to_skip -> satırları atlamaya yarar
*/
--- ilk 5 film getirdi
SELECT
    film_id,
    title,
    release_year
FROM
    film
ORDER BY
    film_id
LIMIT 5;

--- ilk 3 film atladı ve ardındaki 4 taneyi getirdi
SELECT
    film_id,
    title,
    release_year
FROM
    film
ORDER BY
    film_id
LIMIT 4 
OFFSET 3;

