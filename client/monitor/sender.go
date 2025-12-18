// monitor/sender.go
package monitor

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type SenderConfig struct {
	URL     string        // ex: http://localhost:8080/metrics
	Timeout time.Duration // ex: 3 * time.Second
}

var senderConfig SenderConfig

func SetSender(config SenderConfig) {
	senderConfig = config
}

func Send(snapshot Snapshot) error {
	payload, err := json.Marshal(snapshot)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", senderConfig.URL, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: senderConfig.Timeout,
	}
	_, err = client.Do(req)
	return err
}
