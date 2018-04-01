package util

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// This test always adds the same
// values to the chars repo so that
// the result is predictable even if we
// are using multiple goroutines.
// This is to test if there's a data race.
func TestConcurrentCharsRepo(t *testing.T) {
	cr := NewCharsRepo()

	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		cr.Add(uint8(49))
		wg.Done()
	}()
	go func() {
		cr.Add(uint8(49))
		wg.Done()
	}()
	go func() {
		cr.Add(uint8(49))
		wg.Done()
	}()
	wg.Wait()
	res := cr.Reduce()

	assert.Equal(t, 111, *res)
}

func TestEmptyCharsRepo(t *testing.T) {
	cr := NewCharsRepo()
	res := cr.Reduce()

	assert.Nil(t, res)
}
