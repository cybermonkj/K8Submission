package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"sync"
)

func main() {

	// Get list of URLs from file or standard input
	var urls []string
	if flag.NArg() > 0 {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal("Error opening file:", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			urls = append(urls, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal("Error reading file:", err)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			urls = append(urls, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal("Error reading standard input:", err)
		}
	}

	// Add protocol scheme to URLs if not already present
	for i, u := range urls {
		parsedURL, err := url.Parse(u)
		if err != nil {
			log.Printf("Error parsing URL %s: %s\n", u, err)
			continue
		}
		if parsedURL.Scheme == "" {
			urls[i] = "http://" + u
		}
	}

	// Create a slice to hold the results
	results := make([][2]int, len(urls))

	// Visit each URL and record response body size concurrently
	var errors []error
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for i, url := range urls {
		go func(i int, url string) {
			defer wg.Done()

			resp, err := http.Head(url)
			if err != nil {
				errors = append(errors, fmt.Errorf("Error visiting %s: %s", url, err))
				return
			}

			size := -1
			if contentLength := resp.Header.Get("Content-Length"); contentLength != "" {
				fmt.Sscanf(contentLength, "%d", &size)
			}

			if size < 0 {
				// Fallback to inefficient algorithm
				resp, err = http.Get(url)
				if err != nil {
					errors = append(errors, fmt.Errorf("Error visiting %s: %s", url, err))
					return
				}
				defer resp.Body.Close()

				size = 0
				buf := make([]byte, 1024)
				for {
					n, err := resp.Body.Read(buf)
					if n > 0 {
						size += n
					}
					if err != nil {
						break
					}
				}
			}

			results[i] = [2]int{size, i}
		}(i, url)
	}

	wg.Wait()

	// Sort results by response body size
	sort.Slice(results, func(i, j int) bool {
		return results[i][0] < results[j][0]
	})

	// Print results to stdout and log file
	outFileName := "output.txt"
	outFile, err := os.Create(outFileName)
	if err != nil {
		log.Fatal("Error creating output file:", err)
	}
	defer outFile.Close()

	for _, result := range results {
		size := result[0]
		url := urls[result[1]]
		fmt.Printf("%s %d\n", url, size)
		fmt.Fprintf(outFile, "%s %d\n", url, size)
	}

	if len(errors) > 0 {
		fmt.Fprintln(os.Stderr, "Errors:")
		for _, err := range errors {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

