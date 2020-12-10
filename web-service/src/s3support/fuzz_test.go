package s3support

import (
	"bytes"
	"os"
	"path"
	"runtime"
	"testing"

	initialize "web-service/src/init/test_init"

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
	initialize.ConfigureTestingEnvironmentForS3Support()
	InitializeS3Support()
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

	StoreFileByUUID(id, buffer, fileName)
}
