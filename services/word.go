package services

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/whdgus906/wordAdvisor/models"
)

func WordAnalysis(wordList []string, apiKey string) []models.WordResponse {
	// 최대 10개만
	max := 10
	if len(wordList) < max {
		max = len(wordList)
	}
	var wait sync.WaitGroup
	wait.Add(max)
	var channels []models.Channel

	list := wordList[:max]
	for _, v := range list {
		go func(word string) {
			defer wait.Done()
			req, _ := http.NewRequest("GET", fmt.Sprintf("https://stdict.korean.go.kr/api/search.do?key=%s&method=exact&q=%s", apiKey, word), nil)
			client := &http.Client{}
			if res, err := client.Do(req); err == nil {
				defer res.Body.Close()
				if body, err := ioutil.ReadAll(res.Body); err == nil {
					channel := &models.Channel{}
					xml.Unmarshal([]byte(string(body)), channel)
					if channel.Item.Word != "" && channel.Item.Sense.Definition != "" {
						channels = append(channels, *channel)
					}
				}
			}
		}(v)
	}
	wait.Wait()

	var res []models.WordResponse
	for _, channel := range channels {
		res = append(res, *&models.WordResponse{Word: channel.Item.Word, Meaning: channel.Item.Sense.Definition})
	}

	return res
}
