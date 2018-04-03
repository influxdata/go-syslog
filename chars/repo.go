package chars

import (
	"sync"
)

// Repo represents a container for UTF-8 characters
type Repo struct {
	mu    sync.Mutex
	chars []uint8
}

// Add a new UTF-8 (uint8) character to the repo
func (cr *Repo) Add(char uint8) {
	cr.mu.Lock()
	cr.chars = append(cr.chars, char)
	cr.mu.Unlock()
}

// clear empties the internal structure retaining characters
func (cr *Repo) clear() {
	cr.chars = nil
}

// ReductionToInt is a reduction function able to reduce a list of UTF-8 chars into a single integer
type ReductionToInt func(chars []uint8) (*int, error)

// ReductionToString is a reduction function able to reduce a list of UTF-8 chars into a single string
type ReductionToString func(chars []uint8) (*string, error)

// ReduceToInt returns the integer representation of the UTF-8 characters within the chars repo
func (cr *Repo) ReduceToInt(fx ReductionToInt) *int {
	cr.mu.Lock()
	defer cr.clear()
	defer cr.mu.Unlock()

	res, _ := fx(cr.chars)
	return res
}

// ReduceToString returns the string representation of the UTF-8 characters within the chars repo
func (cr *Repo) ReduceToString(fx ReductionToString) *string {
	cr.mu.Lock()
	defer cr.clear()
	defer cr.mu.Unlock()

	res, _ := fx(cr.chars)
	return res
}

// NewRepo creates a Repo
func NewRepo() *Repo {
	return &Repo{
		chars: make([]uint8, 0),
		mu:    sync.Mutex{},
	}
}
