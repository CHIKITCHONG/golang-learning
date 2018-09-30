package main

import (
	"fmt"
	"unicode/utf8"
)

// 字符串

/*

使用 range for 遍历字符串。
range 返回一个 rune （在 byte 数组中）的位置，以及 rune 本身。上面的程序输出为：

*/


//func printCharsAndBytes(s string) {
//	for index, rune := range s {
//		fmt.Printf("%c starts at byte %d\n", rune, index)
//	}
//}
//
//
//func main() {
//	name := "Señor"
//	printCharsAndBytes(name)
//}


// 字符串的长度,用length
func main() {
	str := "hello"
	fmt.Println(utf8.RuneCountInString(str))
}
