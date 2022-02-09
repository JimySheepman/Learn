# Topics

Receiving messages based on a pattern (topics)

```Bash
# If you want to save only 'warning' and 'error' (and not 'info') log messages to a file, just open a console and type:
$go run receive_logs_direct.go warning error > logs_from_rabbit.log
# If you'd like to see all the log messages on your screen, open a new terminal and do:
$go run receive_logs_direct.go info warning error
[*] Waiting for logs. To exit press CTRL+C
$go run emit_log_direct.go error "Run. Run. Or it will explode."
[x] Sent 'error':'Run. Run. Or it will explode.'

```
