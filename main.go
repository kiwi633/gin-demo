package main

import (
	"fmt"
	"gin-demo/router"
	"gin-demo/zlog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"runtime"
)

var log *zap.SugaredLogger

func getGoroutineID() int {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	var id int
	fmt.Sscanf(string(b), "goroutine %d ", &id)
	return id
}

func main() {
	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(zlog.TraceLoggerMiddleware())
	g.MaxMultipartMemory = 8 << 20
	router.Router01(g)
}
