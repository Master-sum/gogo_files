/*
作者   ：bjx
创建时间   ：2020/12/7  11:57 下午
文件名称   ：day22.PY
开发工具   ：GoLand
*/
package main

import (
	"bytes"
	"fmt"
)

func joinStr(slist ...string) string {
	//定义一个字节缓冲，快速地连接字符串
	var b bytes.Buffer
	for _, s := range slist {
		//将遍历出地字符串写入字节数组
		b.WriteString(s)
	}
	//将字节数组转化字符串
	return b.String()

}
func main() {
	fmt.Println(joinStr("bai", "jin", "xing"))
}
