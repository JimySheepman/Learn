/*
 * INTERSECT iki veya daha fazla SELECT sorgusunu tek bir
 * sonuç kümesinde birleştir.(A n B) Veri türleri uyumlu 
 * olmalıdır. Cümlerlerdeki sütun sayısı ve sırasıda aynı
 * olmalıdır.
*/

SELECT *
FROM most_popular_films 
INTERSECT
SELECT *
FROM top_rated_films;