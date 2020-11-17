/*
作者   ：bjx
创建时间   ：2020/11/17  11:42 下午
文件名称   ：切片test.PY
开发工具   ：GoLand
*/
package main

import "fmt"

func main() {
	var num = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Print(num)
	fmt.Println("----------------")
	fmt.Println(num[1:5])
	num = append(num, 10)
	fmt.Println(num)

}
