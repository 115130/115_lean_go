package main

import "container/list"

const englishHelloPrefix = "Hello,"
const chineseHelloPrefix = "你好,"
const frenchHelloPrefix = "Bonjour,"
const chinese = "chinese"
const french = "french"
const repeatCount = 5
const MAX = 100000000

// func main() {
// 	fmt.Println(miao("miao"))
// }

func Miao(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := greetingPrefix(language)

	return prefix + name
}

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case chinese:
		prefix = chineseHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func TestMap() {
	list := list.New()
	for i := 0; i < MAX; i++ {
		list.PushBack(i)
	}
}

func TestArray() {
	// var lists []int
	for i := 0; i < MAX; i++ {
		// lists = append(lists, i)
	}
}

func TestArray1() {
	var lists [MAX]int
	for i := 0; i < MAX; i++ {
		lists[i] = i
	}
}
