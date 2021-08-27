package main

import "math"

var notes = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

type Note int

func NewNote(note string, octave int) Note {
	var num int
	for i, name := range notes {
		if name == note {
			num = i
		}
	}
	n := Note(octave*12 + num)
	return n.constrain()
}

func (n Note) Name() string {
	return notes[n%12]
}

func (n Note) Octave() int {
	return int(n) / 12
}

func (n Note) Frequency() float64 {
	return 440 * math.Pow(2, float64(n-57)/12)
}

func (n Note) Add(i int) Note {
	n += Note(i)
	return n.constrain()
}

func (n Note) constrain() Note {
	if n < 0 {
		n = 0
	}
	if n >= 108 {
		n = 107
	}
	return n
}
