package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		baseUrl = flag.String("baseUrl", "http://localhost:8080", "server base url.")
	)
	flag.Parse()

	fmt.Println(*baseUrl)

}
