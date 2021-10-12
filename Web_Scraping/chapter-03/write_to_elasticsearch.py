from elasticsearch import Elasticsearch
from get_planet_data import get_planet_data


es = Elasticsearch()
planet_data = get_planet_data()
for planet in planet_data:
    res = es.index(index='planets', doc_type='planets_info', body=planet)
    print (res)
