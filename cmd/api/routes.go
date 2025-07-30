package main

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	
	v1 := router.Group("/v1")
	v1.POST("/create-product", )
}
