package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//声明并初始化带缓冲的读取器
	//准备从标准输入读取内容
	inputrueader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")
	//以\n为分隔符读取一段内容
	input, err := inputrueader.ReadString('\n')
	if err != nil {
		fmt.Printf("Found an error")
	} else {
		//对input进行切片操作，去掉内容中最后一个字节\n
		input = input[:len(input)-1]
		fmt.Printf("Hello,%s!\n", input)
	}
}
