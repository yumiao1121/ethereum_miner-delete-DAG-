package main

import (
	"TCP_ceshi/TCP_Client/handle"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
)

func main() {
	times := 1 //开启节点数量
	logInit()
	exit := make(chan bool, 1)
	for i := 0; i < times; i++ {
		go handle.NewWorker(i, exit) //开启节点线程
	}
	i := 0
	for {
		<-exit
		i++
		if i > times-1 {
			break
		}
	}
}

func logInit() {
	var logFileName = flag.String("log", "cServer.log", "Log file name")
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
		os.Exit(1)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate)
}

//NewTcpGetWork()
//NewTcpLogin()
//NewTcpSubmitWork()
