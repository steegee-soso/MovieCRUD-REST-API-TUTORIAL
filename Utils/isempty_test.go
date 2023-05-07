package Utils

import "testing"

func TestEmpty(t *testing.T) {
	actual := isEmptyField("Kofi")
	expected := true
	if expected != actual {
		t.Errorf("Expected %v do not match actual %v", expected, actual)
	} else {
		t.Logf("Expecte Test PASS expected %v ", expected)
	}
}
