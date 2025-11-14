SELECT COUNT(*)  FROM film WHERE length > ( SELECT AVG(length) FROM film);
SELECT COUNT(*) FROM film WHERE rental_rate = ( SELECT MAX(rental_rate) FROM film);

SELECT title
FROM film
WHERE film_id = any
( SELECT film_id
 FROM film
WHERE rental_rate = ( SELECT MIN(rental_rate ) FROM film )
 AND
 replacement_cost = ( SELECT MIN(replacement_cost) FROM film ) );

 SELECT first_name,last_name
FROM customer c
JOIN payment p
ON ( p.customer_id = c.customer_id )
WHERE amount = ( SELECT MAX(amount) from payment );
