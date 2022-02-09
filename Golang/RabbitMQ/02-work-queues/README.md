# Work queues

Distributing tasks among workers (the competing consumers pattern)

```Bash
# shell 1
$go run new_task.go 1
2022/02/09 20:41:13  [x] Sent 1
$go run new_task.go 2
2022/02/09 20:41:14  [x] Sent 2
$go run new_task.go 3 
2022/02/09 20:41:15  [x] Sent 3

```

```Bash
# shell 2
$go run worker.go
2022/02/09 20:41:07  [*] Waiting for messages. To exit press CTRL+C
2022/02/09 20:41:13 Received a message: 1
2022/02/09 20:41:13 Done
```

```Bash
# shell 3
$go run worker.go 
2022/02/09 20:41:07  [*] Waiting for messages. To exit press CTRL+C
2022/02/09 20:41:14 Received a message: 2
2022/02/09 20:41:14 Done
```

```Bash
# shell 4
$go run worker.go
2022/02/09 20:41:07  [*] Waiting for messages. To exit press CTRL+C
2022/02/09 20:41:15 Received a message: 3
2022/02/09 20:41:15 Done
```
