package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	models "news/Models"
)

func FetchNews(country, category, apiKey string) (models.NewsResponse, error) {
	url := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=%s&category=%s&apiKey=%s", country, category, apiKey)
	response, err := http.Get(url)
	if err != nil {
		return models.NewsResponse{}, err
	}
	defer response.Body.Close()

	var newsResp models.NewsResponse
	err1 := json.NewDecoder(response.Body).Decode(&newsResp)
	if err1 != nil {
		return models.NewsResponse{}, err
	}

	return newsResp, nil
}
func APIKEY(apiKey string) bool {
	if apiKey == "2750aaecfd0d43f98f9fd9e65d1d830a" {
		return true
	} else {
		return false
	}
}
