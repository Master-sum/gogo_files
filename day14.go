/*
作者   ：bjx
创建时间   ：2020/12/3  12:12 上午
文件名称   ：day14.PY
开发工具   ：GoLand
*/
package main

import "fmt"

func main() {
	for key, value := range []int{1, 2, 3, 4} {
		fmt.Println(key, value)

	}
	s := "string"
	switch s {
	case "hello":
		fmt.Println("world")
	case "string":
		fmt.Println("12")
	}
}
