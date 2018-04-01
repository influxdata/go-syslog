package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleUTF8DecimalConversion(t *testing.T) {
	slice := []uint8{49, 48, 49}
	res, err := UTF8DecimalCodePointsToInt(slice)
	assert.Nil(t, err)
	assert.Equal(t, 101, *res)
}

func TestNumberStartingWithZero(t *testing.T) {
	slice := []uint8{48, 48, 50}
	res, err := UTF8DecimalCodePointsToInt(slice)
	assert.Nil(t, err)
	assert.Equal(t, 2, *res)
}

func TestNonNumberChars(t *testing.T) {
	slice := []uint8{10} // Line Feed (LF)
	res, err := UTF8DecimalCodePointsToInt(slice)
	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestAllDigits(t *testing.T) {
	slice := []uint8{49, 50, 51, 52, 53, 54, 55, 56, 57, 48}
	res, err := UTF8DecimalCodePointsToInt(slice)
	assert.Nil(t, err)
	assert.Equal(t, 1234567890, *res)
}
