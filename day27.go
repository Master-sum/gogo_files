/*
作者   ：bjx
创建时间   ：2020/12/13  10:36 下午
文件名称   ：day27.PY
开发工具   ：GoLand
*/
package main

type LogWriter interface {
	Writer(data interface{}) error
}
type Logger struct {
	writerlist []LogWriter
}

func (l *Logger) Reg(writer LogWriter) {
	l.writerlist = append(l.writerlist, writer)

}
func (l *Logger) Log(data interface{}) {
	for _, writer := range l.writerlist {
		writer.Writer(data)
	}
}
func NewLogger() *Logger {
	return &Logger{}
}
func main() {
	s := new(Logger)
	s.Log("dddd")
}
