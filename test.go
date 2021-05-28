package main

import "fmt"

const helloPrefix = "Hello,"

func main() {
	fmt.Println(miao("miao"))
}

func miao(name string) string {
	return helloPrefix + name
}
