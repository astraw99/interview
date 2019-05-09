package main

import (
	"flag"
	"log"
	"os"
)

func usage() {
	log.Fatalf("Usage: myProgram [-s server] [-t isShowTimeStamps] <subject> \n")
}

func printMsg(message string) {
	log.Printf("Received message is : %s \n", message)
}

func main() {
	//os.Args 提供原始命令行参数访问功能。注意:切片中的第一个参数是该程序的路径，并且 os.Args[1:]保存所有的的参数。
	argsAll := os.Args
	log.Println("argsAll: ", argsAll)
	//取得对我们有意义的参数内容
	argsUseful := os.Args[1:]
	log.Println("argsUseful: ", argsUseful)

	/**
	第一个参数：设置对应的标签名，可以通过该标签名来或得对应值
	第二个参数：如果没有设置该标签，则采用这个值即该值为默认值
	第三个参数：这个参数为帮助信息，一般用于help调用展示
	*/
	var message = flag.String("s", "default message", "it's user send message[help message]")
	var showTime = flag.Bool("t", false, "Display timestamps")

	/**
	格式化log输入内容，默认为：log.LstdFlags(底层等价:Ldate | Ltime), Ldate:2017/04/01 , Ltime:16:24:36,
	Llongfile:全路径+执行文件+行数, Lshortfile:执行文件名+行数，还有几个其他不常用的，需要的话可以上官方文档查看
	这里设置的0即取消log格式化输出，输出的内容和使用fmt包下的println()格式一样
	*/
	//log.SetFlags(0)
	log.SetFlags(log.LstdFlags)
	//初始化flag包中内置的匿名Usage函数，需要赋一个函数。当flag内部发生异常会调用其内部的Usage函数，继而再调用到我们自己定义的usage函数
	flag.Usage = usage
	//所有标志都声明完成以后，调用 flag.parse() 来执行命令行解析
	flag.Parse()

	//用户没有任何参数输入则不得向下执行
	/*args := flag.Args()
	fmt.Println(args, len(args))
	if len(args) < 1 {
		usage()
	}*/

	log.Println("message:", *message, ", showTime:", *showTime)
	printMsg(*message)
}
