// 10 OMIT
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/siuyin/dflt"
)

// 20 OMIT
func main() {
	endp := dflt.EnvString("ENDPOINT", "http://192.168.39.230:31051/")
	start := time.Now()
	fmt.Println("End to end testing.\n")
	fmt.Printf("getting %s\n", endp)
	resp, err := http.Get(endp)
	if err != nil {
		log.Fatalf("could not reach endpoint: %s: %v", endp, err)
	}

	fmt.Println("checking status code")
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("bad status code received: %v", resp.StatusCode)
	}
	defer resp.Body.Close()
	// 30 OMIT
	fmt.Println("checking response")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading body: %v\n", err)
	}
	if !bytes.Contains(body, []byte("Hello, the time is")) {
		log.Fatalf("unexpected body: %s\n", body)
	}

	fmt.Printf("\ntests complete. Took %g seconds\n", time.Now().Sub(start).Seconds())
}

// 40 OMIT
