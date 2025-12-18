package monitor

import (
	"os"
	"runtime"
	"time"
)

type Snapshot struct {
	AppName     string           `json:"app_name"`
	Environment string           `json:"environment"`
	Hostname    string           `json:"hostname"`
	Timestamp   int64            `json:"timestamp"`
	Goroutines  int              `json:"goroutines"`
	MemStats    runtime.MemStats `json:"memstats"`
	Custom      map[string]int64 `json:"custom"`
}

type Config struct {
	Appname     string
	Environment string
}

var config Config

func Init(cfg Config) {
	config = cfg
}

func Capture(custom map[string]int64) Snapshot {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	hostname, _ := os.Hostname()

	return Snapshot{
		AppName:     config.Appname,
		Environment: config.Environment,
		Hostname:    hostname,
		Timestamp:   time.Now().Unix(),
		Goroutines:  runtime.NumGoroutine(),
		MemStats:    m,
		Custom:      custom,
	}
}
