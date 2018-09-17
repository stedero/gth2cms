// gth2cms adds metadata to Country Chapters for bulkload in the Alfresco CMS.
// Every Country Chapter is stored in a separate XML file.
// Metadata is created by extracting and converting data from a Country Chapter
// and store that in a separate XML file that complies with the JAVA properties
// file DTD (http://java.sun.com/dtd/properties.dtd).
package main

import (
	"ibfd.org/gth2cms/cmd"
	"ibfd.org/gth2cms/exec"
	"ibfd.org/gth2cms/model"
	"ibfd.org/gth2cms/paths"
	"ibfd.org/gth2cms/stats"
	"ibfd.org/gth2cms/gio"
	"ibfd.org/gth2cms/xlate"
)

func main() {
	verbose, mapper, directoryNamer := cmd.ParseCommandLine()
	reporter := stats.NewReporter(verbose)
	visitor := exec.NewVisitor(mapper, directoryNamer, process, reporter)
	visitor.Walk()
	reporter.End()
}

func process(mapper *xlate.Mapper, fileNamer *paths.Filenamer) {
	gthReader := gio.NewGthReader(fileNamer.InputFilename(), fileNamer.OutputFilename())
	defer gthReader.Close()
	gthChapter := model.ReadGthChapter(gthReader)
	metaData := model.NewMetaData(mapper, gthChapter)
	metaFile := gio.CreateFile(fileNamer.MetaFilename())
	defer metaFile.Close()
	metaData.WriteXML(metaFile)
}
