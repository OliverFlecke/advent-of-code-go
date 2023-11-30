package aoc

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var jar, _ = cookiejar.New(nil)
var client = http.Client{
	Jar:       jar,
	Transport: newAddHeaderTransport(nil),
}

const AOC_TOKEN_NAME = "AOC_TOKEN"

// Get the input for a given year and day.
func GetInput(year Year, day Day) string {
	directory := cacheLocation(year)
	filename := filepath.Join(directory, fmt.Sprintf("%d.txt", day))

	if _, err := os.Stat(filename); err == nil {
		// log.Println("Reading file")
		input, err := os.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		return string(input)
	} else if errors.Is(err, os.ErrNotExist) {
		input := fetchInput(year, day)
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

// Submit an answer for a problem.
func SubmitAnswer(year Year, day Day, level Level, answer string) SubmissionResult {
	answerUrl := fmt.Sprintf("%s/answer", baseUrl(year, day))

	form := url.Values{}
	form.Add("level", string(level))
	form.Add("answer", answer)
	request, _ := http.NewRequest(http.MethodPost, answerUrl, strings.NewReader(form.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	return parseSubmissionResponse(string(body))
}

func parseSubmissionResponse(body string) SubmissionResult {
	r, _ := regexp.Compile("(?s)<main>.*</main>")
	main := r.FindString(body)

	if strings.Contains(body, "That's the right anwswer") {
		return Correct
	}
	if strings.Contains(body, "already complete it") {
		return AlreadyCompleted
	}
	if strings.Contains(body, "answer too recently") {
		return TooRecent
	}
	if strings.Contains(body, "not the right answer") {
		fmt.Printf(main)
		return Incorrect
	}

	fmt.Printf(body)
	panic("Could not parse body")
}

// Location to cache input in.
func cacheLocation(year Year) string {
	return fmt.Sprintf(".input/%d/", year)
}

// Fetch the input for a problem, given a year and day.
func fetchInput(year Year, day Day) string {
	log.Println("Fetching input")
	url := fmt.Sprintf("%s/input", baseUrl(year, day))
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Failed to create request %s", err)
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	return string(body)
}

// Create the base URL for a problem.
func baseUrl(year Year, day Day) string {
	return fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
}

type AddHeaderTransport struct {
	T http.RoundTripper
}

func (adt *AddHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", "github.com/OliverFlecke/advent-of-code-go by oliverfl@live.dk")

	token, has_token := os.LookupEnv(AOC_TOKEN_NAME)
	if !has_token {
		log.Panicf("Missing `%s` in environment", AOC_TOKEN_NAME)
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: token})
	return adt.T.RoundTrip(req)
}

func newAddHeaderTransport(T http.RoundTripper) *AddHeaderTransport {
	if T == nil {
		T = http.DefaultTransport
	}
	return &AddHeaderTransport{T}
}
