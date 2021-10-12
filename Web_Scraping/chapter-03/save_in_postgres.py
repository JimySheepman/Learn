import psycopg2
from get_planet_data import get_planet_data

try:

    conn = psycopg2.connect("dbname='scraping', host='localhost',user='postgres', password='mypassword'")

    insert_sql = ('INSERT INTO public."Planets"(name, mass, radius,description, moreinfo) ' +'VALUES (%(Name)s, %(Mass)s, %(Radius)s, %(Description)s,%(MoreInfo)s);')

    cur = conn.cursor()
    planet_data = get_planet_data()
    for planet in planet_data:
        cur.execute(insert_sql, planet)
    conn.commit()
    cur.close()
    conn.close()
    print("Successfully wrote data to the database")
except Exception as ex:
    print(ex)
