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

type column int

// Excel column references. First letter is the column.
const (
	AcollectionCode column = iota
	BcollectionName
	CcountryName
	Dauthors
	Ette
	FregionOrTeam
)

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
		data[record[AcollectionCode]] = record
	}
	return &Mapper{data}
}

// HumanCollection returns the collection code in a human friendly format.
func (mapper *Mapper) HumanCollection(key string) string {
	return mapper.data[key][BcollectionName]
}

// HumanTitle returns the document title in a human friendly format.
func (mapper *Mapper) HumanTitle(key string) string {
	return mapper.data[key][CcountryName] + " - " + mapper.data[key][AcollectionCode]
}

// RegionOrTeam returns a region or team.
func (mapper *Mapper) RegionOrTeam(key string) string {
	return mapper.data[key][FregionOrTeam]
}

// Authors returns a list of authors separated by semicolons.
func (mapper *Mapper) Authors(key string) string {
	return mapper.data[key][Dauthors]
}

// TTE returns TTE values.
func (mapper *Mapper) TTE(key string) string {
	return mapper.data[key][Ette]
}
