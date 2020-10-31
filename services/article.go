package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/whdgus906/wordAdvisor/models"
)

func ArticleAnalysis(article string) []string {
	wordList := []string{}
	q := article

	req, _ := http.NewRequest("GET", fmt.Sprintf("https://open-korean-text-api.herokuapp.com/extractPhrases?text=%s", q), nil)
	client := &http.Client{}
	if res, err := client.Do(req); err == nil {
		defer res.Body.Close()
		if body, err := ioutil.ReadAll(res.Body); err == nil {
			token := &models.Phrases{}
			json.Unmarshal([]byte(string(body)), token)
			wordList = token.Phrases
		}
	}
	return wordList
}
