/*
作者   ：bjx
创建时间   ：2020/11/22  4:39 下午
文件名称   ：day07.PY
开发工具   ：GoLand
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	tipl := "hello world"
	fmt.Println(len(tipl))
	tipl02 := "你好"
	//按字节进行计算
	fmt.Println(len(tipl02))
	//计算中文
	fmt.Println(utf8.RuneCountInString(tipl02))
	//遍历字符串，按照ascii
	theme := "你好 hello world"
	//不能识别中文
	for i := 0; i < len(theme); i++ {
		fmt.Printf("ascii:%c %d\n", theme[i], theme[i])
	}
	//遍历字符串，按照unicode
	for _, s := range theme {
		fmt.Printf("ascii:%c %d\n", s, s)
	}
	for i, ch := range theme {
		fmt.Println(i, ch, string(ch))
	}
}

/*
按照ascii遍历直接使用下标
按照unicode遍历使用range
*/
