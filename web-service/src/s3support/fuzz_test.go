package s3support

import (
	"fmt"
	"testing"
	initialize "web-service/src/init/test_init"
)

func TestFoo(t *testing.T) {
	initialize.ConfigureTestingEnvironment()
	fmt.Println("foo")
}
