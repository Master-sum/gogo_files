/*
作者   ：bjx
创建时间   ：2020/11/22  3:19 下午
文件名称   ：day03.PY
开发工具   ：GoLand
*/
//使用指针修改值
package main

//标准输出
import "fmt"

func change(a, b *int) {
	//将a的值赋值给t
	t := *a
	//将b的值赋值给a
	*a = *b
	//将t的值赋值给b
	*b = t
}

//主函数
func main() {
	//准备两个值
	x, y := 1, 2
	//将x,y的指针进行传递
	change(&x, &y)

	fmt.Print(x, y)
}
