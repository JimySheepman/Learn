/*
 * EXCEPT iki veya daha fazla SELECT sorgusunu tek bir
 * sonuç kümesinde birleştir.((A' n B) || (A n B')) 
 * Veri türleri uyumlu olmalıdır. Cümlerlerdeki 
 * sütun sayısı ve sırasıda aynı olmalıdır.
*/


SELECT *
FROM most_popular_films 
EXCEPT
SELECT *
FROM top_rated_films;