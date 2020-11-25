/*
作者   ：bjx
创建时间   ：2020/11/25  10:52 下午
文件名称   ：day09.PY
开发工具   ：GoLand
*/
package main

import "fmt"

func main() {
	var num []int
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			num = append(num, i*j)
			fmt.Println(num)
		}

	}

}
