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

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", Root).Methods("GET")
	r.HandleFunc("/fsl_run", FSLRun).Methods("POST")
	return r
}

// Root is used for a health check while deployed on the cluster
func Root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// FSLRun is a primary endpoint designed to accept input data and process it.
func FSLRun(w http.ResponseWriter, r *http.Request) {
	inputBytes, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonData := string(inputBytes)
	doc, err := internal.ParseInput(&jsonData)
	if err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	doc.Process() // TODO: return err

	var b bytes.Buffer
	b.WriteString(strings.Join(doc.Output, "\n"))

	w.Header().Add("Content-Type", "text/plain")
	w.Write(b.Bytes())
}
