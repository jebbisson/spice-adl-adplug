// Copyright (c) 2026 Jeb Bisson.

package adplugadl

import (
	"fmt"
	"math"
)

// regToFreq converts OPL2 register values (regAx + regBx) to frequency in Hz.
func regToFreq(regAx, regBx uint8) float64 {
	fnum := uint16(regAx) | (uint16(regBx&0x03) << 8)
	block := (regBx >> 2) & 0x07
	if fnum == 0 {
		return 0
	}
	return float64(fnum) * 49716.0 / math.Pow(2, float64(20-block))
}

func freqToNoteName(freq float64) string {
	if freq <= 0 {
		return "C0"
	}

	midiNote := 69.0 + 12.0*math.Log2(freq/440.0)
	rounded := int(math.Round(midiNote))

	if rounded < 0 {
		rounded = 0
	}
	if rounded > 127 {
		rounded = 127
	}

	noteNames := []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	noteName := noteNames[rounded%12]
	octave := (rounded / 12) - 1

	return fmt.Sprintf("%s%d", noteName, octave)
}
