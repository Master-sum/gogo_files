/*
作者   ：bjx
创建时间   ：2020/11/17  11:33 下午
文件名称   ：Slice.PY
开发工具   ：GoLand
*/
package main

import "fmt"

func main() {
	var num = make([]int, 3, 5)
	print(num)

}
func print(x []int) {
	fmt.Print("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
