// code.mi.com #64
//描述
//有 500 个小孩围成一圈，编号从 1 到 500，从第一个开始报数：1，2，3，1，2，3，1，2，3，……每次报到 3 的小孩退出。问第 n 个被淘汰的小孩，在最开始 500 人里是的编号是几？
//输入
//正整数 N，表示要计算的为第 N 个淘汰的小孩的编号，0 < N <= 500
//输出
//第 N 个淘汰的小孩的编号
//输入样例
//1
//2
//206
//311
//输出样例
//3
//6
//176
//223

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solution(line string) string {
	N, _ := strconv.Atoi(line)
	var aws string
	var cnt int
	var index [1500]int
	for i:=0;i<500;i++ {
		index[i] = i+1
	}
	i := 500
	for j:=1;j<=i;j++{
		if j%3!=0{
			index[i]=index[j-1]
			//fmt.Println(index[i])
			i++
		}else {
			cnt++
			if cnt == N {
				aws = strconv.Itoa(index[j-1])
				break
			}
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
//另一种神奇的方法
//
//package main
//
//import (
//"bufio"
//"fmt"
//"os"
//"strconv"
//)
//
//func solution(line string) string {
//	// please write your code here
//	n,_:=strconv.Atoi(line)
//	shit:=make([]int,501)
//	index:=1
//	killThird:= func() {
//		count:=0
//		for {
//			if shit[index]==0{
//				count++
//				if count==3{
//					shit[index]=1
//					return
//				}
//			}
//			if index==500{
//				index=1
//			}else{
//				index++
//			}
//		}
//	}
//	for i:=0;i<n;i++{
//		killThird()
//	}
//	// 返回处理后的结果
//	return strconv.Itoa(index)
//}
//
//func main() {
//	r := bufio.NewReaderSize(os.Stdin, 20480)
//	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
//		fmt.Println(solution(string(line)))
//	}
//}
