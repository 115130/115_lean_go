package main

const englishHelloPrefix = "Hello,"
const chineseHelloPrefix = "你好,"

// func main() {
// 	fmt.Println(miao("miao"))
// }

func miao(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "chinese" {
		return chineseHelloPrefix + name
	}
	return englishHelloPrefix + name
}
