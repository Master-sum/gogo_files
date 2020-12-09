/*
作者   ：bjx
创建时间   ：2020/12/9  11:34 下午
文件名称   ：day24.PY
开发工具   ：GoLand
*/
package main

import (
	"fmt"
	"runtime"
)

//接受上下文
type pan struct {
	//所在函数
	function string
}

//保护方式允许函数
func Pro(entry func()) {
	//延时函数
	defer func() {
		//接受错误信息
		err := recover()
		//开关语句
		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime err", err)
		default:
			fmt.Println("error", err)

		}
	}()
	entry()
}
func main() {
	fmt.Println("运行时")
	Pro(func() {
		fmt.Println("手动宕机前")
		panic(&pan{
			"手动触发",
		})

		fmt.Println("手动宕机后")

	})
	Pro(func() {
		fmt.Println("赋值前")
		var a *int
		*a = 1
		fmt.Println("赋值后")
	})
	fmt.Println("运行后")
}
