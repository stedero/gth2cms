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

// Fetcher a data fetcher.
type Fetcher struct {
	key    string
	mapper *Mapper
}

type column int

// Excel column references. First letter is the column.
const (
	Auid column = iota
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
		data[record[Auid]] = record
	}
	return &Mapper{data}
}

// Fetcher creates a data fetcher for the specified key.
func (mapper *Mapper) Fetcher(key string) *Fetcher {
	return &Fetcher{key, mapper}
}

// HumanCollection returns the collection code in a human friendly format.
func (fetcher *Fetcher) HumanCollection() string {
	return fetcher.mapper.HumanCollection(fetcher.key)
}

// HumanCollection returns the collection code in a human friendly format.
func (mapper *Mapper) HumanCollection(key string) string {
	return mapper.data[key][BcollectionName]
}

// HumanTitle returns the document title in a human friendly format.
func (fetcher *Fetcher) HumanTitle() string {
	return fetcher.mapper.HumanTitle(fetcher.key)
}

// HumanTitle returns the document title in a human friendly format.
func (mapper *Mapper) HumanTitle(key string) string {
	return mapper.data[key][CcountryName] + " - " + mapper.data[key][BcollectionName]
}

// RegionOrTeam returns a region or team.
func (fetcher *Fetcher) RegionOrTeam() string {
	return fetcher.mapper.RegionOrTeam(fetcher.key)
}

// RegionOrTeam returns a region or team.
func (mapper *Mapper) RegionOrTeam(key string) string {
	return mapper.data[key][FregionOrTeam]
}

// Authors returns a list of authors separated by semicolons.
func (fetcher *Fetcher) Authors() string {
	return fetcher.mapper.Authors(fetcher.key)
}

// Authors returns a list of authors separated by semicolons.
func (mapper *Mapper) Authors(key string) string {
	return mapper.data[key][Dauthors]
}

// TTE returns TTE values.
func (fetcher *Fetcher) TTE() string {
	return fetcher.mapper.TTE(fetcher.key)
}

// TTE returns TTE values.
func (mapper *Mapper) TTE(key string) string {
	return mapper.data[key][Ette]
}
