/*
 * BETWEEN aralık belirler
*/

SELECT
    customer_id,
    payment_id,
    amount
FROM
    payment
WHERE
    amount BETWEEN 8 AND 9;

SELECT
    customer_id,
    payment_id,
    amount,
    payment_date

FROM
    payment
WHERE
    payment_date BETWEEN '2007-02-07' AND '2007-02-15';