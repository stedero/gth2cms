package tio

import (
	"io"
	"log"
	"os"
)

// GthXML info struct.
type GthXML struct {
	FileName string
	Data     []byte
}

// GthReader for copying a file while reading
type GthReader struct {
	reader    io.ReadCloser
	writer    io.WriteCloser
	teeReader io.Reader
}

// IsExistingDirectory determines whether a given directory
// is an existing directory.
func IsExistingDirectory(dir string) bool {
	file, err := os.Stat(dir)
	return err == nil && file.IsDir()
}

// CreateDirIfNotExist creates a directory if it does not
// exist yet. If an error occurs then the program terminates
// with a panic message.
func CreateDirIfNotExist(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatalf("fail to create directory %s: %v", dir, err)
		}
		return true
	}
	return false
}

// NewGthReader creates a reader that copies the output when reading.
func NewGthReader(inputFilename, outputFilename string) *GthReader {
	reader := OpenFile(inputFilename)
	writer := CreateFile(outputFilename)
	teeReader := io.TeeReader(reader, writer)
	return &GthReader{reader, writer, teeReader}
}

func (tr *GthReader) Read(p []byte) (n int, err error) {
	return tr.teeReader.Read(p)
}

// OpenFile opens a file for reading.
// If an error occurs then the program terminates with a panic message.
func OpenFile(filename string) io.ReadCloser {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file %s: %v", filename, err)
	}
	return file
}

// CreateFile creates a file for writing.
// If an error occurs then the program terminates with a panic message.
func CreateFile(filename string) io.WriteCloser {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("failed to create file %s: %v", filename, err)
	}
	return file
}

// Close closes the TNS reader.
func (tr *GthReader) Close() {
	tr.reader.Close()
	tr.writer.Close()
}
