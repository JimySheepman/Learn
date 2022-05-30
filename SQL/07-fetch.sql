/*
 * FETCH bir sorgu tarafından döndürülen sonuçların 
 * bir kısmını almada kullalanır.
*/
SELECT
    film_id,
    title
FROM
    film
ORDER BY
    title 
FETCH FIRST ROW ONLY;