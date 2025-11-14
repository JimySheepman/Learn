SELECT city,country FROM city ci LEFT JOIN country co ON ( ci.country_id = co.country_id );
SELECT first_name,last_name,payment_id FROM customer c RIGHT JOIN payment p ON ( p.customer_id = c.customer_id);
SELECT first_name,last_name,rental_id FROM customer c FULL JOIN rental r ON( c.customer_id = r.customer_id );
