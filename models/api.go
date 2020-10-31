package models

import "encoding/xml"

// 오픈코리아텍스트
// https://github.com/open-korean-text/open-korean-text
type Token struct {
	Tokens        []string
	Token_strings []string
}
type Phrases struct {
	Phrases []string
}

// 국립국어원 표준국어 대사전
// https://stdict.korean.go.kr/openapi/openApiInfo.do
type Channel struct {
	XMLName       xml.Name `xml:"channel"`
	Text          string   `xml:",chardata"`
	Title         string   `xml:"title"`
	Link          string   `xml:"link"`
	Description   string   `xml:"description"`
	LastBuildDate string   `xml:"lastBuildDate"`
	Total         string   `xml:"total"`
	Start         string   `xml:"start"`
	Num           string   `xml:"num"`
	Item          struct {
		Text       string `xml:",chardata"`
		TargetCode string `xml:"target_code"`
		Word       string `xml:"word"`
		SupNo      string `xml:"sup_no"`
		Pos        string `xml:"pos"`
		Sense      struct {
			Text       string `xml:",chardata"`
			Definition string `xml:"definition"`
			Link       string `xml:"link"`
			Type       string `xml:"type"`
		} `xml:"sense"`
	} `xml:"item"`
}
