 This code is a Go program that takes a list of URLs as input from the user, visits each URL, and prints to the output 
 
 
 the list of pairs: URL and response body size. The output is sorted by the size of the response body.

The program begins by prompting the user to enter a list of URLs separated by spaces. It then reads the user's input and splits it into individual URLs using Go's built-in string manipulation functions.

For each URL, the program makes a GET request using the net/http package and records the size of the response body. The program collects the results for every 10 URLs visited and writes them to a file using the writeResults function, which sorts the results by response size and prints them to the console.


Prerequisites

Go installed on your system
A terminal to run the application
 the Application

go run ./Task1/clitool.go and hit enter to launch the tool.

The main function is the entry point of the application. It reads a list of URLs from the user input, visits each URL, and saves the results to a file. The results are sorted by the size of the response body.

The writeResults function writes the results to the output. It first sorts the results by size and then prints out the URL and response body size for each result.

#Task 2

The task 2 file has two different implementations, a Mermaid Diagram which displays the UML diagram of the K8 architecture (Better than Hand drawn). and the second is a .YAML file with the k8 implementation.

The HTML file diaplays the merpaid diagram using mermaid.js but you can also via

https://mermaid-js.github.io/mermaid-live-editor/ Copy the mermaid code and paste it into the editor and the image would be displayed
