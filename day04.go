/*
作者   ：bjx
创建时间   ：2020/11/22  3:33 下午
文件名称   ：day04.PY
开发工具   ：GoLand
*/
//这里就是将两个的地址进行了交换，但是值没有进行交换
package main

import "fmt"

func change1(a, b *int) {
	a, b = b, a
}

func main() {
	x, y := 1, 2
	change1(&x, &y)
	fmt.Print(x, y)
}
