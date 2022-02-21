package main

import (
	"gin_scaffold/golang_common/lib"
	"gin_scaffold/router"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	lib.InitModule("./conf/dev/",[]string{"base","mysql","redis",})
	defer lib.Destroy()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}