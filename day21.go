/*
作者   ：bjx
创建时间   ：2020/12/7  11:50 下午
文件名称   ：day21.PY
开发工具   ：GoLand
*/
package main

import "fmt"

//创建一个玩家生成器，输入名字，输出生成器
func play(name string) func() (string, int) {
	hp := 160
	//创建一个闭包
	return func() (string, int) {
		//返回1
		return name, hp
	}
}
func main() {

	g := play("hello")
	name, hp := g()
	fmt.Println(name, hp)

}
