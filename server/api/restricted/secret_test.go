package restricted_test

import (
	"morseme/server/api/restricted"
	"testing"
)

func TestGenerateSecretLength(t *testing.T) {
	want := 64
	output := len(restricted.GenerateSecret())
	if output != want {
		t.Fatalf("Output length %d does not match wanted length: %d\n", output, want)
	}
}
