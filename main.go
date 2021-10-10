package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// const LOG_FILES_PATH = "/home/ioannis/Desktop"
// const filename = "DJI_0709.JPG"

type params struct {
    TypeOfReport	string  `json:"reportType"`
	RequestId		string	`json:"requestId"`
	DateTime		string	`json:"datetime"`
}

type  detail struct {
    Detail	string  `json:"detail"`
}

func getLogs(c *gin.Context) {
	report_type := c.DefaultQuery("reportType", "all") // shortcut for c.Request.URL.Query().Get("reportType")
	request_id	:= c.DefaultQuery("requestId", "")
	datetime := c.DefaultQuery("datetime", time.Now().String())

	detail := detail {
		Detail: "Request Id is missing, please include a requestId",
	}
	if request_id == "" {
		fmt.Printf("Mpika")
		c.IndentedJSON(http.StatusBadRequest, detail)
		return
	}

	p := params {
		TypeOfReport: report_type,
		RequestId: request_id,
		DateTime: datetime,
	}
	fmt.Printf("kai Vgika")

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