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

// Package main is a program that subjests data to various radomness tests.
package main

func main() {

}

type DataCollection struct {
	data []byte
}

func NewDataCollection(data []byte) *DataCollection {
	dc := new(DataCollection)
	dc.data = data
	return dc
}

// ValueAt returns the Nth item.
func (c *DataCollection) ValueAt(n int) bool {
	byteIndex := n / 8
	bitIndex := (uint)(n % 8)
	mask := (byte)(1 << bitIndex)
	return c.data[byteIndex]&mask != 0
}

// Len returns the number of items in the collection
func (c *DataCollection) Len() int { return len(c.data) * 8 }

// golomb implements the Golomb’s randomness postulate test, as described at
// http://cacr.uwaterloo.ca/hac/about/chap5.pdf section 5.4.3
func (c *DataCollection) FirstGolombTest() bool {
	zeroOneDiff := c.ZeroOneDiff()
	return -1 <= zeroOneDiff && zeroOneDiff <= 1
}

func (c *DataCollection) SecondGolombTest() bool {
	runLengths := c.CalculateRunLengths()
	currLen := 1
	for expectRunsAtLen := c.Len() / 2; expectRunsAtLen > 1; expectRunsAtLen = expectRunsAtLen / 2 {
		if runLengths[currLen] < expectRunsAtLen {
			return false
		}
		currLen++
	}
	return true
}

func (c *DataCollection) ThirdGolombTest() bool {
	firstAutocorrelation := c.Autocorrelation(1)
	for k := 2; k < c.Len(); k++ {
		if firstAutocorrelation != c.Autocorrelation(k) {
			return false
		}
	}
	return true
}

func (c *DataCollection) Autocorrelation(shift int) (autocorr int) {
	for i := 0; i < c.Len(); i++ {
		// (2s_i - 1)(2s_i+shift - 1)
		j := (i + shift) % c.Len()
		if c.ValueAt(i) == c.ValueAt(j) {
			autocorr++
		} else {
			autocorr--
		}
	}
	return
}

func (c *DataCollection) ZeroOneDiff() (difference int) {
	for i := 0; i < c.Len(); i++ {
		isBitSet := c.ValueAt(i)
		if isBitSet {
			difference++
		} else {
			difference--
		}
	}
	return
}

func (c *DataCollection) CalculateRunLengths() (runLengths map[int]int) {
	runLengths = make(map[int]int)
	currRunLen := 0
	lastBitSet := false
	first := true
	for i := 0; i < c.Len(); i++ {
		isBitSet := c.ValueAt(i)
		if first || lastBitSet == isBitSet {
			currRunLen++
		} else {
			runLengths[currRunLen] = runLengths[currRunLen] + 1
			currRunLen = 1
		}
		lastBitSet = isBitSet
		first = false
	}
	runLengths[currRunLen] = runLengths[currRunLen] + 1
	return
}
