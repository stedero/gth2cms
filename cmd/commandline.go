package cmd

import (
	"flag"
	"fmt"
	"os"

	"ibfd.org/gth2cms/gio"
	"ibfd.org/gth2cms/paths"
	"ibfd.org/gth2cms/xlate"
)

// ParseCommandLine extracts flags, Excel mapping filename and
// directory names from the command line.
func ParseCommandLine() (bool, *xlate.Mapper, *paths.DirectoryNamer) {
	var verbose bool
	flag.Usage = usage
	flag.BoolVar(&verbose, "v", false, "verbose output")
	flag.Parse()
	if flag.NArg() != 3 {
		flag.Usage()
	} else {
		mappingFile, indir, outdir := flag.Arg(0), flag.Arg(1), flag.Arg(2)
		if !gio.IsExistingFile(mappingFile) {
			fmt.Printf("%s not found\n", mappingFile)
			exit()
		}
		if !gio.IsExistingDirectory(indir) {
			fmt.Printf("%s is not an existing directory\n", indir)
			exit()
		}
		return verbose, xlate.NewMapper(mappingFile), paths.NewDirectoryNamer(indir, outdir)
	}
	return false, nil, nil
}

func usage() {
	fmt.Printf("Usage of %s:\n", os.Args[0])
	fmt.Printf("\t%s [-v] <Excel mappings file> <input directory> <output directory>\n", os.Args[0])
	flag.PrintDefaults()
	exit()
}

func exit() {
	os.Exit(2)
}
