package cmd

import (
	"flag"
	"fmt"
	"os"

	"ibfd.org/gth2cms/paths"
	"ibfd.org/gth2cms/tio"
)

// ParseCommandLine extracts flags and directory names from the command line.
func ParseCommandLine() (bool, *paths.DirectoryNamer) {
	var verbose bool
	flag.Usage = usage
	flag.BoolVar(&verbose, "v", false, "verbose output")
	flag.Parse()
	if flag.NArg() != 2 {
		flag.Usage()
	} else {
		indir, outdir := flag.Arg(0), flag.Arg(1)
		if !tio.IsExistingDirectory(indir) {
			fmt.Printf("%s is not an existing directory\n", indir)
			exit()
		}
		return verbose, paths.NewDirectoryNamer(indir, outdir)
	}
	return false, nil
}

func usage() {
	fmt.Printf("Usage of %s:\n", os.Args[0])
	fmt.Printf("\t%s [-v] <input directory> <output directory>\n", os.Args[0])
	flag.PrintDefaults()
	exit()
}

func exit() {
	os.Exit(2)
}
