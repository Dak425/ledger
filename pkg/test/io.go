package test

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func CreateTempFile(t *testing.T, data string, name string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	file, err := ioutil.TempFile("", name)
	if err != nil {
		t.Fatalf("unable to create temp file %v", err)
	}

	file.Write([]byte(data))

	cleanUp := func() {
		os.Remove(file.Name())
	}

	return file, cleanUp
}
