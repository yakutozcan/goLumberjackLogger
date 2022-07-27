package main

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM)

	var i = 0
	log.SetOutput(&lumberjack.Logger{
		Filename:   "log/logfile.log",
		MaxSize:    1, // megabytes
		MaxBackups: 10,
		MaxAge:     90,   //days
		Compress:   true, // disabled by default
	})
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("example log start ", "v0.5.0")

	go func() {
		for {
			i++
			log.Println("counterOfNumber", i)
			log.Println("userLogin", "Lorem Ipsum is simply dummy text of the printing and typesetting industry.")
			//log rotation
			time.Sleep(time.Millisecond * 20)
		}
	}()
	<-gracefulShutdown
	log.Println("gracefulShutdown", "done")
}
