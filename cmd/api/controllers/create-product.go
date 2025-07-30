package controllers

import (
	"github.com/gin-gonic/gin"
	"g2l.com/internal/use-cases"
)

type createProductInput {
	Name string `json:"name" bindiging:"required"`
}

func CreateProduct(c *gin.Context) {
	useCase := use
}