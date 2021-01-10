/*
作者   ：bjx
创建时间   ：2020/12/27  10:14 下午
文件名称   ：day31.PY
开发工具   ：GoLand
*/
package main

import "fmt"

type Alipay struct {
}

func (b *Alipay) Can() {
	fmt.Println("sorry")
}

type sun interface {
	Can()
}

func printo(v interface{}) {
	switch v.(type) {
	case sun:
		fmt.Printf("%T is\n", v)
	}
}
func main() {
	a := new(Alipay)
	printo(a)
	a.Can()
}
