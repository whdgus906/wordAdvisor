package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/whdgus906/wordAdvisor/config"
	"github.com/whdgus906/wordAdvisor/models"
	"github.com/whdgus906/wordAdvisor/services"
)

type ArticleController struct {
	Cfg *config.Config
}

func (ac *ArticleController) Analysis(ctx *gin.Context) {
	var req models.ArticleRequest
	ctx.ShouldBindJSON(&req)

	wordList := services.ArticleAnalysis(req.Article)
	wordList = nounFilter(wordList)
	wordList = duplicateFilter(wordList)

	dictList := services.WordAnalysis(wordList, ac.Cfg.DictAPIKey)

	if len(dictList) == 0 {
		ctx.Status(http.StatusNoContent)
		return
	}

	ctx.JSON(http.StatusOK, dictList)
	return
}

func nounFilter(list []string) []string {
	filteredList := []string{}
	for _, v := range list {
		filteredList = append(filteredList, v[:strings.Index(v, "(Noun")])
	}
	return filteredList
}

func duplicateFilter(list []string) []string {
	filteredList := []string{}
	for _, word := range list {
		isPossible := true
		for _, nonDuplicateWord := range filteredList {
			if strings.ContainsAny(nonDuplicateWord, word) {
				isPossible = false
				break
			}
		}
		if isPossible {
			filteredList = append(filteredList, word)
		}
	}
	return filteredList
}
