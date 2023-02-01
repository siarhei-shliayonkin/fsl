package main

import (
	"fmt"

	"github.com/siarhei-shliayonkin/fsl/internal"
)

func main() {
	// now just runs parser w/ test data sample.
	// TODO: http server

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
