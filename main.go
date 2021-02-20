package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, _ := http.Get("https://example.com")

	fmt.Printf("[\x1b[32m%s\x1b[0m]\n", "OK   ")
	fmt.Printf("[\x1b[31m%s\x1b[0m]\n", "   NG")

	//fmt.Println(res.StatusCode)
	fmt.Println(res.Proto)
	fmt.Println(res.Header["Date"])
	fmt.Println(res.Header["Content-Type"])

	fmt.Println(res.Request.Method)
	fmt.Println(res.Request.URL)

	//defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	//fmt.Print(string(body))

}
