package morsecode

import (
	"testing"
)

func TestEncodeLowerABC(t *testing.T) {
	input := "abc"
	want := ".- -... -.-."
	output, _ := Encode(input)
	if output != want {
		t.Fatalf("Output: %s does not match want: %s\n", output, want)
	}
}

func TestEncodeMixedCaseRyan(t *testing.T) {
	input := "Ryan"
	want := ".-. -.-- .- -."
	output, _ := Encode(input)
	if output != want {
		t.Fatalf("Output: %s does not match want: %s\n", output, want)
	}
}

func TestEncodeWithSpaces(t *testing.T) {
	input := "Ryan says Hej"
	want := ".-. -.-- .- -. ... .- -.-- ... .... . .---"
	output, _ := Encode(input)
	if output != want {
		t.Fatalf("Output: %s does not match want: %s\n", output, want)
	}
}

func TestEncodeWithNonAlphaChars(t *testing.T) {
	input := "Hej!"
	want := ".... . .---"
	output, _ := Encode(input)
	if output != want {
		t.Fatalf("Output: %s does not match want: %s\n", output, want)
	}
}
