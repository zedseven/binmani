package binmani

import "testing"

func TestGetMask(t *testing.T) {
	got := GetMask(4, 1)
	if got != 0b00010000 {
		t.Errorf("GetMask(4, 1) = %016b, want 0b00010000", got)
	}

	got = GetMask(2, 3)
	if got != 0b00011100 {
		t.Errorf("GetMask(2, 3) = %016b, want 0b00011100", got)
	}

	got = GetMask(0, 8)
	if got != 0b11111111 {
		t.Errorf("GetMask(0, 8) = %016b, want 0b11111111", got)
	}

	got = GetMask(3, 8)
	if got != 0b0000011111111000 {
		t.Errorf("GetMask(3, 8) = %016b, want 0b0000011111111000", got)
	}
}

func TestReadFrom(t *testing.T) {
	in := uint16(0b0101011101100110)

	got := ReadFrom(in, 5, 1)
	if got != 0b00000001 {
		t.Errorf("ReadFrom(%016b, 5, 1) = %016b, want 0b00000001", in, got)
	}

	got = ReadFrom(in, 5, 3)
	if got != 0b00000011 {
		t.Errorf("ReadFrom(%016b, 5, 3) = %016b, want 0b00000011", in, got)
	}

	got = ReadFrom(in, 0, 16)
	if got != in {
		t.Errorf("ReadFrom(%016b, 0, 16) = %016b, want %016b", in, got, in)
	}
}

func TestWriteTo(t *testing.T) {
	in := uint16(0b0101011101100110)

	got := WriteTo(in, 7, 1, 0)
	if got != in {
		t.Errorf("WriteTo(%016b, 7, 1, 0) = %016b, want %016b", in, got, in)
	}

	got = WriteTo(in, 7, 1, 1)
	if got != 0b0101011111100110 {
		t.Errorf("WriteTo(%016b, 7, 1, 1) = %016b, want 0b0101011111100110", in, got)
	}

	got = WriteTo(in, 7, 3, 0b00000101)
	if got != 0b0101011011100110 {
		t.Errorf("WriteTo(%016b, 7, 3, 0b00000101) = %016b, want 0b0101011011100110", in, got)
	}
}

func TestBytesToBits(t *testing.T) {
	in := []byte{}
	expect := []uint8{}
	got := BytesToBits(in)
	if !arrsEqual(*got, expect) {
		t.Errorf("BytesToBits(%v) = %v, want %v", in, *got, expect)
	}

	in = []byte{0x61, 0x62, 0x63, 0x64}
	expect = []uint8{0, 1, 1, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 0, 0}
	got = BytesToBits(in)
	if !arrsEqual(*got, expect) {
		t.Errorf("BytesToBits(%v) = %v, want %v", in, *got, expect)
	}
}

func TestBitsToBytes(t *testing.T) {
	in := []uint8{}
	expect := []byte{}
	got := BitsToBytes(in, true)
	if !arrsEqual(*got, expect) {
		t.Errorf("BitsToBytes(%v, true) = %v, want %v", in, *got, expect)
	}

	in = []uint8{0, 1, 1, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 0, 0}
	expect = []byte{0x61, 0x62, 0x63, 0x64}
	got = BitsToBytes(in, true)
	if !arrsEqual(*got, expect) {
		t.Errorf("BitsToBytes(%v, true) = %v, want %v", in, *got, expect)
	}

	in = []uint8{0, 1, 1, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1}
	expect = []byte{0x61, 0x62, 0x63, 0x64}
	got = BitsToBytes(in, false)
	if !arrsEqual(*got, expect) {
		t.Errorf("BitsToBytes(%v, false) = %v, want %v", in, *got, expect)
	}

	in = []uint8{1, 1, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 0, 0}
	expect = []byte{0x61, 0x62, 0x63, 0x64}
	got = BitsToBytes(in, true)
	if !arrsEqual(*got, expect) {
		t.Errorf("BitsToBytes(%v, true) = %v, want %v", in, *got, expect)
	}
}

func arrsEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}