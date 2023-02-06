package api

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/siarhei-shliayonkin/fsl/internal"
)

const baseURL = "/fsl/v1"

// NewRouter returns handles router
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(baseURL, Live).Methods("GET")
	r.HandleFunc(baseURL+"/scripts", Script).Methods("POST")
	return r
}

// Live is used for a health check while deployed on the cluster
func Live(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// Script is a primary endpoint designed to accept input data and process it.
func Script(w http.ResponseWriter, r *http.Request) {
	inputBytes, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonData := string(inputBytes)
	doc, err := internal.ParseInput(jsonData)
	if err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	doc.Process()

	var b bytes.Buffer
	b.WriteString(strings.Join(doc.Output, "\n"))

	w.Header().Add("Content-Type", "text/plain")
	w.Write(b.Bytes())
}

// TODO: type ResponseMessage
