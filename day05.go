/*
作者   ：bjx
创建时间   ：2020/11/22  3:39 下午
文件名称   ：day05.PY
开发工具   ：GoLand
*/
//使用指针变量获取命令行的输入信息
package main

//flag的包实现命令行参数的解析

import (
	"flag"
	"fmt"
)

//定义命令行参数
var mode = flag.String("mode", "", "Process mode")

func main() {
	//new创建指针
	str := new(string)
	//指针指向的值
	*str = "usnxuxuuwnxw"
	fmt.Println(*str)
	//解析命令行参数
	flag.Parse()
	//输出命令行参数
	fmt.Println(*mode)

}
