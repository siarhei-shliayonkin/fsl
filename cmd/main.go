package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/siarhei-shliayonkin/fsl/api"
	"github.com/siarhei-shliayonkin/fsl/internal"
	"github.com/sirupsen/logrus"
)

// TODO: documenting func, var, etc.
// TODO: UT, 80% coverage

func main() {
	//runTestSample()

	tcpPort := flag.Int("port", 8081, "istening port")
	timeout := flag.Duration("timeout", time.Second*5, "read/write timeout")
	verboseLevel := flag.Uint("v", 5, "verbose level")
	flag.Parse()

	logrus.SetLevel(logrus.Level(*verboseLevel))

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", *tcpPort),
		Handler:      api.NewRouter(),
		ReadTimeout:  *timeout,
		WriteTimeout: *timeout,
	}

	log.Fatal(s.ListenAndServe())
}

// for debug: runs sample processing directly w/o http server
func runTestSample() {
	jsonData := `{
		"var1":1,
		"var2":2,
		
		"init": [
		  {"cmd" : "#setup" }
		],
		
		"setup": [
		  {"cmd":"update", "id": "var1", "value":3.5},
		  {"cmd":"print", "value": "#var1"},
		  {"cmd":"#sum", "id": "var1", "value1":"#var1", "value2":"#var2"},
		  {"cmd":"print", "value": "#var1"},
		  {"cmd":"create", "id": "var3", "value":5},
		  {"cmd":"delete", "id": "var1"},
		  {"cmd":"#printAll"}
		],
		
		"sum": [
			{"cmd":"add", "id": "$id", "operand1":"$value1", "operand2":"$value2"}
		],
	  
		"printAll":
		[
		  {"cmd":"print", "value": "#var1"},
		  {"cmd":"print", "value": "#var2"},
		  {"cmd":"print", "value": "#var3"}
		]
	  }`

	doc, err := internal.ParseInput(&jsonData)
	if err != nil {
		fmt.Println(err)
	}
	// doc.PrintDoc()

	doc.Process()
}
