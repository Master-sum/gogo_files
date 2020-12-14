/*
作者   ：bjx
创建时间   ：2020/12/14  11:39 下午
文件名称   ：day28.PY
开发工具   ：GoLand
*/
package main

import (
	"fmt"
	"sort"
)

type Mylist []string

func (m Mylist) Len() int {
	return len(m)
}
func (m Mylist) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m Mylist) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func main() {
	name := Mylist{
		"2.999",
		"3.nxx",
		"1.nxunud",
	}
	sort.Sort(name)
	for _, v := range name {
		fmt.Printf("%s\n", v)
	}

}
