package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/temur-shamshidinov/Api_gateway/genproto/content_service"
	"github.com/temur-shamshidinov/Api_gateway/services"
)

func Api() {

	router := gin.Default()

	services := services.Services()

	api := router.Group("api")

	api.POST("/create", func(ctx *gin.Context) {

		var reqBody content_service.CreateContentReq

		ctx.BindJSON(&reqBody)

		resp, err := services.CreateContent(ctx, &reqBody)
		if err != nil {
			log.Println(err)
			return
		}

		ctx.JSON(201,resp)
	})

	router.Run(":8080")
}
