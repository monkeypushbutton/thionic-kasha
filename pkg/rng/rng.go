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

/*
Package rng implements a common interface for random number generators (they
return a stream of bytes that will need to be debiased to be useful)

Implemented prgns range from the simple LCG generators which are of mainly
historical interest, to staples such as the Mersenne Twister.

TODO:
Mersenne Twister
WELL (Well Equidistributed Long-period Linear)
CryptMT (which is basically MS 8bits of MT)
PRBS
Blum Blum Shub
*/
package rng

type rng interface {
	GenerateNextBytes(int) []byte
}

// Causes the supplied rng to output a sequence of biased unsigned integers.
func GenerateNextUInt32(prng rng, numElements int) []uint32 {
	buffer := make([]uint32, 0, numElements)
	byteBuffer := prng.GenerateNextBytes(numElements * 4)
	for i := 0; i < len(byteBuffer); i = i + 4 {
		buffer = append(buffer,
			uint32(byteBuffer[i])|
				uint32(byteBuffer[i+1])<<8|
				uint32(byteBuffer[i+2])<<16|
				uint32(byteBuffer[i+3])<<24)
	}
	return buffer
}
