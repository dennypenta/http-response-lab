package main

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/j1.json
var j1 []byte

//go:embed testdata/j2.json
var j2 []byte

//go:embed testdata/j3.json
var j3 []byte

//go:embed testdata/j4.json
var j4 []byte

//go:embed testdata/j5.json
var j5 []byte

//go:embed testdata/j1_expected.json
var j1Expected string

//go:embed testdata/j2_expected.json
var j2Expected string

//go:embed testdata/j3_expected.json
var j3Expected string

//go:embed testdata/j4_expected.json
var j4Expected string

//go:embed testdata/j5_expected.json
var j5Expected string

func TestSplitFunc(t *testing.T) {
	type testCase struct {
		input    []byte
		expected string
	}
	for i, tt := range []testCase{
		{input: j1, expected: j1Expected},
		{input: j2, expected: j2Expected},
		{input: j3, expected: j3Expected},
		{input: j4, expected: j4Expected},
		{input: j5, expected: j5Expected},
	} {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			split := newSplitFunc()
			advance, buf, err := split(tt.input, false)

			if err != nil {
				t.Error(err)
				t.FailNow()
			}

			assert.JSONEq(t, tt.expected, string(buf))
			// if string(buf) != tt.expected {
			// 	t.Log(string(buf))
			// 	t.Error("jsons must be equal")
			// 	t.FailNow()
			// }

			_ = advance
		})
	}
}
