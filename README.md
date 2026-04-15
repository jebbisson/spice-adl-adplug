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
