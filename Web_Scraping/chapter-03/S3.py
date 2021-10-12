import requests
import boto3

data = requests.get("http://localhost:8080/planets.html").text

s3 = boto3.client('s3')
bucket_name = "planets-content"
s3.create_bucket(Bucket=bucket_name, ACL='public-read')
s3.put_object(Bucket=bucket_name, Key='planet.html',
Body=data, ACL="public-read")
