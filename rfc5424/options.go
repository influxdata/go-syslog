package rfc5424

import (
	syslog "github.com/influxdata/go-syslog/v3"
)

// WithBestEffort enables the best effort mode.
func WithBestEffort() syslog.MachineOption {
	return func(m syslog.Machine) syslog.Machine {
		m.WithBestEffort()
		return m
	}
}

// AllowNonUTF8InMessage allows non utf8 characters in the message part.
func AllowNonUTF8InMessage() syslog.MachineOption {
	return func(m syslog.Machine) syslog.Machine {
		m.(*machine).allowNonUTF8InMessage = true
		return m
	}
}
