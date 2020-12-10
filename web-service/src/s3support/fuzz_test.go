package s3support_test

import (
	"bytes"
	"os"
	"path"
	"runtime"
	"testing"

	initialize "web-service/src/init"
	s3support "web-service/src/s3support"

	fuzz "github.com/google/gofuzz"
	guuid "github.com/google/uuid"
)

func TestMain(m *testing.M) {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
	initialize.Configure()
	os.Exit(m.Run())
}

func TestFoo(t *testing.T) {
	// Configre the Fuzzer, if needed
	f := fuzz.New()

	var raw_bytes []byte
	f.Fuzz(&raw_bytes)

	buffer := bytes.NewBuffer(raw_bytes)

	var id guuid.UUID
	f.Fuzz(&id)

	var fileName string
	f.Fuzz(&fileName)

	err := s3support.StoreFileByUUID(id, buffer, fileName)
	if err != nil {
		panic(err)
	}
}
