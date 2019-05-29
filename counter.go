package lamportts

import "bytes"

// Counter is a dynamic length byte array that expands to handle its capacity
type Counter []byte

// Increment returns a copy of the counter with its value incremented
func (c Counter) Increment() Counter {
	out := make([]byte, len(c)+1)

	copy(out[1:], c)

	for i := len(out) - 1; i >= 0; i-- {
		msb := byte(0x80)
		if i == len(out)-1 {
			msb = 0x00
		}

		// clear the msb
		cur := out[i] & 0x7F
		cur++

		// if incrementing didn't leave the msb set, break
		if cur&0x80 == 0x00 {
			out[i] = cur | msb
			break
		}

		out[i] = msb
	}

	// exclude the first byte if it doesn't include any data
	if out[0] == 0 {
		return out[1:]
	}
	return out
}

// Compare returns <0 if a < b,  >0 if a > b, and 0 if a == b
func Compare(a, b Counter) int {
	return bytes.Compare(a, b)
}
