/*
作者   ：bjx
创建时间   ：2020/11/29  11:07 下午
文件名称   ：day12.PY
开发工具   ：GoLand
*/
package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack(79)
	l.PushFront("test")
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
