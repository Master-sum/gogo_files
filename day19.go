/*
作者   ：bjx
创建时间   ：2020/12/6  9:05 下午
文件名称   ：day19.PY
开发工具   ：GoLand
*/
package main

import (
	"fmt"
	"strings"
)

//func fire()  {
//	fmt.Println("这是个函数")
//}

func main() {
	////定义一个函数
	//var f  func()
	//f = fire
	//f()
	list := []string{
		"go ssssyy",
		"go nxsnxs",
		"go bbbsshan",
		"go  nxu",
	}
	ch := []func(string2 string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}
	StringProccess(list, ch)
	for _, str := range list {
		fmt.Println(str)
	}
}

func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}
func StringProccess(list []string, ch []func(str string) string) {
	for index, str := range list {
		res := str
		for _, proc := range ch {
			res = proc(res)
		}
		list[index] = res

	}

}
