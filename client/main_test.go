package main

import (
	_ "embed"
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
