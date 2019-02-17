package helpers

import "testing"

func TestToInt(t *testing.T) {
	if ToInt("1") != 1 {
		t.Error("ToInt failed")
	}
	if ToInt("0") != 0 {
		t.Error("ToInt failed")
	}
	if ToInt("12345") != 12345 {
		t.Error("ToInt failed")
	}
	if ToInt("-") != 0 {
		t.Error("ToInt failed")
	}
	if ToInt("hello") != 0 {
		t.Error("ToInt failed")
	}
	if ToInt("") != 0 {
		t.Error("ToInt failed")
	}
	if ToInt("\n    ") != 0 {
		t.Error("ToInt failed")
	}
}
