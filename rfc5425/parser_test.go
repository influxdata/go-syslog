package rfc5425

import (
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestAbc(t *testing.T) {
	// r := strings.NewReader("16 <1>1 - - - - - -17 <2>12 A B C D E -")

	// r := strings.NewReader("16 <1>1 A B C D E -xaaa") // longer that MSGLEN

	// what to do in such situation?
	// message is minimally valid but it does not respect MSGLEN
	// discard?
	// r := strings.NewReader("16 <1>1") // shorter than MSGLEN

	r := strings.NewReader("23 <1>1 - - - - - - hellø") // msglen is the octet count

	results := NewParser(r, true).Parse()
	//fmt.Printf("%T\n", results)

	for r := range results {
		spew.Dump(r)
	}

	// what to do when error is a token error? (1) try to continue the parse? or (2)stop?
	// in case (1) the parse will stop only when EOF encountered
	// this choice is about strictiness ..

	// do not return channel but passing it so the user can use buffered channels to bulk N messages?
}
