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
	myEntry := entry
	myEntry.Logger = nil
	err := h.Client.Log(
		myEntry,
	)
	return err
}

// Levels Required for logrus hook implementation
func (h *Hook) Levels() []logrus.Level {
	return h.levels

}
