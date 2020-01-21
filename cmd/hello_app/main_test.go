// 10 OMIT
package main

// 20 OMIT
import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 30 OMIT
func TestRootHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/", nil) // body is nil
	w := httptest.NewRecorder()
	rootHandler(w, req)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("could not read response body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: %v", resp.StatusCode)
	}
	if !bytes.Contains(body, []byte("Hello, the time is")) {
		t.Errorf("unexpected body %s\n", body)
	}
}

// 40 OMIT
