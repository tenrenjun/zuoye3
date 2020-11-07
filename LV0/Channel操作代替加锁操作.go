package  main

import (
	"fmt"
	"time"
)
//使用内置函数make创建通道ch
//ch := make(chan int)//创建ch接受int类型数据,不进行初始化

func factorial(n int, ch chan int ) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
		ch <- res
	}
	close(ch)
}

func main(){
	var ch =make(chan int ,20 )
	go factorial(cap(ch), ch) //cap(c)即为c的容量大小,调用函数cap()查看通道c的数据
	// 不断读取channel数据，直到channel被显式关闭
	time.Sleep(time.Second * 2)
	for i := range ch {
		//遍历通道c
		fmt.Println(i)//打印通道c的每个元素
	}
}
