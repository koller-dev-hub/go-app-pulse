# 📊 Go App Pulse

[![Go Reference](https://pkg.go.dev/badge/github.com/koller-dev-hub/go-app-pulse.svg)](https://pkg.go.dev/github.com/koller-dev-hub/go-app-pulse)
[![Go Report Card](https://goreportcard.com/badge/github.com/koller-dev-hub/go-app-pulse)](https://goreportcard.com/report/github.com/koller-dev-hub/go-app-pulse)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

**Go App Pulse** is a lightweight, zero-dependency monitoring library for Go applications. It captures runtime metrics including memory statistics, goroutine counts, and custom application metrics, then sends them to your monitoring endpoint.

## ✨ Features

- 🚀 **Lightweight** - Minimal overhead on your application
- 📈 **Runtime Metrics** - Captures memory stats, goroutines, and more
- 🎯 **Custom Metrics** - Add your own application-specific metrics
- 🔧 **Easy Integration** - Simple API, quick setup
- 🌐 **HTTP Transport** - Sends metrics via HTTP POST to any endpoint

## 📦 Installation

```bash
go get github.com/koller-dev-hub/go-app-pulse
```

## 🚀 Quick Start

```go
package main

import (
    "log"
    "time"

    "github.com/koller-dev-hub/go-app-pulse/client/monitor"
)

func main() {
    // Initialize the monitor with app info
    monitor.Init(monitor.Config{
        Appname:     "my-awesome-app",
        Environment: "production",
    })

    // Configure the sender endpoint
    monitor.SetSender(monitor.SenderConfig{
        URL:     "https://your-metrics-server.com/api/metrics",
        Timeout: 5 * time.Second,
    })

    // Start collecting and sending metrics periodically
    go func() {
        ticker := time.NewTicker(10 * time.Second)
        defer ticker.Stop()

        for range ticker.C {
            // Capture metrics with optional custom data
            snapshot := monitor.Capture(map[string]int64{
                "active_users":   150,
                "requests_count": 1234,
                "db_latency_ms":  45,
            })

            // Send to your monitoring endpoint
            if err := monitor.Send(snapshot); err != nil {
                log.Printf("Failed to send metrics: %v", err)
            }
        }
    }()

    // Your application code here...
    select {}
}
```

## 📖 API Reference

### Configuration

#### `monitor.Init(cfg Config)`

Initializes the monitor with application information.

```go
monitor.Init(monitor.Config{
    Appname:     "my-app",      // Application name
    Environment: "production",  // Environment (dev, staging, production)
})
```

#### `monitor.SetSender(cfg SenderConfig)`

Configures the HTTP sender.

```go
monitor.SetSender(monitor.SenderConfig{
    URL:     "https://metrics.example.com/api/v1/metrics",
    Timeout: 5 * time.Second,
})
```

### Metrics Collection

#### `monitor.Capture(custom map[string]int64) *Snapshot`

Captures a snapshot of the current application state.

```go
snapshot := monitor.Capture(map[string]int64{
    "db_connections": 25,
    "cache_hits":     1500,
})
```

#### `monitor.Send(snapshot *Snapshot) error`

Sends the captured snapshot to the configured endpoint.

```go
if err := monitor.Send(snapshot); err != nil {
    log.Printf("Error: %v", err)
}
```

### Snapshot Structure

The `Snapshot` struct contains the following fields:

| Field         | Type               | Description                  |
| ------------- | ------------------ | ---------------------------- |
| `AppName`     | `string`           | Application name             |
| `Environment` | `string`           | Deployment environment       |
| `Hostname`    | `string`           | Host machine name            |
| `Timestamp`   | `int64`            | Unix timestamp               |
| `Goroutines`  | `int`              | Number of active goroutines  |
| `MemStats`    | `runtime.MemStats` | Go runtime memory statistics |
| `Custom`      | `map[string]int64` | Custom application metrics   |

### JSON Payload Example

```json
{
  "app_name": "my-awesome-app",
  "environment": "production",
  "hostname": "server-01",
  "timestamp": 1734537600,
  "goroutines": 42,
  "memstats": {
    "Alloc": 2642976,
    "TotalAlloc": 21226712,
    "Sys": 14506248,
    "HeapAlloc": 2642976,
    "HeapSys": 7405568,
    "HeapInuse": 4128768,
    "NumGC": 29
  },
  "custom": {
    "active_users": 150,
    "requests_count": 1234,
    "db_latency_ms": 45
  }
}
```

## 🐳 Docker Usage

When running inside Docker containers, use the appropriate network address:

```go
// Accessing host machine from container
monitor.SetSender(monitor.SenderConfig{
    URL:     "http://host.docker.internal:8080/api/metrics",
    Timeout: 5 * time.Second,
})

// Or use container name in Docker network
monitor.SetSender(monitor.SenderConfig{
    URL:     "http://metrics-server:8080/api/metrics",
    Timeout: 5 * time.Second,
})
```

## 🔧 Environment-Based Configuration

```go
package main

import (
    "os"
    "time"

    "github.com/koller-dev-hub/go-app-pulse/client/monitor"
)

func main() {
    metricsURL := os.Getenv("METRICS_URL")
    if metricsURL == "" {
        metricsURL = "http://localhost:8080/api/metrics"
    }

    monitor.Init(monitor.Config{
        Appname:     os.Getenv("APP_NAME"),
        Environment: os.Getenv("APP_ENV"),
    })

    monitor.SetSender(monitor.SenderConfig{
        URL:     metricsURL,
        Timeout: 5 * time.Second,
    })
}
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## 👤 Author

**William Koller** - [@koller-dev-hub](https://github.com/koller-dev-hub)
