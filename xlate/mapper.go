package xlate

import (
	"log"
	"strings"
	"github.com/tealeg/xlsx"
)

// Mapper translates country chapter content to CMS equivalents.
type Mapper struct {
	data map[string][]string
}

// NewMapper creates a mapper.
func NewMapper(filename string) *Mapper {
	data := make(map[string][]string)
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		log.Fatalf("failed to open %s: %v", filename, err)
	}
	sheet := xlFile.Sheets[0]
	for _, row := range sheet.Rows {
		record := make([]string, 0, len(row.Cells))
		for _, cell := range row.Cells {
			record = append(record, strings.TrimSpace(cell.String()))
		}
		data[record[0]] = record
	}
	return &Mapper{data}
}

// HumanCollection returns the collection code in a human friendly format.
func (mapper *Mapper) HumanCollection(key string) string {
	data := mapper.get(key)
	if data != nil {
		return data[1]
	} else {
		return ""
	}
}

// HumanTitle returns the document title in a human friendly format.
func (mapper *Mapper) HumanTitle(key string) string {
	data := mapper.get(key)
	if data != nil {
		return data[2] + " - " + data[1] 
	} else {
		return ""
	}
}

// RegionOrTeam returns a region or team.
func (mapper *Mapper) RegionOrTeam(key string) string {
	data := mapper.get(key)
	if data != nil {
		return data[5]
	} else {
		return ""
	}
}

// Authors returns a lits of authors.
func (mapper *Mapper) Authors(key string) string {
	data := mapper.get(key)
	if data != nil {
		return data[4] 
	} else {
		return ""
	}
}

func (mapper *Mapper) get(key string) []string {
	if data, ok := mapper.data[key]; ok {
		return data
	} else {
		log.Printf("No entry for: %s\n", key)
		return data
	}
}
