package splunk

import (
	"github.com/sirupsen/logrus"
)

// Hook is a logrus hook for splunk
type Hook struct {
	Client *Client
	levels []logrus.Level
}

// NewHook creates new hook
// client - splunk client instance (use NewClient)
// level - log level
func NewHook(client *Client, levels []logrus.Level) *Hook {
	return &Hook{client, levels}
}

// Fire triggers a splunk event
func (h *Hook) Fire(entry *logrus.Entry) error {
	event := map[string]string{
		"message": entry.Message,
		"time": entry.Time.String(),
		"level": entry.Level.String(),

	}
	for _, m := range entry.Data {
		for k, v := range m {
			event[k] = v
		}
	}

	err := h.Client.Log(
		event,
	)
	return err
}

// Levels Required for logrus hook implementation
func (h *Hook) Levels() []logrus.Level {
	return h.levels

}
