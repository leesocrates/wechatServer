package main

import (
	"wechatServer/src/config"
	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadConfig()

	r := gin.New()
	initRouters(r)
	r.Run("8888")
}

func initRouters(r *gin.Engine){

}