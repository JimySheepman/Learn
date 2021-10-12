import boto3
import botocore
import requests
from bs4 import BeautifulSoup

print("Starting")

sqs = boto3.client('sqs', "us-west-2")

print("Created client")

queue = sqs.create_queue(QueueName="PlanetMoreInfo")
queue_url = queue["QueueUrl"]
print ("Opened queue: %s" % queue_url)

while True:
	print ("Attempting to receive messages")
	response = sqs.receive_message(QueueUrl=queue_url,
								   MaxNumberOfMessages=1,
								   WaitTimeSeconds=1)
	if not 'Messages' in response:
		print ("No messages")
		continue

	message = response['Messages'][0]
	receipt_handle = message['ReceiptHandle']
	url = message['Body']

	html = requests.get(url)
	bsobj = BeautifulSoup(html.text, "lxml")

	planet=bsobj.findAll("h1", {"id": "firstHeading"} )[0].text
	albedo_node = bsobj.findAll("a", {"href": "/wiki/Geometric_albedo"})[0]
	root_albedo = albedo_node.parent
	albedo = root_albedo.text.strip()

	sqs.delete_message(
		QueueUrl=queue_url,
		ReceiptHandle=receipt_handle
	)

	print("%s: %s" % (planet, albedo))