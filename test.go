package main

const englishHelloPrefix = "Hello,"
const chineseHelloPrefix = "你好,"
const frenchHelloPrefix = "Bonjour,"
const chinese = "chinese"
const french = "french"

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
