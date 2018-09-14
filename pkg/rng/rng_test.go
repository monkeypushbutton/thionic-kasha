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

func TestGenerateNextUInt32(t *testing.T) {
	var tests = []struct {
		generator rng
		numBytes  int
		expect    []uint32
	}{
		{NewBsdRand(0), 5, []uint32{12345, 1406932606, 654583775, 1449466924, 229283573}},
		{NewBsdRand(1), 5, []uint32{1103527590, 377401575, 662824084, 1147902781, 2035015474}},
	}
	for _, c := range tests {
		got := GenerateNextUInt32(c.generator, c.numBytes)
		if !reflect.DeepEqual(c.expect, got) {
			expectStr, _ := json.MarshalIndent(c.expect, "", "  ")
			gotStr, _ := json.MarshalIndent(got, "", "  ")
			t.Errorf("Test failed, expected %s, got: %s", expectStr, gotStr)
		}
	}
}
