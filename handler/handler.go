package handler

import (
	"gin-demo/zlog"
	"github.com/gin-gonic/gin"
	"time"
)

func FindPerson(c *gin.Context) {
	zlog.WithContext(c).Info("FindPerson")
	time.Sleep(5000)
	ccc(c)
}

func ccc(c *gin.Context) {
	zlog.WithContext(c).Info("FindPerson ccccc")
	ddd(c)
}
func ddd(c *gin.Context) {
	zlog.WithContext(c).Info("FindPerson ddd")
	eee(c)
}
func eee(c *gin.Context) {
	zlog.WithContext(c).Info("FindPerson eee")
	zlog.WithContext(c).Info("FindPerson eee")
	zlog.WithContext(c).Info("FindPerson eee")
	zlog.WithContext(c).Info("FindPerson eee")
	zlog.WithContext(c).Info("FindPerson eee")
	zlog.WithContext(c).Info("FindPerson eee")
}
