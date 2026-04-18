package adplugadl

import (
	"testing"
)

func TestParseBytesRejectsTooSmallFile(t *testing.T) {
	_, err := ParseBytes([]byte{0x00, 0x01})
	if err == nil {
		t.Fatal("expected ParseBytes to reject undersized ADL data")
	}
}

func TestExtractInstrumentsEmptyFile(t *testing.T) {
	f := &File{}
	insts := f.ExtractInstruments("test")
	if len(insts) != 0 {
		t.Fatalf("expected no instruments, got %d", len(insts))
	}
}

func TestSnapshotChannelsIncludesDerivedState(t *testing.T) {
	d := NewDriver(nil)
	d.channels[0] = channel{
		dataptr:              123,
		currentInstrumentID: 7,
		rawNote:             0x24,
		regAx:               0x34,
		regBx:               0x32,
		currentNoteDuration: 9,
		duration:            4,
		spacing1:            2,
		spacing2:            1,
		volumeModifier:      0x7F,
		opLevel1:            0x45,
		opLevel2:            0x0B,
		opExtraLevel1:       3,
		opExtraLevel2:       4,
		opExtraLevel3:       5,
		twoChan:             1,
		feedback:            6,
	}

	states := d.SnapshotChannels()
	if len(states) != 10 {
		t.Fatalf("SnapshotChannels() returned %d states, want 10", len(states))
	}

	got := states[0]
	if got.RegAx != 0x34 || got.RegBx != 0x32 {
		t.Fatalf("register snapshot mismatch: got regAx=0x%02X regBx=0x%02X", got.RegAx, got.RegBx)
	}
	if got.BaseModulatorLevel != 0x05 {
		t.Fatalf("BaseModulatorLevel = %d, want 5", got.BaseModulatorLevel)
	}
	if got.BaseCarrierLevel != 0x0B {
		t.Fatalf("BaseCarrierLevel = %d, want 11", got.BaseCarrierLevel)
	}
	if got.ExtraLevel1 != 3 || got.ExtraLevel2 != 4 || got.ExtraLevel3 != 5 {
		t.Fatalf("extra levels = (%d,%d,%d), want (3,4,5)", got.ExtraLevel1, got.ExtraLevel2, got.ExtraLevel3)
	}
	if got.Feedback != 6 {
		t.Fatalf("Feedback = %d, want 6", got.Feedback)
	}
	if got.Connection != 1 {
		t.Fatalf("Connection = %d, want 1", got.Connection)
	}
	if !got.KeyOn {
		t.Fatal("expected KeyOn to be true")
	}
	if !got.TwoOperatorCarrier {
		t.Fatal("expected TwoOperatorCarrier to be true")
	}
}

func TestDriverStructuredEvents(t *testing.T) {
	d := NewDriver(nil)
	d.channels[0] = channel{
		dataptr:              77,
		currentInstrumentID: 5,
		rawNote:             0x24,
		regAx:               0x34,
		regBx:               0x12,
		volumeModifier:      0x7F,
		opLevel1:            0x05,
		opLevel2:            0x0B,
		twoChan:             1,
		feedback:            2,
	}

	var events []ChannelEvent
	d.SetEventFunc(func(ev ChannelEvent) {
		events = append(events, ev)
	})

	d.emitEvent(EventInstrumentChange, 0)
	d.noteOn(&d.channels[0])
	d.noteOff(&d.channels[0])

	if len(events) != 3 {
		t.Fatalf("got %d events, want 3", len(events))
	}
	if events[0].Type != EventInstrumentChange {
		t.Fatalf("event 0 type = %v, want EventInstrumentChange", events[0].Type)
	}
	if events[1].Type != EventNoteOn {
		t.Fatalf("event 1 type = %v, want EventNoteOn", events[1].Type)
	}
	if events[2].Type != EventNoteOff {
		t.Fatalf("event 2 type = %v, want EventNoteOff", events[2].Type)
	}
	if !events[1].State.KeyOn {
		t.Fatal("expected note-on event state to have KeyOn=true")
	}
	if events[2].State.KeyOn {
		t.Fatal("expected note-off event state to have KeyOn=false")
	}
	if events[1].State.InstrumentID != 5 {
		t.Fatalf("event state InstrumentID = %d, want 5", events[1].State.InstrumentID)
	}
}
