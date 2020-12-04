/*
作者   ：bjx
创建时间   ：2020/12/4  12:06 上午
文件名称   ：day17.PY
开发工具   ：GoLand
*/
package main

import "fmt"

func main() {
	err := firstCheckError()
	if err != nil {
		goto onExit
	}
	err := secondCheckError()
	if err != nil {
		goto onExit
	}
	fmt.Println("dd")
	return
onExit:
	fmt.Println(err)
	exitProcess()
}
