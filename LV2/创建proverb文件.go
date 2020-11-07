package  main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println()
	file, error := os.Create ("./proverb.txt")//创建文件
	if error != nil {
		fmt.Println(error)
	}
	deta :="don't communicate by sharing memory share memory by communicating";
	file.WriteString(deta);//读取字符
	fmt.Println(file);
	file.Close();//关闭文件
}
