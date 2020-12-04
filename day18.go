/*
作者   ：bjx
创建时间   ：2020/12/4  11:14 下午
文件名称   ：day18.PY
开发工具   ：GoLand
*/
package main

import "fmt"

const (
	min   = 60
	hores = min * 60
	day   = hores * 60
)

func test(S int) (d int, h int, m int) {
	d = S / day
	h = S / hores
	m = S / min
	return
}

func main() {
	fmt.Println(test(1000))
	_, h, m := test(1800)
	fmt.Println(h, m)
	d, _, _ := test(900000)
	fmt.Println(d)
}
