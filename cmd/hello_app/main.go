// 10 OMIT
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/siuyin/dflt"
	myt "github.com/siuyin/present-go_development_kubernetes/time"
)

// 20 OMIT
func main() {
	fmt.Println("hello_app")
	webServer()
	heartBeat()
	select {}
}

// 30 OMIT
func webServer() {
	go func() {
		c := myt.Clock{}
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, the time is %s\n", myt.Now(c)) // my time package // HL
		})
		port := dflt.EnvString("PORT", "8080")
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}()
}

// 40 OMIT
func heartBeat() {
	go func() {
		for {
			c := myt.Clock{}
			fmt.Println(myt.Now(c))
			time.Sleep(5 * time.Second) // go's time package // HL
		}
	}()

}

// 50 OMIT
