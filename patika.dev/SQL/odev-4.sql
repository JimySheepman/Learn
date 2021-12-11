SELECT DISTINCT replacement_cost FROM film;
SELECT COUNT( DISTINCT replacement_cost ) FROM film;
SELECT COUNT(*) FROM film WHERE title = 'T%' AND rating = 'G';
SELECT COUNT(*) FROM film WHERE LENGTH(country) = 5;
SELECT COUNT(*) FROM City WHERE ILIKE city = '%R';
