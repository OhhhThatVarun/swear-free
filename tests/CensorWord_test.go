package tests

import (
	swearFree "github.com/OhhhThatVarun/swear-free"
	"testing"
)

func TestNoCensoring(t *testing.T) {
	var expected = "This is a test with no swear words."
	var result = swearFree.CensorWord("This is a test with no swear words.")
	if result != expected {
		t.Errorf("Failed expected %v, got %v", expected, result)
	}
}

func TestCensoring(t *testing.T) {
	var expected = "**** my *** Dude"
	var result = swearFree.CensorWord("Fuck my ass Dude")
	if result != expected {
		t.Errorf("Failed expected %v, got %v", expected, result)
	}
}

func TestReplaceCharacter(t *testing.T) {
	var expected = "#### Dude"
	swearFree.SetReplaceCharacter("#")
	var result = swearFree.CensorWord("Fuck Dude")
	if result != expected {
		t.Errorf("Failed expected %v, got %v", expected, result)
	}
}
