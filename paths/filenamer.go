package paths

import (
	"os"
	"path/filepath"
	"strings"
)

const extension = ".xml"
const metaFileSuffix = ".metadata.properties" + extension

// Action defines the possible actions to perform on a file or directory
type Action int

// The Action response values
const (
	AcceptDir Action = iota
	AcceptFile
	RejectDir
	RejectFile
	NumActions
)

// Filenamer holds all data needed to create input- and outputfilenames.
type Filenamer struct {
	DirectoryNamer
	fileName string
}

// NewFilenamer creates a filenamer from root directory info.
func NewFilenamer(rootDirNamer *DirectoryNamer, currentPath string, fileInfo os.FileInfo) *Filenamer {
	source := strings.TrimSuffix(currentPath, fileInfo.Name())
	part := strings.TrimPrefix(source, rootDirNamer.InDir())
	dest := filepath.Join(rootDirNamer.OutDir(), part)
	return &Filenamer{*NewDirectoryNamer(source, dest), fileInfo.Name()}
}

// Validate determines what to do with a file or directory.
func Validate(fileInfo os.FileInfo) Action {
	if fileInfo.IsDir() {
		if hasValidDirectoryName(fileInfo) {
			return AcceptDir
		}
		return RejectDir
	} else if hasValidExtension(fileInfo) {
		return AcceptFile
	} else {
		return RejectFile
	}
}

// InputFilename returns the path of an input file.
func (fileNamer *Filenamer) InputFilename() string {
	return filepath.Join(fileNamer.indir, fileNamer.fileName)
}

// OutputFilename returns the path of an output file.
func (fileNamer *Filenamer) OutputFilename() string {
	return filepath.Join(fileNamer.outdir, fileNamer.fileName)
}

// MetaFilename returns the path of a meta data file.
func (fileNamer *Filenamer) MetaFilename() string {
	return fileNamer.OutputFilename() + metaFileSuffix
}

func hasValidDirectoryName(fileInfo os.FileInfo) bool {
	return !strings.HasPrefix(fileInfo.Name(), ".")
}

func hasValidExtension(fileInfo os.FileInfo) bool {
	return strings.EqualFold(filepath.Ext(fileInfo.Name()), extension)
}
