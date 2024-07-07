# gLogger

gLogger is a flexible and extensible logging library for Go applications. It supports various log levels and allows you
to subscribe services like Sentry, Logstash, or any other service as a subscriber. When you log an error or other
messages, gLogger sends the data to all subscribed services based on the log level.

## Install

```bash
go get github.com/mjedari/glogger
```

## Features

### Subscriber System:
Allows you to subscribe external services (e.g., Sentry, Logstash) to receive log data.

### Log levels:
Supports different log levels (Debug, Info, Warn, Error, Fatal).

## Contributing to gLogger

To contribute please have a look at
to [CONTRIBUTION.md](https://repo.abanicon.com/abantheter-microservices/glogger/-/blob/main/CONTRIBUTION.md).
