from elasticsearch import Elasticsearch

es = Elasticsearch()
res = es.search(index="planets", body={"query": {"match_all": {}}})
print("Got %d Hits:" % res['hits']['total'])

for hit in res['hits']['hits']:
    print("%(Name)s %(Mass)s: %(Radius)s" % hit["_source"])
