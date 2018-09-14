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

import "math/big"

func NewBsdRand(seed uint32) rng {
	return NewLcg32(seed, 1<<31, 1103515245, 12345)
}

func NewDRand48(seed uint32) rng {
	return NewLcg64(uint64(seed), 1<<48, 25214903917, 11)
}

type Lcg64 struct {
	state         uint64
	modulus       *big.Int
	multiplier    *big.Int
	increment     *big.Int
	bytesConsumed byte
}

func NewLcg64(seed uint64, modulus int64, multiplier int64, increment int64) rng {
	return &Lcg64{seed, big.NewInt(modulus), big.NewInt(multiplier), big.NewInt(increment), 0}
}

func (prng *Lcg64) GenerateNextBytes(numBytes int) []byte {
	buffer := make([]byte, 0, numBytes)
	for bufferIdx := 0; bufferIdx < numBytes; bufferIdx++ {
		if prng.bytesConsumed == 0 {
			tmpState := big.NewInt(0)
			tmpState.SetUint64(prng.state)
			tmpState.Mul(tmpState, prng.multiplier)
			tmpState.Add(tmpState, prng.increment)
			tmpState.Mod(tmpState, prng.modulus)
			prng.state = tmpState.Uint64()
		}
		nextByte := byte((prng.state >> (prng.bytesConsumed * 8)) % 256)
		buffer = append(buffer, nextByte)
		prng.bytesConsumed = (prng.bytesConsumed + 1) % 8
	}
	return buffer
}

type Lcg32 struct {
	state         uint64
	modulus       uint64
	multiplier    uint64
	increment     uint64
	bytesConsumed byte
}

func NewLcg32(seed uint32, modulus uint32, multiplier uint32, increment uint32) rng {
	return &Lcg32{uint64(seed), uint64(modulus), uint64(multiplier), uint64(increment), 0}
}

func (rng *Lcg32) GenerateNextBytes(numBytes int) []byte {
	buffer := make([]byte, 0, numBytes)
	for bufferIdx := 0; bufferIdx < numBytes; bufferIdx++ {
		if rng.bytesConsumed == 0 {
			rng.state = ((rng.state * rng.multiplier) + rng.increment) % rng.modulus
		}
		nextByte := byte((rng.state >> (rng.bytesConsumed * 8)) % 256)
		buffer = append(buffer, nextByte)
		rng.bytesConsumed = (rng.bytesConsumed + 1) % 4
	}
	return buffer
}
