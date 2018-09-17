package stats

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/inflection"
	"ibfd.org/gth2cms/paths"
)

// Reporter definition
type Reporter struct {
	start              time.Time
	verbose            bool
	actionCounts       []int
	createdDirectories int
	createdFiles       int
}

// NewReporter creates a reporter.
func NewReporter(verbose bool) *Reporter {
	return &Reporter{time.Now(), verbose, make([]int, paths.NumActions), 0, 0}
}

// Register an action.
func (reporter *Reporter) Register(action paths.Action, path string) {
	reporter.actionCounts[action]++
	if reporter.verbose {
		switch action {
		case paths.RejectDir:
			log.Printf("skipped directory: %s\n", path)
		case paths.RejectFile:
			log.Printf("skipped file: %s\n", path)
		default:
		}
	}
}

// CreatedFiles increments the count for the number of files created.
func (reporter *Reporter) CreatedFiles(count int) {
	reporter.createdFiles += count
}

// CreatedDir counts the number of directories created
func (reporter *Reporter) CreatedDir(didCreate bool) {
	if didCreate {
		reporter.createdDirectories++
	}
}

// End logs the termination of the process with counts.
func (reporter *Reporter) End() {
	counts := reporter.actionCounts
	elapsed := time.Since(reporter.start)
	dirFmt := countFmt("directory")
	fileFmt := countFmt("file")
	log.Printf("skipped %s", dirFmt(counts[paths.RejectDir]))
	log.Printf("skipped %s", fileFmt(counts[paths.RejectFile]))
	log.Printf("created %s", dirFmt(reporter.createdDirectories))
	log.Printf("created %s", fileFmt(reporter.createdFiles))
	log.Printf("processing %s in %s took %s", fileFmt(counts[paths.AcceptFile]), dirFmt(counts[paths.AcceptDir]), elapsed)
}

func countFmt(singular string) func(int) string {
	return func(count int) string {
		return fmtPluralize(singular, count)
	}
}

func fmtPluralize(singular string, count int) string {
	return fmt.Sprintf("%d %s", count, pluralize(singular, count))
}

func pluralize(singular string, count int) string {
	if count == 1 {
		return singular
	}
	return inflection.Plural(singular)
}
