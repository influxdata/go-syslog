package rfc5424

import (
	"sync"
)

var (
	repo CharsRepo
)

// CharsRepo represents a container for UTF-8 characters
type CharsRepo interface {
	Add(char uint8)
	Reduce() *int
	Clear()
}

type charsRepo struct {
	chars []uint8
	mu    sync.RWMutex
}

func (cr *charsRepo) Add(char uint8) {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	cr.chars = append(cr.chars, char)
}

// Clear empties the internal structure retaining characters
func (cr *charsRepo) Clear() {
	cr.chars = nil
}

// Reduce returns the integer representation of the utf8 characters added to the chars repo
func (cr *charsRepo) Reduce() *int {
	if len(cr.chars) == 0 {
		return nil
	}

	cr.mu.Lock()
	defer cr.mu.Unlock()

	out := 0
	ord := 1
	for i := len(cr.chars) - 1; i >= 0; i-- {
		out += (int(cr.chars[i]) - '0') * ord
		ord *= 10
	}

	cr.Clear()

	return &out
}

// GetCharsRepo returns the characters repository
func GetCharsRepo() CharsRepo {
	repo = &charsRepo{
		chars: make([]uint8, 0),
	}

	return repo
}
