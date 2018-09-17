package exec

import (
	"os"
	"path/filepath"

	"ibfd.org/gth2cms/gio"
	"ibfd.org/gth2cms/paths"
	"ibfd.org/gth2cms/stats"
	"ibfd.org/gth2cms/xlate"
)

// Processor is the function to process every file
type Processor func(mapper *xlate.Mapper, Filenamer *paths.Filenamer)

// Visitor defines a visitor
type Visitor struct {
	mapper		 *xlate.Mapper
	rootDirNamer *paths.DirectoryNamer
	process      Processor
	reporter     *stats.Reporter
}

// NewVisitor creates a directory visitor which scans a directory tree
// recursively and processes each valid file using the supplied processor.
func NewVisitor(mapper *xlate.Mapper, rootDirNamer *paths.DirectoryNamer, processor Processor, reporter *stats.Reporter) *Visitor {
	return &Visitor{mapper, rootDirNamer, processor, reporter}
}

// Walk the directory tree and process each file.
func (visitor *Visitor) Walk() error {
	return filepath.Walk(visitor.rootDirNamer.InDir(), visitor.walker())
}

// Walker returns a directory walker function.
func (visitor *Visitor) walker() func(string, os.FileInfo, error) error {
	rootDirNamer := visitor.rootDirNamer
	reporter := visitor.reporter
	return func(path string, fileInfo os.FileInfo, err error) error {
		action := paths.Validate(fileInfo)
		defer reporter.Register(action, path)
		switch action {
		case paths.AcceptFile:
			reporter.CreatedFiles(2) // We create 2 new files for every inputfile.
			fileNamer := paths.NewFilenamer(rootDirNamer, path, fileInfo)
			visitor.process(visitor.mapper, fileNamer)
		case paths.AcceptDir:
			dest := rootDirNamer.NewOutdirName(path)
			reporter.CreatedDir(gio.CreateDirIfNotExist(dest))
		case paths.RejectDir:
			return filepath.SkipDir
		case paths.RejectFile:
		default:
		}
		return err
	}
}
