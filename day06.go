/*
作者   ：bjx
创建时间   ：2020/11/22  3:50 下午
文件名称   ：day06.PY
开发工具   ：GoLand
*/
//变量逃逸分析
package main

import "fmt"

//定义函数参数和返回值
func dummy(b int) int {
	var c int
	c = b
	return c
}

//定义一个空函数
func void() {

}

//主函数调用
func main() {
	//声明a
	var a int
	//调用空函数
	void()
	//打印a变量和dummy函数
	fmt.Println(a, dummy(0))
}

/*
 go run -gcflags "-m -l" day06.go    //-gcflags是编译参数 -m表示进行内存分配分析，-l表示避免程序内联，也就是避免进行程序化优化
# command-line-arguments
./day06.go:27:13: ... argument does not escape
./day06.go:27:13: a escapes to heap//变量逃逸到堆里面
./day06.go:27:21: dummy(0) escapes to heap
0 0

*/
