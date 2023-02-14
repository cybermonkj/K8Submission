package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
)

type result struct {
	url  string
	size int64
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a list of URLs separated by spaces: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	urls := strings.Split(text, " ")

	var results []result
	for i, url := range urls {
		if i%10 == 0 {
			// Write the results to a file and reset the results slice
			// every 10 URLs to avoid exceeding the file size limit.
			writeResults(results)
			results = nil
		}

		res, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error fetching URL %s: %s\n", url, err)
			continue
		}
		defer res.Body.Close()
		results = append(results, result{
			url:  url,
			size: res.ContentLength,
		})
	}

	writeResults(results)
}

func writeResults(results []result) {
	sort.Slice(results, func(i, j int) bool {
		return results[i].size > results[j].size
	})

	for _, r := range results {
		fmt.Printf("%s: %d bytes\n", r.url, r.size)
	}
}

