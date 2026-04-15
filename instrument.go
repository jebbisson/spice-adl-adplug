// Copyright (c) 2026 Jeb Bisson.

package adplugadl

// Instrument represents an extracted OPL instrument definition.
type Instrument struct {
	Name string
	Op1  Operator
	Op2  Operator

	Feedback   uint8
	Connection uint8
}

// Operator represents a single OPL operator definition.
type Operator struct {
	Attack        uint8
	Decay         uint8
	Sustain       uint8
	Release       uint8
	Level         uint8
	Multiply      uint8
	KeyScaleRate  bool
	KeyScaleLevel uint8
	Tremolo       bool
	Vibrato       bool
	Sustaining    bool
	Waveform      uint8
}
