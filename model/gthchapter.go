package model

import (
	"encoding/xml"
	"io"
	"log"
)

// GthChapter defines (partially) the XML structure of a
// Country Chapter.
type GthChapter struct {
	Collection string `xml:"collection,attr"`
	Anchor     string `xml:"anchor,attr"`
	ReviewDate struct {
		Date string `xml:"reviewdate,attr"`
	} `xml:"chaphead>latestinfo>reviewdate"`
	Country struct {
		CC string `xml:",attr"`
	} `xml:"chaphead>country"`
}

// ReadGthChapter transforms a Country Chapter in XML into an internal structure.
func ReadGthChapter(r io.Reader) *GthChapter {
	var gthChapter GthChapter
	decoder := xml.NewDecoder(r)
	err := decoder.Decode(&gthChapter)
	if err != nil {
		log.Fatalf("error unmarshaling Country Chapter: %v", err)
	}
	return &gthChapter
}
