package model

import (
	"encoding/xml"
	"io"
	"log"
)

// TnsArticle defines the XML structure of a
// Tax News Service article.
type TnsArticle struct {
	GUID           string `xml:"guid,attr"`
	Collection     string `xml:"collection,attr"`
	ReportType     string `xml:"reporttype,attr"`
	TnsArticleInfo struct {
		CountryList struct {
			Main      string `xml:"main,attr"`
			Countries []struct {
				CC   string `xml:"cc,attr"`
				Name string `xml:"countryname"`
			} `xml:"country"`
		} `xml:"countrylist"`
		Topics []struct {
			TC          string `xml:"tc,attr"`
			Score       string `xml:"score,attr"`
			Description string `xml:",innerxml"`
		} `xml:"topics>topic"`
		OnlinetTitle string `xml:"onlinetitle"`
		ArticleDate  struct {
			IsoDate   string `xml:"isodate,attr"`
			HumanDate string `xml:",innerxml"`
		} `xml:"articledate"`
		Author struct {
			Initials string `xml:"initials,attr"`
			Name     string `xml:",innerxml"`
		} `xml:"author"`
		Correspondent string `xml:"correspondent"`
		Reference     []struct {
			Target  string `xml:"target,attr"`
			AltText string `xml:"alttext,attr"`
			Xref    string `xml:",innerxml"`
		} `xml:"reference>extxref"`
		Source string `xml:"source"`
	} `xml:"tnsarticleinfo"`
}

// ReadTnsArticle transforms a TNS article in XML into an internal structure.
func ReadTnsArticle(r io.Reader) *TnsArticle {
	var tnsArticle TnsArticle
	decoder := xml.NewDecoder(r)
	err := decoder.Decode(&tnsArticle)
	if err != nil {
		log.Fatalf("error unmarshaling TNS article: %v", err)
	}
	tnsArticle.addDTDDefaults()
	return &tnsArticle
}

func (tnsArticle *TnsArticle) addDTDDefaults() {
	if tnsArticle.Collection == "" {
		tnsArticle.Collection = "tns"
	}
	if tnsArticle.ReportType == "" {
		tnsArticle.ReportType = "standard"
	}
}
