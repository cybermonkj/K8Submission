
reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter a list of URLs separated by spaces: ")
text, _ := reader.ReadString('\n')
text = strings.TrimSpace(text)
urls := strings.Split(text, " ")

type result struct {
	url  string
	size int64
}

var results []result
for _, url := range urls {
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

sort.Slice(results, func(i, j int) bool {
	return results[i].size > results[j].size
})

for _, r := range results {
	fmt.Printf("%s: %d bytes\n", r.url, r.size)
}
}
