package main

import (
	"encoding/json"
	"fmt"
	"gin-demo/handler"
	"gin-demo/zlog"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"runtime"
)

var log *zap.SugaredLogger

func router01() http.Handler {
	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(traceLoggerMiddleware())
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
		zlog.WithContext(c).Info("测试日志")
		zlog.WithContext(c).Info("测试日志")
		zlog.WithContext(c).Info("测试日志")
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

func traceLoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 每个请求生成的请求traceId具有全局唯一性
		u1, _ := uuid.NewV6()
		traceId := u1.String()
		zlog.NewContext(ctx, zap.String("traceId", traceId))

		// 为日志添加请求的地址以及请求参数等信息
		zlog.NewContext(ctx, zap.String("request.method", ctx.Request.Method))
		headers, _ := json.Marshal(ctx.Request.Header)
		zlog.NewContext(ctx, zap.String("request.headers", string(headers)))
		zlog.NewContext(ctx, zap.String("request.url", ctx.Request.URL.String()))

		// 将请求参数json序列化后添加进日志上下文
		if ctx.Request.Form == nil {
			ctx.Request.ParseMultipartForm(32 << 20)
		}
		form, _ := json.Marshal(ctx.Request.Form)
		zlog.NewContext(ctx, zap.String("request.params", string(form)))

		ctx.Next()
	}
}
