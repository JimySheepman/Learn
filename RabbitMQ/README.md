# Hand-on RabbitMQ

Bu projede hem interval value ile hemde json value ile kuyruklama yapısını ne olduğunu ve nasıl çalıştığını deneyimleyerek öğrendim.

### Docker RabbitMQ Kurulumu

```Bash
docker pull rabbitmq
docker run -d --hostname my-rabbit --name some-rabbit rabbitmq:3
```

### Node.js Ayarları

```Bash
npm init -y
npm install --save amqplib
```
