package model

import (
	"encoding/xml"
	"io"
	"log"
)

// GthChapter defines the XML structure of a
// Country Chapter.
type GthChapter struct {
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

// ReadGthChapter transforms a TNS article in XML into an internal structure.
func ReadGthChapter(r io.Reader) *GthChapter {
	var gthChapter GthChapter
	decoder := xml.NewDecoder(r)
	err := decoder.Decode(&gthChapter)
	if err != nil {
		log.Fatalf("error unmarshaling TNS article: %v", err)
	}
	gthChapter.addDTDDefaults()
	return &gthChapter
}

func (gthChapter *GthChapter) addDTDDefaults() {
	if gthChapter.Collection == "" {
		gthChapter.Collection = "tns"
	}
	if gthChapter.ReportType == "" {
		gthChapter.ReportType = "standard"
	}
}
