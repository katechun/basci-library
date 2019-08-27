package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFrom(reader io.Reader,num int)([]byte,error){
	p := make([]byte,num)
	n,err := reader.Read(p)
	if n >0 {
		return p[:n],nil
	}
	return p,err
}

func main() {
	//第一种读取方法
	data,err := ReadFrom(os.Stdin,11)
	if err != nil {
		return
	}
	fmt.Println(string(data))


	f,err := os.Open("E:\\tmp\\dat.txt")

	//第二种读取方法
	data1,err := ReadFrom(f,9)
	fmt.Println(string(data1))

	//第三种读取方法
	data3,err := ReadFrom(strings.NewReader("from string"),12)

	fmt.Println(string(data3))

}