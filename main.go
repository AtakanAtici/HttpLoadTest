package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main(){

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <url>")
		return
	}

	url := args[1]
	

	resp, err := http.Get(url)
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