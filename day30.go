/*
作者   ：bjx
创建时间   ：2020/12/27  4:26 下午
文件名称   ：day30.PY
开发工具   ：GoLand
*/
package main

import "fmt"

func pro(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println(v, "is int")
	case string:
		fmt.Println(v, "is string")
	case bool:
		fmt.Println(v, "is bool")
	}
}

func main() {
	pro("hello")
	pro(12)
	pro(true)

}
