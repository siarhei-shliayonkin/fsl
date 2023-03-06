package internal

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// integation test for http server

var testSamples = []string{
	`{
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
}`,
	`3.5
5.5
undefined
2
5`,
	`{
	"init": [
	  {"cmd" : "#setup" }
	],

	"setup": [
	  {"cmd":"create", "id": "var1", "value":10},
	  {"cmd":"print", "value": "#var1"},
	  {"cmd":"#sum", "id": "var1", "value1":"#var1", "value2":"#var3"},
	  {"cmd":"#printAll"}
	]
  }`,
	`10
15
2
5`,
}

func Post(t *testing.T, srv *httptest.Server, data string) string {
	res, err := http.Post(srv.URL+baseURL+"/scripts",
		"application/json", strings.NewReader(data),
	)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("status not OK")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	return string(body)
}

func check(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("check body: %+v\nwant:\n%v", got, want)
		t.Fail()
	}
}

func Test_httpServer(t *testing.T) {
	InitGlobals()
	srv := httptest.NewServer(NewRouter())
	defer srv.Close()

	out := Post(t, srv, testSamples[0])
	check(t, out, testSamples[1])

	// cumulative test
	out = Post(t, srv, testSamples[2])
	check(t, out, testSamples[3])
}
