package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/valyala/fasthttp"
)

// standardHTTPDemo demonstrates a simple HTTP POST request using the standard http package
func standardHTTPDemo() {
	fmt.Println("Standard HTTP Demo")

	url := "https://httpbin.org/post"
	method := "POST"

	payload := strings.NewReader("This is sample POST data")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("CUSTOM-Header", "MyCustomValue")

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

	var buf bytes.Buffer
	res.Header.WriteSubset(&buf, nil)
	headers := buf.String()

	fmt.Println("Response Headers:")
	fmt.Print(headers)

	fmt.Println("\nResponse Body:")
	fmt.Println(string(body))
}

// fastHttpDemo demonstrates a simple HTTP POST request using the fasthttp package
func fastHttpDemo() {
	fmt.Println("Fast HTTP Demo")

	url := "https://httpbin.org/post"

	payload := []byte("This is sample POST data")

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("FAST-HTTP-Header", "MyCustomValue")
	req.SetBody(payload)

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	client := &fasthttp.Client{}
	err := client.Do(req, res)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Response Headers:")
	res.Header.VisitAll(func(key, value []byte) {
		fmt.Printf("%s: %s\n", string(key), string(value))
	})

	fmt.Println("\nResponse Body:")
	fmt.Println(string(res.Body()))
}

// responseHeadersExample demonstrates retrieving case-sensitive response headers using fasthttp
func responseHeadersExample() {
	fmt.Println("Case-sensitive Response Headers Example")

	url := "https://httpbin.org/response-headers?Custom-Header=MyCustomValue&Another-Header=AnotherValue"

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("GET")

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	client := &fasthttp.Client{}
	err := client.Do(req, res)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Response Headers:")
	res.Header.VisitAll(func(key, value []byte) {
		fmt.Printf("%s: %s\n", string(key), string(value))
	})
}

func main() {
	// standardHTTPDemo()

	// fmt.Println("\n**********")
	// fastHttpDemo()

	// fmt.Println("\n**********")
	// responseHeadersExample()
	goClientDemo()
}
