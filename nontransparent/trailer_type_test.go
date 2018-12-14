package nontransparent

import (
	"encoding/json"

	"github.com/stretchr/testify/assert"

	"testing"
)

type trailerWrapper struct {
	Trailer TrailerType `json:"trailer"`
}

func TestUnmarshalLowercase(t *testing.T) {
	x := &trailerWrapper{}
	in := []byte(`{"trailer": "lf"}`)
	err := json.Unmarshal(in, x)
	assert.Nil(t, err)
	assert.Equal(t, &trailerWrapper{Trailer: LF}, x)
}

func TestUnmarshalUnknown(t *testing.T) {
	x := &trailerWrapper{}
	in := []byte(`{"trailer": "UNK"}`)
	err := json.Unmarshal(in, x)
	assert.Error(t, err)
	assert.Equal(t, &trailerWrapper{Trailer: -1}, x)
}

func TestUnmarshal(t *testing.T) {
	x := &trailerWrapper{}
	in := []byte(`{"trailer": "NUL"}`)
	err := json.Unmarshal(in, x)
	assert.Nil(t, err)
	assert.Equal(t, &trailerWrapper{Trailer: NUL}, x)
}

func TestMarshalUnknown(t *testing.T) {
	res, err := json.Marshal(&trailerWrapper{Trailer: TrailerType(-2)})
	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestMarshal(t *testing.T) {
	res, err := json.Marshal(&trailerWrapper{Trailer: NUL})
	assert.Nil(t, err)
	assert.Equal(t, `{"trailer":"NUL"}`, string(res))
}
