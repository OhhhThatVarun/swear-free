package tests

import (
	swearFree "github.com/OhhhThatVarun/swear-free"
	"testing"
)

func TestNoCensoring(t *testing.T) {
	var expected = "This is a test with no swears words."
	var result = swearFree.CensorWord("This is a test with no swears words.")
	if result != expected {
		t.Errorf("Failed expected %v, got %v", expected, result)
	}
}

func TestCensoring(t *testing.T) {
	var expected = "**** *** Dude"
	var result = swearFree.CensorWord("Fuck ass Dude")
	if result != "**** *** Dude" {
		t.Errorf("Failed expected %v, got %v", expected, result)
	}
}

func TestReplaceCharacter(t *testing.T) {
	var expected = "#### Dude"
	swearFree.SetReplaceCharacter("#")
	var result = swearFree.CensorWord("Fuck Dude")
	if result != "#### Dude" {
		t.Errorf("Failed expected %v, got %v", expected, result)
	}
}
