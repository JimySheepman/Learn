SELECT city,country FROM city ci JOIN country co ON (co.country_id = ci.country_id );
SELECT first_name,last_name,payment_id FROM customer c JOIN payment p ON ( p.customer_id = c.customer_id);
SELECT first_name,last_name,rental_id FROM customer c JOIN rental r ON( c.customer_id = r.customer_id );
