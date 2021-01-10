/*
作者   ：bjx
创建时间   ：2021/1/6  11:45 下午
文件名称   ：2021_01_06day.PY
开发工具   ：GoLand
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	exit := make(chan struct{}) //创建通道

	go func() {
		time.Sleep(time.Second)
		fmt.Println("hello world")
		close(exit) //关闭通道，发送信号
	}()
	fmt.Println("main...")
	<-exit //接受信号，解除阻塞
	fmt.Println("main end")
}
