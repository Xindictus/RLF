package controller

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
)

// Endpoint used to return a subset of the log files with just the section of the request id in question
func (c *Controller) GetSection(ctx *gin.Context) {
	// Retrieve param
	request_id := ctx.Param("request_id")

	fmt.Printf("request_id: %s\n", request_id)

	// Open file
	// TODO: Combine with logs.go in order to parser dynamically through files
	f, err := os.Open("./test_files/test01.log")
	if err != nil {
		fmt.Println(err)
	}

	// Init variables
	var tot_bytes_start, section_start, lines int64 = 0, 0, 0
	var section_match string

	// Regex patterns
	reg_sect_start := "^#{3}\\sGENERATION\\sREPORT\\sSTART\\s#{3}$"
	reg_sect_end := "^#{3}\\sGENERATION\\sREPORT\\sEND\\s#{3}$"

	// Read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines += 1
		line := scanner.Text()

		// Mark section start
		matched, _ := regexp.MatchString(reg_sect_start, line)

		if matched {
			fmt.Println("Section start found in line ", lines)
			section_start = tot_bytes_start
		}

		// Check if request id exists in text
		matched, _ = regexp.MatchString(request_id, line)

		if matched {
			fmt.Println("Request ID found in line ", lines)
			break
		}

		// Increment bytes read count
		tot_bytes_start += int64(len(line)) + 1
	}

	// Move file cursor to the start of last marked section
	f.Seek(section_start, 0)
	scanner = bufio.NewScanner(f)

	// Build string containing the section
	for scanner.Scan() {
		line := scanner.Text()

		section_match += "\n" + line

		matched, _ := regexp.MatchString(reg_sect_end, line)

		if matched {
			fmt.Println("Section end found")
			break
		}

	}

	f.Close()

	ctx.String(http.StatusOK, "%t", section_match)
}

// Will be used to return multiple sections from log files
// func (c *Controller) GetSections(ctx *gin.Context) {}
