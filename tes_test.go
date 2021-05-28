package main

import "testing"

func TestMiao(t *testing.T) {
	got := miao()
	want := "miao"
	if got != want {
		t.Errorf("got是 '%q' want是 '%q123'", got, want)
	}
}
