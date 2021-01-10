/*
作者   ：bjx
创建时间   ：2020/12/27  10:08 下午
文件名称   ：day01.PY
开发工具   ：GoLand
*/
package files

import "fmt"

type Alipay struct {
}

func (b *Alipay) Can() {

}

type sun interface {
	Can()
}

func print(v interface{}) {
	switch v.(type) {
	case sun:
		fmt.Printf("%T is", v)
	}
}
func main() {
	print(new(Alipay))
}
