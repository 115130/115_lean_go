package main

import (
	"testing"
)

func TestMiao(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello people", func(t *testing.T) {
		got := miao("chris")
		want := "Hello,chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := miao("")
		want := "Hello,World"
		assertCorrectMessage(t, got, want)
	})
}
