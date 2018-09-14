/*
MIT License

Copyright (c) 2018 David Hatfield

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package rng

// MT19937
func NewMT19937(seed uint32) rng {
	w, n, m, r := uint(32), uint(624), uint(397), uint(31)
	a := uint64(0x9908B0DF16)
	u, d := uint64(11), uint64(0xFFFFFFFF16)
	s, b := uint64(7), uint64(0x9D2C568016)
	t, c := uint64(15), uint64(0xEFC6000016)
	l := uint64(18)
	f := uint64(1812433253)

	state := make([]uint64, 0, n)
	last := uint64(seed)
	state = append(state, last)
	for i := uint(1); i < n; i++ {
		last = (f*(last^(last>>30)) + uint64(i)) & 0xFFFFFFFF
		state = append(state, last)
	}

	return &MersenneTwister{w, n, m, r, a, u, d, s, b, t, c, l, state, n - 1, 0, 0}
}

type MersenneTwister struct {
	w, n, m, r uint
	a          uint64
	u, d       uint64
	s, b       uint64
	t, c       uint64
	l          uint64
	state      []uint64
	index      uint

	y             uint64
	bytesConsumed uint
}

func (prng *MersenneTwister) GenerateNextBytes(numBytes int) []byte {
	buffer := make([]byte, 0, numBytes)
	for bufferIdx := 0; bufferIdx < numBytes; bufferIdx++ {
		if prng.bytesConsumed == 0 {
			if prng.index >= prng.n {
				prng.twist()
			}
			y := uint64(prng.state[bufferIdx])
			y = y ^ ((y >> prng.u) & prng.d)
			y = y ^ ((y << prng.s) & prng.b)
			y = y ^ ((y << prng.t) & prng.c)
			y = y ^ (y >> prng.l)
			prng.y = y
		}
		nextByte := byte((prng.y >> (prng.bytesConsumed * 8)) % 256)
		buffer = append(buffer, nextByte)
		prng.bytesConsumed = (prng.bytesConsumed + 1) % 4
	}
	return buffer
}

// Generate the next n values from the series x_i
func (prng *MersenneTwister) twist() {
	lowerMask := uint64((1 << prng.r) - 1)
	// lowest w bits of (not lower_mask)
	upperMask := ^lowerMask
	for i := uint(0); i < prng.n-1; i++ {
		x := (prng.state[i] & upperMask) + (prng.state[i+1] & lowerMask)
		xA := uint64(x >> 1)
		if (x % 2) != 0 {
			xA = xA ^ prng.a
		}
		prng.state[i] = prng.state[(i+prng.m)] ^ xA
	}
}
