/*
作者   ：bjx
创建时间   ：2021/1/10  1:55 下午
文件名称   ：2021_10day.PY
开发工具   ：GoLand
*/
package main

import "sync"

//func count()  {
//	x:=0
//	for i:=0;i<math.MaxUint32;i++{
//		x+=i
//	}
//	println(x)
//}
//
//func test11(n int)  {
//	for i:=0;i<n;i++{
//		count()
//	}
//}
func main() {
	//n:=runtime.GOMAXPROCS(0)
	//fmt.Println(n)
	//test11(n)

	//		go func(n int) {
	//			for n:=0;n<2;n++{
	//				println(n)
	//			}
	//		}(0)
	//
	//runtime.Goexit()
	//println("dd")
	//	done := make(chan struct{})
	//	c := make(chan string)
	//	go func() {
	//		s := <-c
	//		println(s)
	//		close(done)
	//	}()
	//	c<-"hi"
	//	<-done
	var we sync.WaitGroup
	//	ret := make(chan struct{})
	//	for i:=0;i<3;i++{
	//
	//		we.Add(1)
	//		println("hello",i)
	//		go func(n int) {
	//			defer  we.Done()
	//			println("id",n)
	//			<-ret
	//			println("run",n)
	//		}(i)
	//	}
	//	time.Sleep(time.Second)
	//	close(ret)
	//	we.Wait()
	//	println("end")
	n := 10
	cl := make(chan struct{})
	we.Add(1)
	go func(n int) {
		println("tt")
		defer println("hello")
		we.Done()
		defer println("8888")
		defer println("99990")
		<-cl
		n++
		println("hllo", n)
	}(n)
	close(cl)
	we.Wait()

}
