package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//从标准输入读入数据
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input you name:")
	//读入数据直到碰到\n
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("An error occured:%s\n", err)
		//异常退出
		os.Exit(1)
	} else {
		name := input[:len(input)-1]
		fmt.Printf("Hello, %s! what can I help you?\n", name)
	}
	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An err occured: %s\n", err)
			continue
		}
		input = input[:len(input)-1]
		//全部转换成小写
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			fmt.Println("Sorry,I didn`t catch you.")

		}
	}
}
