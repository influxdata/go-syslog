package syslog

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestProva(t *testing.T) {
	input := []byte("<22>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [id aaaa] freeformmessage?")
	fsm := NewMachine()
	res, err := fsm.Parse(input)

	spew.Dump(res)
	spew.Dump(err)
}

// [id1 a=\"b\" c=\"d\" e=\"\"][id2 z=\"w\"]
