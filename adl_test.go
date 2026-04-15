package adplugadl

import "testing"

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
