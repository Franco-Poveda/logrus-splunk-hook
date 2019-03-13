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
	//TODO: should I map more of these fields over? like err?
	event := map[string]interface{}{
		"message": entry.Message,
		"time": entry.Time.String(),
		"level": entry.Level.String(),
	}
	for k, v := range entry.Data {
		event[k] = v
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
