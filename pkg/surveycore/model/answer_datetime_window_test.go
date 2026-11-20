package model

import "testing"

func TestNormalizeAnswerDatetimeWindow(t *testing.T) {
	window := NormalizeAnswerDatetimeWindow([2]string{
		" 2024-03-10 09:00:00 ",
		"bad",
	})
	if window != [2]string{"2024-03-10 09:00:00", ""} {
		t.Fatalf("window = %#v", window)
	}
}
