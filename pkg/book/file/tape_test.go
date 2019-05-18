package file

import (
	"gitlab.com/patchwell/ledger/pkg/test"
	"io/ioutil"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := test.CreateTempFile(t, "123", "test")
	defer clean()

	tape := &tape{file}

	tape.Write([]byte("456"))

	file.Seek(0, 0)
	contents, _ := ioutil.ReadAll(file)

	got := string(contents)
	want := "456"

	if got != want {
		t.Errorf("got incorrect data from reading file, got '%s', wanted '%s'", got, want)
	}
}