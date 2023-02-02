package api

import (
	"bytes"
	"io"
	"net/http"

	"github.com/siarhei-shliayonkin/fsl/internal"
	"github.com/sirupsen/logrus"
)

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
	for _, s := range doc.Output {
		b.WriteString(s)
	}

	w.Header().Add("Content-Type", "text/plain")
	w.Write(b.Bytes())
}
