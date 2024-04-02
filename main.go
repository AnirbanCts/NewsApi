package main

import (
	"fmt"
	"net/http"
	controller "news/Controller"
)

func main() {

	http.HandleFunc("/news/headlines", controller.NewsController)
	fmt.Println("Starting web server @ http://localhost:9999")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println(err)
	}
}
