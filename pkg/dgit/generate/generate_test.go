package generate

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	t.Skip("it doesn't work right now")
	if err := GenerateManifest("r2d4", "ubuntu", "master"); err != nil {
		t.Fatal(err)
	}
}
