package aoc

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
)

var jar, _ = cookiejar.New(nil)
var client = http.Client{
	Jar: jar,
}

const AOC_TOKEN_NAME = "AOC_TOKEN"

// Get the input for a given year and day.
func GetInput(year Year, day Day) string {
	return fetchInput(year, day)
}

// Fetch the input for a problem, given a year and day.
func fetchInput(year Year, day Day) string {
	var token, has_token = os.LookupEnv(AOC_TOKEN_NAME)
	if !has_token {
		log.Panicf("Missing `%s` in environment", AOC_TOKEN_NAME)
	}

	var url = fmt.Sprintf("%s/input", baseUrl(year, day))
	var request, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Failed to create request %s", err)
	}
	request.AddCookie(&http.Cookie{Name: "session", Value: token})

	var response, _ = client.Do(request)
	defer response.Body.Close()

	fmt.Println(response.StatusCode)
	var body, _ = io.ReadAll(response.Body)

	return string(body)
}

// Create the base URL for a problem.
func baseUrl(year Year, day Day) string {
	return fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
}
