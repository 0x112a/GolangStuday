package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solution(line string) string {
	// please write your code here
	//var a string
	s := strings.Split(line, "")
	m,i:=0,0
	for _,t := range s{
		if t=="m" {
			m++
		} else if t == "i"{
			i++
		}
	}
	for t :=0;t<m||t<i;t++ {
		line = findmi(line)
	}
	//aws+=sline[]
	return line
	// 返回处理后的结果
	// return ans
}

func findmi(line string) string{
	var aws string

	sline := strings.Split(line,"")
	//fmt.Printf("%T\t%v",line,line[2])
	for i:=0;i<len(sline);i++{
		if i==len(sline)-1 {
			aws+=sline[i]
		}else if sline[i] !="m"{
			aws += sline[i]
		}else if sline[i+1] != "i"{
			aws += sline[i]
		}else {
			i++
		}
	}
	return aws
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}