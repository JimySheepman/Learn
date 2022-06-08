/*
 * CUBE birden çok  grouping sets oluşturmaya
 * yarar. Gruplandırma kümesi, gruplamak 
 * istediğiniz bir sütun kümesidir.
*/

SELECT
    brand,
    segment,
    SUM (quantity)
FROM
    sales
GROUP BY
    CUBE (brand, segment)
ORDER BY
    brand,
    segment;


SELECT
    brand,
    segment,
    SUM (quantity)
FROM
    sales
GROUP BY
    brand,
    CUBE (segment)
ORDER BY
    brand,
    segment;
