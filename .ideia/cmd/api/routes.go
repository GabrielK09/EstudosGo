package main

<<<<<<< HEAD:.ideia/cmd/api/routes.go
func Routes() {

=======
import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	
	v1 := router.Group("/v1")
	v1.POST("/create-product", )
>>>>>>> 42f08843a2c6a26d99224e2206c8b142a50648e0:cmd/api/routes.go
}
