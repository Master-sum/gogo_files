/*
作者   ：bjx
创建时间   ：2020/12/8  12:10 上午
文件名称   ：day23.PY
开发工具   ：GoLand
*/
package main

import "fmt"

func main() {
	//d := []int{1,2,3}
	//
	//for i:=0;i<10;i++{
	//	fmt.Println(i)
	//}
	//for _,j:=range d{
	//	fmt.Println(j)
	//}
	fmt.Println("hello")
	//延时进栈，defer相当于进栈操作
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	fmt.Println("world")
}
