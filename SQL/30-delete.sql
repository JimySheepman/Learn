/*
 * DELETE ifadesi, bir tablodan bir veya daha fazla satırı 
 * silmenizi sağlar.
*/

DROP TABLE IF EXISTS links;

CREATE TABLE links (
   id serial PRIMARY KEY,
   url varchar(255) NOT NULL,
   name varchar(255) NOT NULL,
   description varchar(255),
   rel varchar(10),
   last_update date DEFAULT now()
);

INSERT INTO  
   links 
VALUES 
   ('1', 'https://www.postgresqltutorial.com', 'PostgreSQL Tutorial', 'Learn PostgreSQL fast and easy', 'follow', '2013-06-02'),
   ('2', 'http://www.oreilly.com', 'O''Reilly Media', 'O''Reilly Media', 'nofollow', '2013-06-02'),
   ('3', 'http://www.google.com', 'Google', 'Google', 'nofollow', '2013-06-02'),
   ('4', 'http://www.yahoo.com', 'Yahoo', 'Yahoo', 'nofollow', '2013-06-02'),
   ('5', 'http://www.bing.com', 'Bing', 'Bing', 'nofollow', '2013-06-02'),
   ('6', 'http://www.facebook.com', 'Facebook', 'Facebook', 'nofollow', '2013-06-01'),
   ('7', 'https://www.tumblr.com/', 'Tumblr', 'Tumblr', 'nofollow', '2013-06-02'),
   ('8', 'http://www.postgresql.org', 'PostgreSQL', 'PostgreSQL', 'nofollow', '2013-06-02');

SELECT * FROM links;

DELETE FROM links
WHERE id = 8;

DELETE FROM links
WHERE id = 10;

DELETE FROM links
WHERE id = 7
RETURNING *;

DELETE FROM links;

DROP TABLE IF EXISTS contacts;
CREATE TABLE contacts(
   contact_id serial PRIMARY KEY,
   first_name varchar(50) NOT NULL,
   last_name varchar(50) NOT NULL,
   phone varchar(15) NOT NULL
);


DROP TABLE IF EXISTS blacklist;
CREATE TABLE blacklist(
   phone varchar(15) PRIMARY KEY
);


INSERT INTO contacts(first_name, last_name, phone)
VALUES ('John','Doe','(408)-523-9874'),
      ('Jane','Doe','(408)-511-9876'),
      ('Lily','Bush','(408)-124-9221');


INSERT INTO blacklist(phone)
VALUES ('(408)-523-9874'),
      ('(408)-511-9876');

DELETE FROM contacts 
USING blacklist
WHERE contacts.phone = blacklist.phone;

SELECT * FROM contacts;

DELETE FROM contacts
WHERE phone IN (SELECT phone FROM blacklist);