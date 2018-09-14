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

import (
	"encoding/json"
	"reflect"
	"testing"
)

func BenchmarkGenerateNextBytes32(b *testing.B) {
	lcg := NewLcg32(1, (2<<31)-1, 16807, 0)
	lcg.GenerateNextBytes(b.N)
}

func TestGenerateNextBytes(t *testing.T) {
	var tests = []struct {
		generator rng
		numBytes  int
		expect    []byte
	}{
		{NewDRand48(0), 16, []byte{
			11, 0, 0, 0, 0, 0, 0, 0,
			0xBA, 0xE6, 0x2D, 0x94, 0x40, 0, 0, 0}},
	}
	for _, c := range tests {
		got := c.generator.GenerateNextBytes(c.numBytes)
		if !reflect.DeepEqual(c.expect, got) {
			expectStr, _ := json.MarshalIndent(c.expect, "", "  ")
			gotStr, _ := json.MarshalIndent(got, "", "  ")
			t.Errorf("Test failed, expected %s, got: %s", expectStr, gotStr)
		}
	}
}

// func printSlice(s []byte) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// }
