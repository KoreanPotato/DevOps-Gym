package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main () {

	if len(os.Args) < 3 {
		fmt.Println("you need three arguments!")
		os.Exit(1)
	}

	host := os.Args[1]
	name := os.Args[2]

	url := fmt.Sprintf("http://%s?name=%s", host,  name)

	res, err := http.Get(url)

	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	fmt.Println(string(body))
}
