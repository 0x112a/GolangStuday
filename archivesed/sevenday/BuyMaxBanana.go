//买香蕉
//描述
//我是一个爱吃香蕉的强迫症。今天我要去水果店论筐买香蕉。 现在水果店有好多筐香蕉，我的要求是买回来的每一筐里必须有相同数量的香蕉。
//
//为了实现这个目标，你可以每次做两件事情。
//1）放弃框里的一部分香蕉 2）连筐带香蕉放弃一整筐
//
//我想知道我能得到最多多少香蕉。
//
//输入
//以空格分割的多个正整数，每个正整数表示一筐香蕉的总香蕉数
//
//输出
//最多能得到的香蕉数
//
//输入样例
//1 2 3
//5 0 29 14
//复制样例
//输出样例
//4
//29
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(line string) string  {
	var aws int
	var N []int
	n := strings.Split(line," ")
	for i :=0 ;i < len(n);i++{
		s, _ := strconv.Atoi(n[i])
		N = append(N,s)
	}
	sort.Ints(N)
	for i := 0 ; i < len(N);i++{
		if N[i]*(len(N)-i) > aws{
			aws = N[i]*(len(N)-i)
		}
	}

	//fmt.Printf("%T\t%v",N,N)

	return strconv.Itoa(aws)
}

func main(){
	r := bufio.NewReaderSize(os.Stdin,20480)
	for line,_,err := r.ReadLine();err==nil;line,_,err = r.ReadLine(){
		fmt.Println(solution(string(line)))
	}
}
