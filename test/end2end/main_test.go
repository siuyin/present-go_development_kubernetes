// 10 OMIT
package main

// 20 OMIT
import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/cloudflare/cfssl/log"
	"github.com/siuyin/dflt"
)

// 30 OMIT
func TestEndToEnd(t *testing.T) {
	endp := dflt.EnvString("ENDPOINT", "http://192.168.39.230:31377")
	t.Run("root", func(t *testing.T) { get(endp+"/", t) })
	t.Run("/a", func(t *testing.T) { get(endp+"/a", t) })
}

// 40 OMIT
func get(endp string, t *testing.T) {
	t.Helper() // marks this as a test helper // HL
	t.Logf("connecting to endpoint: %s\n", endp)
	resp, err := http.Get(endp)
	if err != nil {
		t.Fatalf("could not reach endpoint %s: %v", endp, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("bad status code received: %v", resp.StatusCode)
	}
	defer resp.Body.Close()
	t.Logf("status code: %v", resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading body: %v", err)
	}
	if !bytes.Contains(body, []byte("Hello, the time is")) {
		log.Errorf("unexpected body contents: %s\n", body)
	}
	t.Logf("body contents OK")
}

// 50 OMIT
