package deployer

import (
	"io"
	"io/ioutil"
	"kafka-connect-deployer/internal/config"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestKafkaConnectDeployer_Deploy(t *testing.T) {
	t.Run("should deploy sink connector", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			assertEqual(t, uri, req.URL.Path)
			assertEqual(t, http.MethodPost, req.Method)
			assertEqual(t, contentType, req.Header.Get("Content-Type"))

			file, err := os.Open("testdata/sink/file-sink.json")
			if err != nil {
				t.Fatal("fail to test file", err)
			}
			expected := getContent(t, "testdata/sink/file-sink.json", file)
			body := getContent(t, "body", req.Body)
			assertEqual(t, expected, body)

			res.WriteHeader(http.StatusAccepted)
		}))
		defer func() { testServer.Close() }()

		c := config.Config{
			ConnectUrl:    testServer.URL,
			ConnectorsDir: "testdata/sink",
		}

		deployer := New(c)
		if err := deployer.Deploy(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("should fail deploy connector when was a invalid directory", func(t *testing.T) {
		c := config.Config{
			ConnectUrl:    "locahost",
			ConnectorsDir: "test",
		}

		deployer := New(c)
		err := deployer.Deploy()
		assertEqual(t, "fail to read dir test. error: open test: no such file or directory", err.Error())
	})

	t.Run("should fail deploy connector when API requested return not ok", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusBadRequest)
			if _, err := res.Write([]byte("invalid connector")); err != nil {
				t.Fatal("fail to write response body")
			}
		}))
		defer func() { testServer.Close() }()

		c := config.Config{
			ConnectUrl:    testServer.URL,
			ConnectorsDir: "testdata/sink",
		}

		deployer := New(c)
		err := deployer.Deploy()
		assertEqual(t, "response not ok: file-sink.json - code: 400 - body: invalid connector", err.Error())
	})
}

func getContent(t *testing.T, name string, c io.ReadCloser) string {
	fileBody, err := ioutil.ReadAll(c)
	if err != nil {
		t.Fatal("fail to get content from "+name, err)
	}

	return string(fileBody)
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected == actual {
		return
	}
	t.Errorf("Expected: %v (type %v) but actual: %v (type %v)",
		expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual))
}
