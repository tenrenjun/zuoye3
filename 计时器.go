package main
import (

"fmt"

"time"

)


func testTimer() {

	go func() {

		fmt.Println("芜湖！起飞")

	}()


}
func timer1() {
	for {
		now := time.Now()  //获取当前时间，放到now里面，要给next用
		next := now.Add(time.Hour * 24) //通过now偏移24小时
		next = time.Date(next.Year(), next.Month(), next.Day(), 8, 0, 0, 0, next.Location()) //获取下一个早上八点的日期
		t := time.NewTimer(next.Sub(now))//计算当前时间到凌晨的时间间隔，设置一个定时器
		<-t.C
		fmt.Print("早安，打工人！",time.Now())
		//以下为定时执行的操作

	}

}

func timer2() {

	for {
		now := time.Now()  //获取当前时间，放到now里面，要给next用
		next := now.Add(time.Hour * 24) //通过now偏移24小时
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location()) //获取下一个凌晨的日期
		t := time.NewTimer(next.Sub(now))//计算当前时间到凌晨的时间间隔，设置一个定时器
		<-t.C
		fmt.Print("没有困难的工作，只有勇敢的打工人！",time.Now())
		//以下为定时执行的操作

	}

}

func timer3() {

	timer3 := time.NewTicker(1 * time.Hour)

	for {

		select {

		case <-timer3.C:

			testTimer()

		}

	}

}

func main() {

	go
		timer1()
	timer3()
     timer2()

}