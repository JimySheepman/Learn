/*
 * UNION iki veya daha fazla SELECT sorgusunu tek bir
 * sonuç kümesinde birleştir.(A U B) Veri türleri 
 * uyumlu olmalıdır. Birleştirilen veri kümesinde tüm
 * yenilenen satırları kaldırır. Bu satırları korumak
 * için UNION ALL kullanılır.
*/

DROP TABLE IF EXISTS top_rated_films;
CREATE TABLE top_rated_films(
	title VARCHAR NOT NULL,
	release_year SMALLINT
);

DROP TABLE IF EXISTS most_popular_films;
CREATE TABLE most_popular_films(
	title VARCHAR NOT NULL,
	release_year SMALLINT
);

INSERT INTO 
    top_rated_films(title,release_year)
VALUES
    ('The Shawshank Redemption',1994),
    ('The Godfather',1972),
    ('12 Angry Men',1957);

INSERT INTO 
    most_popular_films(title,release_year)
VALUES
    ('An American Pickle',2020),
    ('The Godfather',1972),
    ('Greyhound',2020);

SELECT * FROM top_rated_films;

SELECT * FROM most_popular_films;



SELECT * FROM top_rated_films
UNION
SELECT * FROM most_popular_films;


SELECT * FROM top_rated_films
UNION ALL
SELECT * FROM most_popular_films
ORDER BY title;