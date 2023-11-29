package aoc

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path/filepath"
)

var jar, _ = cookiejar.New(nil)
var client = http.Client{
	Jar: jar,
}

const AOC_TOKEN_NAME = "AOC_TOKEN"

// Get the input for a given year and day.
func GetInput(year Year, day Day) string {
	var directory = cacheLocation(year)
	var filename = filepath.Join(directory, fmt.Sprintf("%d.txt", day))

	if _, err := os.Stat(filename); err == nil {
		log.Println("Reading file")
		var input, err = os.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		return string(input)
	} else if errors.Is(err, os.ErrNotExist) {
		var input = fetchInput(year, day)
		os.MkdirAll(directory, os.ModePerm)
		if err := os.WriteFile(filename, []byte(input), 0666); err != nil {
			log.Fatalf("Failed to write file: %s", err)
		}

		return input
	} else {
		log.Fatal(err)
	}

	panic(nil)
}

// Location to cache input in.
func cacheLocation(year Year) string {
	return fmt.Sprintf(".input/%d/", year)
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

	var body, _ = io.ReadAll(response.Body)

	return string(body)
}

// Create the base URL for a problem.
func baseUrl(year Year, day Day) string {
	return fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
}
