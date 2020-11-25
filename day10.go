/*
作者   ：bjx
创建时间   ：2020/11/25  11:03 下午
文件名称   ：day10.PY
开发工具   ：GoLand
*/
package main

import "fmt"

func main() {
	//设置元素大小
	const number = 1000
	//预分配足够多的元素切片
	srtdata := make([]int, number)

	for i := 0; i < number; i++ {
		//每一个切片位置的元素
		srtdata[i] = i
	}
	//chag
	chag := srtdata
	//c获取足够大的空间
	c := make([]int, number)
	//将srtdata复制给c
	copy(c, srtdata)

	srtdata[0] = 90

	fmt.Println(chag[0])

	fmt.Println(c[0], c[number-1])

	copy(c, srtdata[2:8])

	for j := 0; j < 10; j++ {
		fmt.Println(c[j])
	}

}
