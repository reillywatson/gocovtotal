package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: gocovtotal <filename>")
		os.Exit(1)
	}
	f, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		os.Exit(1)
	}
	totalLines := 0.0
	totalActive := 0.0

	for _, l := range strings.Split(string(f), "\n") {
		if strings.Contains(l, "test_services") {
			continue
		}
		words := strings.Split(l, " ")
		if len(words) != 3 {
			continue
		}
		lines, _ := strconv.ParseFloat(words[1], 64)
		active, _ := strconv.ParseFloat(words[2], 64)
		totalLines += lines
		if active > 0 {
			totalActive += lines
		}
	}
	fmt.Printf("%.2f%% covered\n", totalActive/totalLines*100.0)
}
