package  main

import (
	"fmt"
	"time"
)

func goroutine1(n int,chan1 chan int ){
	for i:=0;i<100;i++{
		 x :=i
		chan1<-x
	}
	close(chan1) //关闭通道
}
func goroutine2(n,int,chan2 chan int ){
	for i:=0;i<100;i=i+2{
		chan2<-i
	}
	close(chan2) //关闭通道
}
func main() {
	var chan1= make(chan int, 100)//创建通道chan1为int类型而且定义大小为100；
	go  goroutine1(cap(chan1),chan1) //cap(c)即为chan1的容量大小,调用函数cap()查看通道c的数据
	// 不断读取channel数据，直到channel被显式关闭
	time.Sleep(time.Second * 2)
	for i := range chan1 {//遍历通道chan1
		fmt.Println(i)//打印通道chan1的每个元素
	}
	var chan2= make(chan int, 100)//创建通道chan2为int类型而且定义大小为10；
	go  goroutine1(cap(chan2),chan2) //cap(chan2)即为chan2的容量大小,调用函数cap()查看通道c的数据
	// 不断读取channel数据，直到channel被显式关闭
	time.Sleep(time.Second * 4)
	for i := range chan2 {//遍历通道chan2
		fmt.Println(i)//打印通道chan2的每个元素
	}
}
