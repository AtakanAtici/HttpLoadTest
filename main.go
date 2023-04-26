package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	resp, err := http.Get("https://example.com")
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err :=ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error readin response:", err)
		return
	}
	fmt.Println(string(body))
}