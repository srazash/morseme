package morsecode_test

import (
	"morseme/server/morsecode"
	"testing"
)

func TestEncodeLowerABC(t *testing.T) {
	input := "abc"
	want := ".- -... -.-."
	output, _ := morsecode.Encode(input)
	if output != want {
		t.Fatalf("Output: %s does not match want: %s\n", output, want)
	}
}

func TestEncodeMixedCaseRyan(t *testing.T) {
	input := "Ryan"
	want := ".-. -.-- .- -."
	output, _ := morsecode.Encode(input)
	if output != want {
		t.Fatalf("Output: %s does not match want: %s\n", output, want)
	}
}

func TestEncodeWithSpaces(t *testing.T) {
	input := "Ryan says Hej"
	want := ".-. -.-- .- -. ... .- -.-- ... .... . .---"
	output, _ := morsecode.Encode(input)
	if output != want {
		t.Fatalf("Output: %s does not match want: %s\n", output, want)
	}
}

func TestEncodeWithNonAlphaChars(t *testing.T) {
	input := "Hej!"
	want := ""
	output, err := morsecode.Encode(input)
	if output != want && err != nil {
		t.Fatalf("Output: %s does not match want: %s, and error should not be nil\n", output, want)
	}
}
