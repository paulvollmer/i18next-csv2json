package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	testdata := []byte(`key,en,de
hello.world,Hello World,Hallo Welt`)
	langs, data, err := GenerateFromBytes(testdata)
	assert.NoError(t, err)

	assert.Len(t, langs, 2)
	assert.Equal(t, "en", langs[0])
	assert.Equal(t, "de", langs[1])

	assert.Len(t, data, 2)
	assert.Equal(t, `{"hello":{"world":"Hello World"}}`, string(data[0]))
	assert.Equal(t, `{"hello":{"world":"Hallo Welt"}}`, string(data[1]))
}
