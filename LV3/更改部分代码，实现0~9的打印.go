package  main

import (
	"fmt"
	"time"
)
func goroutine(n,int,ch chan int ){
	for i:=0;i<10;i++{
		ch<-i
	}
	close(ch) //关闭通道
}
func main() {
	over := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
		time.Sleep(time.Second * 1)
		if i == 9 {
			over <- true
		}
	}
	close(over) //关闭通道
	fmt.Println("over!!!")
}