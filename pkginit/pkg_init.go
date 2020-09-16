package main

import (
	"fmt"
	"runtime"
)

//代码初始化function
func init() {
	//格式化的打印
	fmt.Printf("Map:%v\n", m)
	//通过调用 runtime 包代码获取当前机器的的操作系统和计算机架构
	//然后通过fmt包的Sprintf方法进行格式化字符串生成并赋值给变量info
	info = fmt.Sprintf("OS: %s,Arch:%s", runtime.GOOS, runtime.GOARCH)
}

var m = map[int]string{1: "A", 2: "B", 3: "C"}

var info string

func main() {
	fmt.Println(info)
}
