package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNotes(t *testing.T) {

	n := NewNote("C", 0)
	assert.Equal(t, "C", n.Name())
	assert.Equal(t, 0, n.Octave())
	assert.InDelta(t, 16.35, n.Frequency(), 0.01)

	n = NewNote("A", 0)
	assert.Equal(t, "A", n.Name())
	assert.Equal(t, 0, n.Octave())
	assert.InDelta(t, 27.50, n.Frequency(), 0.01)

	n = NewNote("C", 2)
	assert.Equal(t, "C", n.Name())
	assert.Equal(t, 2, n.Octave())
	assert.InDelta(t, 65.41, n.Frequency(), 0.01)

	n = NewNote("F#", 5)
	assert.Equal(t, "F#", n.Name())
	assert.Equal(t, 5, n.Octave())
	assert.InDelta(t, 739.99, n.Frequency(), 0.01)

	n = NewNote("B", 8)
	assert.Equal(t, "B", n.Name())
	assert.Equal(t, 8, n.Octave())
	assert.InDelta(t, 7902.13, n.Frequency(), 0.01)

}