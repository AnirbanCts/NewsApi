package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	service "news/Services"
)

func NewsController(write http.ResponseWriter, readvalue *http.Request) {
	country := readvalue.FormValue("Country")
	category := readvalue.FormValue("Category")
	apiKey := readvalue.Header.Get("x-api-key")
	isValid := service.APIKEY(apiKey)
	if !isValid {
		http.Error(write, "Unauthorized", http.StatusUnauthorized)
		return
	}
	news, err := service.FetchNews(country, category, apiKey)
	if err != nil {
		fmt.Println("Error fetching news:", err)
		return
	}
	//json.NewEncoder(write).Encode(news.TotalResults)
	//json.NewEncoder(write).Encode(news.Status)
	for _, article := range news.Articles {
		json.NewEncoder(write).Encode("Title:    " + article.Title)
		json.NewEncoder(write).Encode("Description:    " + article.Description)
	}
}
