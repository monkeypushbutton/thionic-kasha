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
package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestFirstGolombTest(t *testing.T) {
	var tests = []struct {
		input  []byte
		expect bool
	}{
		{[]byte{0xFF}, false},
		{[]byte{0xF0}, true},
		{[]byte{0x01}, false},
		{[]byte{0xAA}, true},
	}
	for _, c := range tests {
		dc := NewDataCollection(c.input)
		got := dc.FirstGolombTest()
		if c.expect != got {
			t.Errorf("Test failed, input %v expected %v, got: %v", c.input, c.expect, got)
		}
	}
}

func TestSecondGolombTest(t *testing.T) {
	var tests = []struct {
		input  []byte
		expect bool
	}{
		{[]byte{0xFF}, false},
		{[]byte{0xF0}, false},
		{[]byte{0x01}, false},
		{[]byte{0xB2}, true},
	}
	for _, c := range tests {
		dc := NewDataCollection(c.input)
		got := dc.SecondGolombTest()
		if c.expect != got {
			t.Errorf("Test failed, input %v expected %v, got: %v", c.input, c.expect, got)
		}
	}
}

func TestThirdGolombTest(t *testing.T) {
	var tests = []struct {
		input  []byte
		expect bool
	}{
		{[]byte{0xFF}, true},
		{[]byte{0xF0}, false},
		{[]byte{0x01}, false},
		{[]byte{0xB2}, true},
	}
	for _, c := range tests {
		dc := NewDataCollection(c.input)
		got := dc.ThirdGolombTest()
		if c.expect != got {
			t.Errorf("Test failed, input %v expected %v, got: %v", c.input, c.expect, got)
		}
	}
}

func TestAutocorrelation(t *testing.T) {
	var tests = []struct {
		input  []byte
		shift  int
		expect int
	}{
		{[]byte{0xFF}, 1, 8},
		{[]byte{0xF0}, 1, 4},
		{[]byte{0xF0}, 4, -8},
		{[]byte{0x01}, 1, 4},
		{[]byte{0x00}, 1, 8},
	}
	for _, c := range tests {
		dc := NewDataCollection(c.input)
		got := dc.Autocorrelation(c.shift)
		if c.expect != got {
			t.Errorf("Test failed, input %v shift %d, expected %d, got: %d", c.input, c.shift, c.expect, got)
		}
	}
}

func TestZeroOneDiff(t *testing.T) {
	var tests = []struct {
		input  []byte
		expect int
	}{
		{[]byte{0xFF}, 8},
		{[]byte{0xF0}, 0},
		{[]byte{0x01}, -6},
		{[]byte{0x00, 0xF0}, -8},
	}
	for _, c := range tests {
		dc := NewDataCollection(c.input)
		got := dc.ZeroOneDiff()
		if c.expect != got {
			t.Errorf("Test failed, input %v expected %d, got: %d", c.input, c.expect, got)
		}
	}
}

func TestCalculateRunLengths(t *testing.T) {
	var tests = []struct {
		input  []byte
		expect map[int]int
	}{
		{[]byte{0xFF}, map[int]int{8: 1}},
		{[]byte{0xF0}, map[int]int{4: 2}},
		{[]byte{0x01}, map[int]int{1: 1, 7: 1}},
		{[]byte{0x00, 0xF0}, map[int]int{4: 1, 12: 1}},
	}
	for _, c := range tests {
		dc := NewDataCollection(c.input)
		got := dc.CalculateRunLengths()
		if !reflect.DeepEqual(c.expect, got) {
			expectStr, _ := json.MarshalIndent(c.expect, "", "  ")
			gotStr, _ := json.MarshalIndent(got, "", "  ")
			t.Errorf("Test failed, input %v expected %s, got: %s", c.input, expectStr, gotStr)
		}
	}
}
