package api_test

import (
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"testing"
	"net/http/httptest"
	"encoding/json"
	"mime/multipart"
	"bytes"

	api "web-service/src/api/controllers"
	initialize "web-service/src/init"
)

func UploadFilesActivityWorksInNormalUseCase(t *testing.T) {
	w := httptest.NewRecorder()
	content := []byte("This is a file with some content")
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	part, _ := writer.CreateFormFile("file_uploads", "sample_file.txt")
	part.Write(content)
	writer.Close()

	req := httptest.NewRequest("POST", "/", &buffer)
	req.Header.Set("Content-Type", "multipart/form-data; " + "boundary=" + "\"" + writer.Boundary() + "\"")
	api.UploadFilesHandler(w, req)
 	result := w.Result()

	var unmarshaled map[string]interface{}
	body_bytes, _ := ioutil.ReadAll(result.Body)
	json.Unmarshal(body_bytes, &unmarshaled)
	t.Log(unmarshaled)

	if result.StatusCode != 200 {
		t.Error("Response status code should be 200")
	}
}

func UploadFilesActivityHandlesErrors(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	api.UploadFilesHandler(w, req)
 	result := w.Result()
	if result.StatusCode == 200 {
		t.Error("Status code should not be 200")
	}
	if result.Body == nil {
		t.Error("Body should not be nil")
	}
	if value, ok := result.Header["Content-Type"]; ok {
		if len(value) != 1 {
			t.Error("Header slice should have length 1")
		}
		expected := "application/json; charset=utf-8"
		if value[0] != expected {
			t.Errorf("Content-Type should be %s", expected)
		}
	} else {
		t.Error("Header should contain Content-Type")
	}

	var body map[string]string
	json.NewDecoder(result.Body).Decode(&body)
	if _, ok := body["Error"]; !ok {
		t.Error("Header should contain Content-Type")
	}
}

func UploadFilesActivity(t *testing.T) {
	t.Run("UploadFiles", UploadFilesActivityHandlesErrors)
	t.Run("UploadFiles", UploadFilesActivityWorksInNormalUseCase)
}

func TestUploadFilesHandler(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
	initialize.Configure()

	t.Run("UploadFiles", UploadFilesActivity)
}
