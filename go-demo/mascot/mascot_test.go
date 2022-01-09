package mascot_test

import (
	"testing"

	"example.com/go-demo/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong Mascot :(")
	}
}
