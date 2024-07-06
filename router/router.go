package router

import (
	"gin-demo/handler"
	"gin-demo/zlog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"text/template"
	"time"
)

var log *zap.SugaredLogger

func Router01(g *gin.Engine) http.Handler {
	g.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			log.Infoln("filename=", file.Filename, " filezie=", file.Size)
			c.SaveUploadedFile(file, file.Filename)
		}
	})
	g.GET("/demo/v1/get-name", func(c *gin.Context) {
		defer measureTime("/demo/v1/get-name")()
		zlog.WithContext(c)
		zlog.Logger.Info("========================测试日志1==============")
		zlog.Logger.Info("========================测试日志2==============")
		zlog.Logger.Info("========================测试日志3==============")
		handler.Loggeraaaaa()
	})
	g.GET("/hello", func(c *gin.Context) {
		t, err := template.ParseFiles("./html/hello.html")
		zlog.WithContext(c).Info("hello template...")
		if err != nil {
			zlog.Logger.Error("加载静态文件失败！")
			return
		}
		p := Person{Name: "suntong", IdNo: "111111111111"}
		t.Execute(c.Writer, p)
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

type Person struct {
	Name    string `json:"name"`
	IdNo    string `json:"idNo"`
	TraceId string `json:"traceId"`
}

func measureTime(funcName string) func() {
	start := time.Now()
	return func() {
		zlog.Logger.Info("Time taken by function is ", zap.String("funcName", funcName), zap.Duration("time", time.Since(start)))
	}
}
