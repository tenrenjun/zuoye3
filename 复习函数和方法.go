package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
//字符串处理

	var s string = "hello go"
	//判断字符串s是否包含子串
	r := strings.Contains(s, "go")
	fmt.Println(r) //true

	var s1 []string = []string{"1", "2", "3", "4"}
	//将一系列字符串连接为一个字符串，之间用sep来分隔
	s2 := strings.Join(s1, ",")
	fmt.Println(s2) //1,2,3,4

	//子串sep在字符串s中第一次出现的位置，不存在则返回-1
	n1 := strings.Index(s, "go")
	fmt.Println(n1) //6

	//返回count个s串联的字符串
	s3 := strings.Repeat(s, 2)
	fmt.Println(s3) //hello gohello go

	//返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串
	s4 := strings.Replace(s, "o", "e", -1)
	fmt.Println(s4) //helle ge

	//字符串切割，如果后一个参数为空，则按字符切割，返回字符串切片
	s5 := strings.Split(s, " ")
	fmt.Println(s5) //[hello go]

	//切除字符串两端的某类字符串
	s6 := strings.Trim(s, "o")
	fmt.Println(s6) //hello g

	//去除字符串的空格符，并且按照空格分割返回slice
	s7 := " hello go "
	s8 := strings.Fields(s7)
	fmt.Println(s8) //[hello go]

//类型转换&strconv

	/*****Format系列******/
	// bool转字符串
	fmt.Println(strconv.FormatBool(true))
	//'f'指打印格式以小数方式，3：指小数位数， 64：指float64处理
	fmt.Println(strconv.FormatFloat(2.12, 'f', 3, 64))
	//整数转字符串
	fmt.Println(strconv.Itoa(19))

	/*****Parse系列******/
	//字符串转bool
	flag, _ := strconv.ParseBool("true")
	fmt.Println(flag)
	//字符串转浮点
	float, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println(float)
	//字符串转整型
	i, _ := strconv.Atoi("123")
	fmt.Println(i)

	/*****Append系列******/
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	// 10:指十进制
	slice = strconv.AppendInt(slice, 123, 10)
	slice = strconv.AppendFloat(slice, 3.14, 'f', 2, 64)
	slice = strconv.AppendQuote(slice, "hello go")
	fmt.Println(string(slice))

//map

	var dict map[string]string //定义dict为map类型
	dict = make(map[string]string) //让dict可编辑

	dict["a"] = "打工人" //加值1
	dict["b"] = "996" //加值2

	value, ok := dict["a"] //ok是看当前key是否存在返回布尔，value返回对应key的值
	if ok {
		fmt.Println("存在",value)
	}else{
		fmt.Println("不存在")
	}

	for key := range dict{ //取map中的值
		fmt.Println(key,"**********",dict[key])
	}

	delete(dict,"a") //删除集合元素

	for key,value := range dict{ //取map中的值
		fmt.Println(key,"**********",value)
	}
}