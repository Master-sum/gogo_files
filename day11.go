/*
作者   ：bjx
创建时间   ：2020/11/29  11:03 下午
文件名称   ：day11.PY
开发工具   ：GoLand
*/
package main

import "fmt"

func main() {
	const A, B = "hello ", " world"
	fmt.Printf("%s", test1(A))
	fmt.Printf("%s", test1(B))

}
func test1(s string) string {
	return s

}
