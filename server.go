package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Job struct {
	JobType     string `json:"type"`
	Company     string `json:"company"`
	Title       string `json:"title"`
	Location    string `json:"location"`
	Description string `json:"description"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Logger middleware
	e.Use(middleware.Logger())

	// Calling GitHub jobs API
	response, _ := http.Get("https://jobs.github.com/positions.json?description=python&full_time=true&location=sf")
	// Get body of the response
	data, _ := ioutil.ReadAll(response.Body)
	// Instantiate a new job struct
	var job []Job
	// Unmarshal data into a pointer to the job struct
	json.Unmarshal(data, &job)
	// Create HTML string
	jobPosting := "Company: <br>" + job[0].Company + "<br><br> Description: " + job[0].Description

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, jobPosting)
	})

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
