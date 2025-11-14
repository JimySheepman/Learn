# Learn

This repository is where I share what I've tried and what I'm trying to learn.

## Docker

### Docker run command for systems

```Bash
# PostgreSQL
docker run --name postgres -e POSTGRES_PASSWORD=root -d -p 5432:5432 postgres
# Nginx
docker run --name nginx -d -p 8080:80 nginx
# Mongo
docker run --name mongodb -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=root -d -p 27017:27017 mongo
# zookeeper
docker run --name zookeeper -p 2181:2181 zookeeper
# kafka
docker run --name kafka -p 9092:9092 -e KAFKA_ZOOKEEPER_CONNECT=192.168.1.2:2181 -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://192.168.1.2:9092 -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 confluentinc/cp-kafka
```

## Docker Compose

### Redis docker-compose

```Bash
version: '3.9'
services:
  redis:
    container_name: redis
    image: redislabs/redismod:latest
    restart: always
    hostname: 'redis'
    ports:
      - 6379:6379

  redisinsight:
    container_name: redisinsight
    image: redislabs/redisinsight:latest
    restart: always
    ports:
      - 8001:8001
    volumes:
      - '~/.docker-conf/redisinsight:/db'
```

### RabbitMQ docker-compose

```Bash
version: "3.9"
services:
    rabbitmq:
        hostname: 'dev-rabbitmq'
        image: rabbitmq:3.8.14-management
        container_name: 'project-rabbitmq'
        restart: always
        ports:
            - 5672:5672
            - 15672:15672
        volumes:
            - '~/.docker-conf/rabbitmq/data/:/var/lib/rabbitmq/mnesia/'
        networks:
            - test-network

networks:
    test-network:
        driver: bridge
```
