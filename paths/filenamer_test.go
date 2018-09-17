package paths

import (
	"os"
	"path/filepath"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestInputFilename(t *testing.T) {
	t.Log("Given the need to test filenamer functionality")
	file, _ := os.Stat("filenamer_test.go")
	dirNamer := NewDirectoryNamer("/test/indir", "/test/outdir")
	fileNamer := &Filenamer{*dirNamer, file.Name()}

	expected := "/test/indir/filenamer_test.go"
	actual := filepath.ToSlash(fileNamer.InputFilename())
	assert(t, expected, actual, "InputFilename for \"filenamer_test.go\"")

	expected = "/test/outdir/filenamer_test.go"
	actual = filepath.ToSlash(fileNamer.OutputFilename())
	assert(t, expected, actual, "OutputFilename for \"filenamer_test.go\"")

	expected = "/test/outdir/filenamer_test.go.metadata.properties.xml"
	actual = filepath.ToSlash(fileNamer.MetaFilename())
	assert(t, expected, actual, "MetaFilename for \"filenamer_test.go\"")
}

func assert(t *testing.T, expected string, actual string, what string) {
	t.Logf("\tWhen checking %s", what)
	if actual == expected {
		t.Logf("\t\tIt should be \"%s\" %s", actual, checkMark)
	} else {
		t.Fatalf("\t\tIt should be \"%s\" but it was \"%s\" %s", expected, actual, ballotX)
	}
}
