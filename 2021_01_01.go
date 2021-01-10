/*
作者   ：bjx
创建时间   ：2021/1/1  11:24 下午
文件名称   ：2021_01_01.PY
开发工具   ：GoLand
*/
package main

func test10(a, b int) {
	//延迟调用
	defer println("dispose,,,,")
	println(a + b)
}
func main() {
	test10(10, 8)
}
