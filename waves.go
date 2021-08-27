package main

import (
	"math"
	"math/rand"
)

func Sine(phase *float64, freq float64, rate float64, prev float64) float64 {

	delta := 2 * math.Pi * freq / rate
	*phase += delta

	// keep phase between 0 and 2*pi
	for *phase >= 2*math.Pi {
		*phase -= 2 * math.Pi
	}
	for *phase < 0 {
		*phase += 2 * math.Pi
	}

	return math.Sin(*phase)

}

func Square(phase *float64, freq float64, rate float64, prev float64) float64 {

	*phase += freq / rate

	for *phase > 1 {
		*phase -= 1
	}
	for *phase < 0 {
		*phase += 1
	}

	if *phase <= 0.5 {
		return -1
	}

	return 1

}

func Saw(phase *float64, freq float64, rate float64, prev float64) float64 {

	*phase += freq / rate

	for *phase > 1 {
		*phase -= 1
	}
	for *phase < 0 {
		*phase += 1
	}

	return (*phase * 2) - 1

}

func Triangle(phase *float64, freq float64, rate float64, prev float64) float64 {

	*phase += freq / rate

	for *phase > 1 {
		*phase -= 1
	}
	for *phase < 0 {
		*phase += 1
	}

	val := *phase * 2
	if *phase > 0.5 {
		val = (1 - *phase) * 2
	}

	return (val * 2) - 1

}

func Noise(phase *float64, freq float64, rate float64, prev float64) float64 {

	prevSeed := int64(*phase)
	*phase += freq / rate
	seed := int64(*phase)

	for *phase > 2 {
		*phase -= 1
	}

	if seed == prevSeed {
		return prev
	}

	val := rand.Float64()

	return (val * 2) - 1

}
