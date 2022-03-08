# Topics

Receiving messages based on a pattern (topics)

```Bash
# To receive all the logs:
$go run receive_logs_topic.go "#"
# To receive all logs from the facility "kern":
$go run receive_logs_topic.go "kern.*"
# Or if you want to hear only about "critical" logs:
$go run receive_logs_topic.go "*.critical"
# You can create multiple bindings:
$go run receive_logs_topic.go "kern.*" "*.critical"
# And to emit a log with a routing key "kern.critical" type:
$go run emit_log_topic.go "kern.critical" "A critical kernel error"
```
