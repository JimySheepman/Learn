/*
 * INSERT ifadesi, bir tabloya yeni bir satır eklemenizi sağlar.
 * RETURNING Yalnızca eklenen kimlik listesini döndürmek 
 * istiyorsanız , yan tümcedeki idsütunu şu şekilde 
 * belirtebilirsiniz.
*/

DROP TABLE IF EXISTS links;

CREATE TABLE links (
	id SERIAL PRIMARY KEY,
	url VARCHAR(255) NOT NULL,
	name VARCHAR(255) NOT NULL,
	description VARCHAR (255),
        last_update DATE
);

INSERT INTO links (url, name)
VALUES('https://www.postgresqltutorial.com','PostgreSQL Tutorial');

INSERT INTO links (url, name)
VALUES('http://www.oreilly.com','O''Reilly Media');

INSERT INTO links (url, name, last_update)
VALUES('https://www.google.com','Google','2013-06-01');

INSERT INTO links (url, name)
VALUES('http://www.postgresql.org','PostgreSQL') 
RETURNING id;

--- multiple rows
INSERT INTO 
    links (url, name)
VALUES
    ('https://www.google.com','Google'),
    ('https://www.yahoo.com','Yahoo'),
    ('https://www.bing.com','Bing');

INSERT INTO 
    links(url,name, description)
VALUES
    ('https://duckduckgo.com/','DuckDuckGo','Privacy & Simplified Search Engine'),
    ('https://swisscows.com/','Swisscows','Privacy safe WEB-search')
RETURNING *;

INSERT INTO 
    links(url,name, description)
VALUES
    ('https://www.searchencrypt.com/','SearchEncrypt','Search Encrypt'),
    ('https://www.startpage.com/','Startpage','The world''s most private search engine')
RETURNING id;
