// Copyright (c) 2026 Jeb Bisson.

package adplugadl

// Runtime is the minimal playback-engine surface used by Player.
type Runtime interface {
	InitDriver()
	SetVersion(v int)
	SetSoundData(data []byte)
	StartSound(track int, volume uint8)
	StopAllChannels()
	Callback()
	IsChannelPlaying(ch int) bool
	IsChannelRepeating(ch int) bool
	SnapshotChannels() []ChannelState
	SetTraceFunc(fn func(format string, args ...interface{}))
	SetEventFunc(fn ChannelEventFunc)
}

// RuntimeFactory constructs a runtime bound to a backend.
type RuntimeFactory func(opl Backend) Runtime

var runtimeFactory RuntimeFactory = func(opl Backend) Runtime {
	return NewDriver(opl)
}

// NewRuntime constructs a runtime using the configured factory.
func NewRuntime(opl Backend) Runtime {
	return runtimeFactory(opl)
}

// SetRuntimeFactory overrides the runtime constructor used by Player. Passing
// nil restores the in-repo driver implementation.
func SetRuntimeFactory(factory RuntimeFactory) {
	if factory == nil {
		runtimeFactory = func(opl Backend) Runtime {
			return NewDriver(opl)
		}
		return
	}
	runtimeFactory = factory
}
