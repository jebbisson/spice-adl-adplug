# spice-adl-adplug

`spice-adl-adplug` is a thin Go wrapper/repackaging of the ADL playback logic
derived from AdPlug's Westwood ADL support.

## Purpose

This module isolates the LGPL-derived ADL parser and playback runtime from the
higher-level SpiceSynth library.

It currently provides:

- ADL file parsing
- subsong classification helpers
- instrument extraction
- a 72 Hz ADL bytecode runtime
- a simple `io.Reader` player API

## Runtime Introspection

In addition to playback, the runtime exposes structured state for converters and
debugging tools.

### Structured events

`Driver` and `Player` both support `SetEventFunc`, which receives semantic
`ChannelEvent` callbacks during bytecode execution.

Current event kinds:

- `EventInstrumentChange`
- `EventNoteOn`
- `EventNoteOff`
- `EventVolumeChange`

Each event includes a full `ChannelState` snapshot at the moment the event is
emitted. This is intended for higher-level tooling that wants to observe ADL
playback without parsing debug trace strings.

### Channel snapshots

`SnapshotChannels()` and `Player.ChannelStates()` return richer per-channel
state than just note/frequency activity.

Notable fields include:

- raw OPL frequency registers: `RegAx`, `RegBx`
- current instrument ID and note state
- effective operator levels: `CarrierLevel`, `ModulatorLevel`
- base instrument levels: `BaseCarrierLevel`, `BaseModulatorLevel`
- ADL level modifiers: `ExtraLevel1`, `ExtraLevel2`, `ExtraLevel3`
- channel-level instrument settings: `Feedback`, `Connection`

This makes it possible for converter code to reconstruct note-start state and
post-note level changes more directly from the runtime.

## Build Requirements

- Go 1.24.4+
- `CGO_ENABLED=1`
- a working C compiler, because the default backend uses
  `github.com/jebbisson/spice-opl3-nuked`

## License

This repository is distributed under `LGPL-2.1-or-later`.

See:

- `LICENSE`
- `THIRD_PARTY_LICENSES.md`

## Notes

This repository intentionally keeps the ADL-derived logic separate from the main
`spice-synth` module so downstream projects can reason about licensing and
replacement boundaries more clearly.
