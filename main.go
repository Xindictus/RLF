package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// const LOG_FILES_PATH = "/home/ioannis/Desktop"
// const filename = "DJI_0709.JPG"

type params struct {
    TypeOfReport	string  `json:"reportType"`
	DateTime		string	`json:datetime`
}

func getLogs(c *gin.Context) {
	report_type := c.DefaultQuery("reportType", "all") // shortcut for c.Request.URL.Query().Get("reportType")
	datetime := c.DefaultQuery("datetime", time.Now().String())

	p := params {
		TypeOfReport: report_type,
		DateTime: datetime,
	}

	// TODO Function that searches for the file based on the previosu parameters
	// targetPath := filepath.Join(LOG_FILES_PATH, filename)
	// c.FileAttachment(targetPath, filename)
	c.IndentedJSON(http.StatusOK, p)
}


func main() {
	router := gin.Default()
	router.GET("/logs", getLogs)

	router.Run("localhost:8080")
}