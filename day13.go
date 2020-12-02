/*
作者   ：bjx
创建时间   ：2020/12/2  11:56 下午
文件名称   ：day13.PY
开发工具   ：GoLand
*/
package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println(" ")
	}

}
