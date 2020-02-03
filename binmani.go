// Package binmani provides rudimentary binary manipulation functions.
package binmani

const bitsPerByte int = 8

// Bit manipulation functions

// GetMask creates a bitmask of size shifted left index bits.
// 	GetMask(2, 3) -> 00011100
func GetMask(index, size uint8) uint16 {
	return ((1 << size) - 1) << index
}

// ReadFrom reads a specified bit or set of consecutive bits from data.
// index works from the right of the data to the left.
func ReadFrom(data uint16, index, size uint8) uint16 {
	return (data & GetMask(index, size)) >> index
}

// WriteTo writes a value to a specified bit or set of consecutive bits in data, and returns the result.
// index works from the right of the data to the left.
func WriteTo(data uint16, index, size uint8, value uint16) uint16 {
	return (data & (^GetMask(index, size))) | (value << index)
}

// BytesToBits converts a byte slice to a slice of each individual bit of the bytes.
func BytesToBits(bytes *[]byte) *[]uint8 {
	bits := make([]uint8, len(*bytes) * bitsPerByte)
	for i := 0; i < len(*bytes); i++ {
		for j := 0; j < bitsPerByte; j++ {
			bits[i * bitsPerByte + j] = uint8(ReadFrom(uint16((*bytes)[i]), uint8(bitsPerByte - j - 1), 1))
		}
	}
	return &bits
}

// BitsToBytes converts a slice of individual bits into a slice of bytes, effectively compressing them together.
func BitsToBytes(bits []uint8, padStart bool) *[]byte {
	numBytes := len(bits) / bitsPerByte
	if len(bits) % bitsPerByte != 0 {
		numBytes++
	}

	// Zero-pad the beginning/end of the array if the number of bits is not a multiple of 8
	extraBits := make([]uint8, len(bits) % bitsPerByte)
	if padStart {
		bits = append(extraBits, bits...)
	} else {
		bits = append(bits, extraBits...)
	}

	bytes := make([]byte, numBytes)
	for i := 0; i < numBytes; i++ {
		for j := 0; j < bitsPerByte; j++ {
			bytes[i] = byte(WriteTo(uint16(bytes[i]), uint8(bitsPerByte - j - 1), 1, uint16(bits[i * bitsPerByte + j])))
		}
	}
	return &bytes
}