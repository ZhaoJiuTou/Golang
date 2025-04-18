package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	// 打印 Hello World
	fmt.Println("Hello World")
	
	// 提示用户按任意键退出
	fmt.Println("按任意键退出...")
	
	// 创建一个通道来监听 Ctrl+C 信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	
	// 等待信号或者标准输入
	go func() {
		// 读取一个字节，任何键都会导致退出
		buffer := make([]byte, 1)
		os.Stdin.Read(buffer)
		c <- os.Interrupt
	}()
	
	// 等待信号
	<-c
}