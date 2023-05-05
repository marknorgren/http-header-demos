package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func goClientDemo() {
	url := "http://localhost:8000"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Go-Client")
	req.Header.Add("Custom-Header", "MyCustomValue")
	req.Header.Add("Another-Header", "AnotherValue")
	req.Header.Add("User-Agent", "Python-Client")
	req.Header.Add("MARK-norgren", "MakdaLLAsl")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Response Headers:")
	for key, value := range res.Header {
		fmt.Printf("%s: %s\n", key, value[0])
	}

	fmt.Println("\nResponse Body:")
	fmt.Println(string(body))
}
