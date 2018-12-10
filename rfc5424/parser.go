package rfc5424

import (
	"github.com/influxdata/go-syslog"
	"sync"
)

// Parser represent a RFC5424 parser with mutex capabilities.
type Parser struct {
	sync.Mutex
	*machine
}

// NewParser creates a new parser and the underlying FSM.
func NewParser(options ...syslog.MachineOption) syslog.Machine {
	p := &Parser{
		machine: NewMachine(options...).(*machine),
	}

	return p
}

// HasBestEffort tells whether the receiving parser has best effort mode on or off.
func (p *Parser) HasBestEffort() bool {
	return p.bestEffort
}

// Parse parses the input RFC5424 syslog message using its FSM.
//
// Best effort mode enables the partial parsing.
func (p *Parser) Parse(input []byte) (syslog.Message, error) {
	p.Lock()
	defer p.Unlock()

	msg, err := p.machine.Parse(input)
	if err != nil {
		if p.bestEffort {
			return msg, err
		}
		return nil, err
	}

	return msg, nil
}
