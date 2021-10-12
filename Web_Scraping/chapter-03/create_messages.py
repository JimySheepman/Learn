from urllib.request import urlopen
from bs4 import BeautifulSoup

import boto3
import botocore

access_key="AKIAIXFTCYO7FEL55TCQ"
access_secret_key="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"

sqs = boto3.client('sqs', "us-west-2",aws_access_key_id = access_key,aws_secret_access_key = access_secret_key)

queue = sqs.create_queue(QueueName="PlanetMoreInfo")
print (queue)

html = urlopen("http://127.0.0.1:8081/pages/planets.html")
bsobj = BeautifulSoup(html, "lxml")
planets = []
planet_rows = bsobj.html.body.div.table.findAll("tr", {"class": "planet"})

for i in planet_rows:
    tds = i.findAll("td")
    more_info_url = tds[5].findAll("a")[0]["href"].strip()
    sqs.send_message(QueueUrl=queue["QueueUrl"],
    MessageBody=more_info_url)
    print("Sent %s to %s" % (more_info_url, queue["QueueUrl"]))
