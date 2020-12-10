package s3support

import (
	"os"
	"path"
	"runtime"
	"testing"
	initialize "web-service/src/init/test_init"
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

}
