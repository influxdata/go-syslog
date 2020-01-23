package rfc3164

import (
	syslog "github.com/influxdata/go-syslog/v2"
)

// WithBestEffort enables the best effort mode.
func WithBestEffort() syslog.MachineOption {
	return func(m syslog.Machine) syslog.Machine {
		m.WithBestEffort()
		return m
	}
}

// WithYear sets the strategy to decide the year for the Stamp timestamp of RFC 3164.
func WithYear(o YearOperator) syslog.MachineOption {
	return func(m syslog.Machine) syslog.Machine {
		m.(*machine).WithYear(o)
		return m
	}
}
