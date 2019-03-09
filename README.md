# Gollector

### Build and Run

```
$ go build
$ ./gollector --help
  Usage of ./gollector:
    -amqpUrl string
      AQMP URI format (default "amqp://guest:guest@localhost:5672/")
    -collection string
      Database Collection name (default "gollector")
    -database string
      Database name (default "gollector")
    -mongoUrl string
      Mongo URL (default "mongodb://localhost:27017/")
    -production
      When production is activated, the log level is set to Warn
    -queue string
      AMQP Queue Name (default "gollector")
```
