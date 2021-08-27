package main

import (
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

const SAMPLERATE = 44100

func main() {

	sr := beep.SampleRate(SAMPLERATE)
	err := speaker.Init(sr, sr.N(time.Second))
	if err != nil {
		log.Fatal(err)
	}

	tone := &Tone{
		rate: SAMPLERATE,
		freq: 261.63,
	}

	speaker.Play(tone)
	select {}

}

type Tone struct {
	rate       float64
	freq       float64
	sinePhase  float64
	sawPhase   float64
	noisePhase float64
	prev       float64
	n          int
}

var CMajorScale = []Note{
	NewNote("C", 4),
	NewNote("D", 4),
	NewNote("E", 4),
	NewNote("F", 4),
	NewNote("G", 4),
	NewNote("A", 4),
	NewNote("B", 4),
	NewNote("C", 5),
	NewNote("B", 4),
	NewNote("A", 4),
	NewNote("G", 4),
	NewNote("F", 4),
	NewNote("E", 4),
	NewNote("C", 4),
}

// var keyboard = map[string]Note{
// 	"b":
// }

func (t *Tone) Stream(samples [][2]float64) (int, bool) {
	for i := range samples {

		quarterNote := t.n * 4 / int(t.rate)
		if quarterNote >= len(CMajorScale) {
			quarterNote = 0
			t.n = 0
		}

		freq := CMajorScale[quarterNote].Frequency()
		val := Sine(&t.sinePhase, freq, t.rate, t.prev)
		val += 0.1 * Saw(&t.sawPhase, freq, t.rate, t.prev)

		samples[i][0] = val
		samples[i][1] = val

		t.n++
		t.prev = val
	}
	return len(samples), true
}

func (t *Tone) Err() error {
	return nil
}
