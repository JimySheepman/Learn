import mysql.connector
import get_planet_data
from mysql.connector import errorcode
from get_planet_data import get_planet_data

try:
    cnx = mysql.connector.connect(user='root',password='mypassword',database="scraping")
    host="127.0.0.1",insert_sql = ("INSERT INTO Planets (Name, Mass, Radius,Description) " + "VALUES (%(Name)s, %(Mass)s, %(Radius)s,%(Description)s)")

    planet_data = get_planet_data()

    cursor = cnx.cursor()
    for planet in planet_data:
        print("Storing data for %s" % (planet["Name"]))
        cursor.execute(insert_sql, planet)

    cnx.commit()

    cursor.close()
    cnx.close()

except mysql.connector.Error as err:
    if err.errno == errorcode.ER_ACCESS_DENIED_ERROR:
        print("Something is wrong with your user name or password")
    elif err.errno == errorcode.ER_BAD_DB_ERROR:
        print("Database does not exist")
    else:
        print(err)
else:
    cnx.close()
