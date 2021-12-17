# RabbitMQ Tutorial Go

İnstall AMQP package

```Bash
go get github.com/streadway/amqp
```

## Code

Code examples are executed via `go run`:

["Hello World!"](https://www.rabbitmq.com/tutorials/tutorial-one-go.html):

```Bash
    go run send.go
    go run receive.go
```

[Work Queues](https://www.rabbitmq.com/tutorials/tutorial-two-go.html):

```Bash
    go run new_task.go hello world
    go run worker.go
```

[Publish/Subscribe](https://www.rabbitmq.com/tutorials/tutorial-three-go.html)

```Bash
    go run receive_logs.go
    go run emit_log.go hello world
```

[Routing](https://www.rabbitmq.com/tutorials/tutorial-four-go.html)

```Bash
    go run receive_logs_direct.go info warn
    go run emit_log_direct.go warn "a warning"
```

[Topics](https://www.rabbitmq.com/tutorials/tutorial-five-go.html)

```Bash
    go run receive_logs_topic.go "kern.*" "*.critical"
    go run emit_log_topic.go kern.critical "A critical kernel error"
```

[RPC](https://www.rabbitmq.com/tutorials/tutorial-six-go.html)

```Bash
    go run rpc_server.go
    go run rpc_client.go 10
```

To learn more, see [`rabbitmq/amqp091-go`](https://github.com/rabbitmq/amqp091-go).
