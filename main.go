// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server. \n Developed by Konstantinos Vytiniotis & Katsikavelas Ioannis
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi
package main

import (
	"fmt"
	"test/web-service/controller"
	"test/web-service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	conf "rlf.com/conf"
)

func main() {

	// Programmatically set swagger info
	docs.SwaggerInfo.Title = "Retrieve LOG File (RLF) API"
	docs.SwaggerInfo.Description = "This is a basic server for retrieving K8 log files \n Created by: Vytiniotis Konstantinos & Katsikavelas Ioannis"
	docs.SwaggerInfo.Version = "0.1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	ip, port := conf.GetAPIConf()

	router := gin.Default()
	c := controller.NewController()

	v1 := router.Group("/api/v1")
	{
		log := v1.Group("/logs")
		{
			log.GET("", c.GetLogs)
		}

		search := v1.Group("/search")
		{
			search.GET("/section/:request_id", c.GetSection)
			// search.GET("/sections/:request_id", c.GetSections)
			// search.GET("/text/:request_id", c.GetSections)
		}
	}

	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(fmt.Sprintf("%s:%d", ip, port))
}
