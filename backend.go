// Copyright (c) 2026 Jeb Bisson.

package adplugadl

import nukedopl3 "github.com/jebbisson/spice-opl3-nuked"

// Backend is the minimal OPL engine surface required by the ADL runtime.
type Backend interface {
	Reset()
	Close()
	WriteRegister(port uint16, reg uint8, val uint8)
	WriteRegisterBuffered(port uint16, reg uint8, val uint8)
	GenerateSamples(n int) ([]int16, error)
	GenerateSamplesWithMeters(n int) ([]int16, []uint16, error)
	SetSoloChannel(ch int)
	SoloChannel() int
	ChannelMeter(ch int) float64
}

// BackendFactory constructs an OPL backend for the requested sample rate.
type BackendFactory func(sampleRate int) Backend

var backendFactory BackendFactory = func(sampleRate int) Backend {
	return nukedopl3.New(uint32(sampleRate))
}

// NewBackend constructs a backend using the currently configured factory.
func NewBackend(sampleRate int) Backend {
	return backendFactory(sampleRate)
}

// SetBackendFactory overrides the backend constructor used by this module.
// Passing nil restores the default Nuked-OPL3-backed implementation.
func SetBackendFactory(factory BackendFactory) {
	if factory == nil {
		backendFactory = func(sampleRate int) Backend {
			return nukedopl3.New(uint32(sampleRate))
		}
		return
	}
	backendFactory = factory
}
