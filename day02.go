/*
作者   ：bjx
创建时间   ：2020/11/22  3:04 下午
文件名称   ：day02.PY
开发工具   ：GoLand
*/
//对于&和*的指针取地址和取值进行操作
package main

import "fmt"

func main() {
	var test = "hello test string"
	//这里将test的指针地址赋值给pr
	pr := &test
	//输出pr的类型
	fmt.Printf("pr:Teyp %T\n", pr)
	//输出pr的地址
	fmt.Printf("pr address:%p\n", pr)
	//将pr进行赋值end
	end := *pr
	//输出类型
	fmt.Printf("end %T\n", end)
	//输出值
	fmt.Printf("end_value:%s\n", end)
}
