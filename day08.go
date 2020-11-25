/*
作者   ：bjx
创建时间   ：2020/11/22  4:55 下午
文件名称   ：day08.PY
开发工具   ：GoLand
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b string
	for i := 1; i < len(os.Args); i++ {
		a += b + os.Args[i]
		b = " "
	}
	fmt.Println(a)
}
