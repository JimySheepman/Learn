SELECT country From Country WHERE country like 'A%a';
SELECT country From Country WHERE LENGTH(country) > 6 AND country like '%n';
SELECT title FROM film WHERE title ILIKE '%t%t%t%t%';
SELECT * From film Where title LIKE 'C%' AND length > 90 AND rental_rate = 2.99 ;
