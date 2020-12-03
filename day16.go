/*
作者   ：bjx
创建时间   ：2020/12/3  11:58 下午
文件名称   ：day16.PY
开发工具   ：GoLand
*/
package main

import "fmt"

func main() {

	for i := 1; i < 10; i++ {
		for j := 1; j < i; j++ {
			if j == 3 {
				//goto到标签
				goto d
			}

		}

	}
	//标签
d:
	fmt.Println("hello")
}
