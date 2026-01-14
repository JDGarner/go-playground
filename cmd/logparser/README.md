# Log Parser

## Requirements

input log files are structued like:
```
{"timestamp": "2025-10-28T10:00:00Z", "service": "auth", "level": "error", "message": "Invalid token"}
{"timestamp": "2025-10-28T10:00:01Z", "service": "payments", "level": "info", "message": "Payment completed"}
{"timestamp": "2025-10-28T10:00:02Z", "service": "auth", "level": "warning", "message": "Slow login"}
```

This command-line tool:
1. Load logs from multiple files
2. Support queries like:
   - `service=auth`
   - `level=error`
   - `service=auth level=warning`
3. Print results in chronological order


## Usage

```
go run main.go -path=/path/to/logs -query="service=auth level=error"
```

just running `go run main.go` will result in the default args being supplied (the path will be set to "logs" and no query will be given so all log lines will be printed)
