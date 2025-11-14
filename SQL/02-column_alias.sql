/*
 * sutüna verilen takma adtır. 
 * AS ile belirtilebilir yada direkt yanına yazılır
*/

SELECT
    first_name AS fname,
    last_name lname,
FROM customer;

/*
 * aşağıdaki sorgu isim soy isimi tamamını araya boşluk
 * koyarak birleşik bir şekilde döner tam isiim olarak 
 * yanıtlar.
*/
SELECT
    first_name || ' ' || last_name  AS full_name
FROM
    customer;


/*
 * boşluk içermesini istiyorsak "" işaretini kullanarak
 * sutün içine birden çok boşluklu ifade yazılabilir.
*/

SELECT
    first_name || ' ' || last_name  AS "full name"
FROM
    customer;