package main

import (
	"reflect"
	"testGo/array"
	"testing"
)

func TestSumAll(t *testing.T) {

	got := array.SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want int, number []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d given ,%v", got, want, number)
		}
	}
	t.Run("测试", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := array.Sum(numbers)
		want := 15
		assertCorrectMessage(t, got, want, numbers)
	})
	t.Run("测试", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := array.Sum(numbers)
		want := 6
		assertCorrectMessage(t, got, want, numbers)
	})
	t.Run("测试", func(t *testing.T) {
		numbers := []int{0}

		got := array.Sum(numbers)
		want := 0
		assertCorrectMessage(t, got, want, numbers)
	})

}

func TestMiao(t *testing.T) { //普通测试
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello people", func(t *testing.T) {
		got := Miao("chris", "")
		want := "Hello,chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Miao("", "")
		want := "Hello,World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in chinese", func(t *testing.T) {
		got := Miao("喵", "chinese")
		want := "你好,喵"
		assertCorrectMessage(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) { //基准测试 可以测试速度
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
func BenchmarkRepeat0(b *testing.B) { //基准测试
	for i := 0; i < b.N; i++ {
		TestArray1()
	}
}
func BenchmarkRepeat1(b *testing.B) { //基准测试
	for i := 0; i < b.N; i++ {
		TestArray()
	}
}
func BenchmarkRepeat2(b *testing.B) { //基准测试
	for i := 0; i < b.N; i++ {
		TestMap()
	}
}
