package main

import (
	"SportAnalyse/models"
	v1 "SportAnalyse/routers/api/v1"
	"SportAnalyse/setting"
	"github.com/gin-gonic/gin"
)

func main() {

	setting.Setup()
	models.Setup()


	r := gin.Default()
	models.Setup()

	r.GET("/info", v1.GetInfos)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
