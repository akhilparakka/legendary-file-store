package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathTranspformFunc(t *testing.T) {
	key := "specialpic"
	pathKey := CASPathTransformFunc(key)
	expectedFilename := "2dba7048445a58cf9734188e62a4011f574cec00"
	expectedPathname := "2dba7/04844/5a58c/f9734/188e6/2a401/1f574/cec00"
	assert.Equal(t, pathKey.Pathname, expectedPathname)
	assert.Equal(t, pathKey.Filename, expectedFilename)
}

func TestDelete(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "momsspecials"
	data := []byte("some jpg bytes")

	assert.Nil(t, s.writeStream(key, bytes.NewReader(data)))

	assert.Nil(t, s.Delete(key))
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "momsspecials"
	data := []byte("some jpg bytes")

	assert.Nil(t, s.writeStream(key, bytes.NewReader(data)))

	r, err := s.Read(key)
	assert.Nil(t, err)

	b, err := io.ReadAll(r)
	assert.Nil(t, err)

	assert.Equal(t, b, data)

	s.Delete(key)
}
