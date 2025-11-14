/*
 * HAVING bir grup veya bir küme için bir arama koşulu belirtir.
 * filtrelemek için kullanılır.
*/

SELECT
	customer_id,
	SUM (amount)
FROM
	payment
GROUP BY
	customer_id
HAVING
	SUM (amount) > 200;

SELECT
	store_id,
	COUNT (customer_id)
FROM
	customer
GROUP BY
	store_id