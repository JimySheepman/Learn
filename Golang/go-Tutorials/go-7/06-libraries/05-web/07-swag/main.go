package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"swag/api"
	_ "swag/docs"
)

func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/testapi/get-string-by-int/", api.GetStringByInt)
	r.GET("/testapi/get-struct-array-by-string/", api.GetStructArrayByString)
	r.POST("/testapi/upload", api.Upload)

	log.Println("open http://localhost:3000/swagger/index.html")

	r.Run(":3000")
}
