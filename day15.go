/*
作者   ：bjx
创建时间   ：2020/12/3  11:52 下午
文件名称   ：day15.PY
开发工具   ：GoLand
*/
package main

import "fmt"

func main() {
	s := "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "ss":
		fmt.Println("world")

	}
}
