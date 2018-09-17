package xlate

// Mapper translates country chapter content to CMS equivalents.
type Mapper struct {
}

// NewMapper creates a mapper.
func NewMapper(filename string) *Mapper {
	return &Mapper{}
}

// HumanCollection returns the collection code in a human friendly format.
func (mapper *Mapper)HumanCollection(col string) string {
	return ""
}

// HumanTitle returns the document title in a human friendly format.
func (mapper *Mapper)HumanTitle(col string) string {
	return ""
}

// RegionOrTeam returns a region or team.
func (mapper *Mapper)RegionOrTeam(col string) string {
	return ""
}

// Authors returns a lits of authors.
func (mapper *Mapper)Authors(col string) []string {
	authors := make([]string, 3)
	return authors
}
