package paths

import (
	"path/filepath"
	"strings"
)

// DirectoryNamer holds the names of the input- and
// output directories.
type DirectoryNamer struct {
	indir  string
	outdir string
}

// NewDirectoryNamer creates a new instance.
func NewDirectoryNamer(indir string, outdir string) *DirectoryNamer {
	return &DirectoryNamer{indir, outdir}
}

// InDir returns the input directory name.
func (directoryNamer *DirectoryNamer) InDir() string {
	return directoryNamer.indir
}

// OutDir returns the output directory name.
func (directoryNamer *DirectoryNamer) OutDir() string {
	return directoryNamer.outdir
}

// NewOutdirName creates a new output directory name from the specified input path.
func (directoryNamer *DirectoryNamer) NewOutdirName(inputPath string) string {
	part := strings.TrimPrefix(inputPath, directoryNamer.indir)
	return filepath.Join(directoryNamer.outdir, part)
}
