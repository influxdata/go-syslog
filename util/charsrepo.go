package util

import (
	"sync"
)

// CharsRepo represents a container for UTF-8 characters
type CharsRepo struct {
	mu    sync.Mutex
	chars []uint8
}

// Add a new UTF-8 (uint8) character to the repo
func (cr *CharsRepo) Add(char uint8) {
	cr.mu.Lock()
	cr.chars = append(cr.chars, char)
	cr.mu.Unlock()
}

// clear empties the internal structure retaining characters
func (cr *CharsRepo) clear() {
	cr.chars = nil
}

// Reduce returns the integer representation of the UTF-8 characters added to the chars repo
func (cr *CharsRepo) Reduce() *int {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	defer cr.clear()
	res, _ := UTF8DecimalCodePointsToInt(cr.chars)
	return res
}

// NewCharsRepo creates a CharsRepo
func NewCharsRepo() *CharsRepo {
	return &CharsRepo{
		chars: make([]uint8, 0),
		mu:    sync.Mutex{},
	}
}
