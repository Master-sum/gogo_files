/*
作者   ：bjx
创建时间   ：2021/1/6  11:52 下午
文件名称   ：2021_01_06_01day.PY
开发工具   ：GoLand
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var va sync.WaitGroup
	for i := 0; i < 10; i++ {
		va.Add(1) //累加
		go func(id int) {
			defer va.Done() //延时递减
			time.Sleep(time.Second)
			println(id)
		}(i)
	}
	fmt.Println("hello")
	va.Wait()
	fmt.Print("exit")
}
