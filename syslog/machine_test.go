package syslog

import (
	"testing"
)

func TestProva(t *testing.T) {
	input := []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [id1 a=\"b\" c=\"d\" e=\"\"][id2 z=\"w\"] freeformmessage?")
	fsm := NewMachine()
	fsm.Parse(input)
}
