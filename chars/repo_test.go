package chars

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
func TestConcurrentRepo(t *testing.T) {
	cr := NewRepo()

	var wg sync.WaitGroup

	wg = sync.WaitGroup{}
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
	resInt := cr.ReduceToInt(UTF8DecimalCodePointsToInt)

	assert.Equal(t, 111, *resInt)

	wg = sync.WaitGroup{}
	wg.Add(3)
	go func() {
		cr.Add(uint8(109))
		wg.Done()
	}()
	go func() {
		cr.Add(uint8(109))
		wg.Done()
	}()
	go func() {
		cr.Add(uint8(109))
		wg.Done()
	}()
	wg.Wait()
	resStr := cr.ReduceToString(UTF8DecimalCodePointsToString)

	assert.Equal(t, "mmm", *resStr)
}

func TestEmptyRepo(t *testing.T) {
	cr := NewRepo()
	res := cr.ReduceToInt(UTF8DecimalCodePointsToInt)

	assert.Nil(t, res)
}
