package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/siarhei-shliayonkin/fsl/api"
	"github.com/siarhei-shliayonkin/fsl/internal"
)

// TODO: documenting func, var, etc.
// TODO: UT

func main() {
	//runTestSample()

	tcpPort := flag.Int("port", 8081, "Listening port.")
	flag.Parse()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", *tcpPort),
		Handler:        api.NewRouter(),
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
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
