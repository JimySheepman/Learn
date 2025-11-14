/*
 * CASE -> ifadesi, diğer programlama dillerindeki IF/ELSE ifadesi ile aynıdır.
 * COALESCE ->  Boş değerleri etkili bir şekilde işlemek için SELECT deyiminde 
 *              bu işlevi nasıl uygulayacağınızı öğreneceksiniz.
 * ISNULL -> İfade NULL ise, ISNULL işlevi değiştirmeyi döndürür. Aksi takdirde,
 *           ifadenin sonucunu döndürür.
 * NULLIF -> işlevi, PostgreSQL tarafından sağlanan en yaygın koşullu ifadelerden
 *           biridir. 
 * CAST -> Bir veri türünün değerini dönüştürmek için kullanılır.
 */
SELECT
    title,
    length,
    CASE
        WHEN length > 0
        AND length <= 50 THEN 'Short'
        WHEN length > 50
        AND length <= 120 THEN 'Medium'
        WHEN length > 120 THEN 'Long'
    END duration
FROM
    film
ORDER BY
    title;

SELECT
    SUM (
        CASE
            WHEN rental_rate = 0.99 THEN 1
            ELSE 0
        END
    ) AS "Economy",
    SUM (
        CASE
            WHEN rental_rate = 2.99 THEN 1
            ELSE 0
        END
    ) AS "Mass",
    SUM (
        CASE
            WHEN rental_rate = 4.99 THEN 1
            ELSE 0
        END
    ) AS "Premium"
FROM
    film;

SELECT
    product,
    (price - COALESCE(discount, 0)) AS net_price
FROM
    items;

SELECT
    CASE
        WHEN expression IS NULL THEN replacement
        ELSE expression
    END AS column_alias;

CREATE TABLE posts (
    id serial primary key,
    title VARCHAR (255) NOT NULL,
    excerpt VARCHAR (150),
    body TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

INSERT INTO
    posts (title, excerpt, body)
VALUES
    (
        'test post 1',
        'test post excerpt 1',
        'test post body 1'
    ),
    ('test post 2', '', 'test post body 2'),
    ('test post 3', null, 'test post body 3');

SELECT
    ID,
    title,
    excerpt
FROM
    posts;

SELECT
    id,
    title,
    COALESCE (excerpt, LEFT(body, 40))
FROM
    posts;

SELECT
    id,
    title,
    COALESCE (NULLIF (excerpt, ''), LEFT (body, 40))
FROM
    posts;

SELECT
	CAST ('10.2' AS DOUBLE);