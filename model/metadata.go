package model

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
	"ibfd.org/gth2cms/xlate"
)

const metaDataDoctype = "<!DOCTYPE properties SYSTEM \"http://java.sun.com/dtd/properties.dtd\">\n"

// MetaData element for marshaling to XML
type MetaData struct {
	XMLName xml.Name `xml:"properties"`
	Entries []Entry  `xml:"entry"`
}

// Entry element for marshaling to XML
type Entry struct {
	Key   string `xml:"key,attr"`
	Value string `xml:",innerxml"`
}

const multiValueJoiner = ","

// NewMetaData transforms a Country Chapter into a meta data structure.
func NewMetaData(mapper *xlate.Mapper, gthChapter *GthChapter) *MetaData {
	col := gthChapter.Collection
	metaData := &MetaData{}
	metaData.add("type", "ibfd:onlinecontent")
	metaData.add("aspects", "ibfd:countryChapter,ibfd:onlineContentProperties,cm:titled")
	metaData.add("cm:title", mapper.HumanTitle(col))
	metaData.add("ibfd:collectionCode", col)
	metaData.add("ibfd:collectionCodeHumanReadable", mapper.HumanCollection(col))
	metaData.add("ibfd:regionOrTeam", mapper.RegionOrTeam(col))
	metaData.add("ibfd:lastReviewDate", gthChapter.ReviewDate + "T00:00:00.000+02:00")
	return metaData
}

// WriteXML writes the metadata as XML.
func (m *MetaData) WriteXML(w io.Writer) {
	writeString(w, xml.Header)
	writeString(w, nowAsComment())
	writeString(w, metaDataDoctype)
	encoder := xml.NewEncoder(w)
	encoder.Indent("", "   ")
	err := encoder.Encode(m)
	if err != nil {
		log.Fatalf("error encoding XML: %v", err)
	}
}

func (m *MetaData) add(key string, value string) {
	m.Entries = append(m.Entries, Entry{key, value})
}

func nowAsComment() string {
	return fmt.Sprintf("<!-- Generated by gth2cms on %s -->\n", time.Now().Format(time.RFC3339))
}

func mapJoin(len int, get func(int) string) string {
	result := make([]string, len)
	for i := 0; i < len; i++ {
		result[i] = get(i)
	}
	return strings.Join(result, multiValueJoiner)
}

func writeString(w io.Writer, s string) {
	_, err := io.WriteString(w, s)
	if err != nil {
		log.Fatalf("error writing %s: %v", s, err)
	}
}
