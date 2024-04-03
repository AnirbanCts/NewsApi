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
	response := []models.ArticleResponse{}
	for _, article := range news.Articles {
		response = append(response, models.ArticleResponse{Title: article.Title, Description: article.Description})

	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(write, "Could not convert to JSON", 500)
		return
	} else {

		write.Write(jsonResponse)
	}
}
