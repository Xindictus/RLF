package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// let's get to hacking your nice pathS
const LOG_FILES_PATH = "/home/ioannis/Desktop/"

// recursively walk through a directory and return the paths to all files whose name matches the given pattern
func WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

type detail struct {
	Detail string `json:"detail"`
}

func searchFile(request_id string, filepath string) bool {
	// read the whole file at once
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	s := string(b)
	// //check whether s contains substring text
	// fmt.Println(strings.Contains(s, request_id))
	request_id_exists := strings.Contains(s, request_id)
	return request_id_exists
}

// Search for log file godoc
// @Summary Search for log file
// @Description Search for log file by request_id
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param requestId query string false "Log search by requestId"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /logs [get]
func (c *Controller) GetLogs(ctx *gin.Context) {
	// report_type := c.DefaultQuery("reportType", "all") // shortcut for c.Request.URL.Query().Get("reportType")
	request_id := ctx.DefaultQuery("requestId", "")
	if request_id == "" {
		detail := detail{
			Detail: "Request Id is missing, please include a requestId",
		}
		ctx.IndentedJSON(http.StatusBadRequest, detail)
		return
	}

	files, err := WalkMatch(LOG_FILES_PATH, "*.txt")
	if err != nil {
		fmt.Print(err)
	}

	for _, item := range files {
		request_id_exists := searchFile(request_id, item)
		if request_id_exists {
			last_element_of_path := filepath.Base(item)
			ctx.FileAttachment(item, last_element_of_path)
		}
		fmt.Println("Result", item, request_id_exists)
	}

	// p := params {
	// 	TypeOfReport: report_type,
	// 	RequestId: request_id,
	// 	DateTime: datetime,
	// }
	detail := detail{
		Detail: "Unfortunatelly no log files found with this specific requestId: " + request_id,
	}
	ctx.IndentedJSON(http.StatusNotFound, detail)
}
