package main

import (
	"fmt"
	"gin-demo/handler"
	"gin-demo/zlog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"runtime"
)

var log *zap.SugaredLogger

func router01() http.Handler {
	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(zlog.TraceLoggerMiddleware())
	g.MaxMultipartMemory = 8 << 20
	g.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			log.Infoln("filename=", file.Filename, " filezie=", file.Size)
			c.SaveUploadedFile(file, file.Filename)
		}
	})
	g.GET("/demo/v1/get-name", func(c *gin.Context) {
		zlog.WithContext(c)
		zlog.Logger.Info("========================测试日志1==============")
		zlog.Logger.Info("========================测试日志2==============")
		zlog.Logger.Info("========================测试日志3==============")
		handler.Loggeraaaaa()
	})
	g.POST("/get-person", func(c *gin.Context) {
		zlog.WithContext(c).Info("get-person1测试日志")
		zlog.WithContext(c).Info("get-person2测试日志")
		zlog.WithContext(c).Info("get-person3测试日志")
		handler.FindPerson(c)
	})
	g.Run(":8089")
	return g
}
func getGoroutineID() int {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	var id int
	fmt.Sscanf(string(b), "goroutine %d ", &id)
	return id
}

type Person struct {
	Name string `json:"name"`
	IdNo string `json:"idNo"`
}

func main() {
	router01()
}
