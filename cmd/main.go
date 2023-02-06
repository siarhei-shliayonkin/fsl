package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/siarhei-shliayonkin/fsl/api"
)

func main() {
	tcpPort := flag.Int("port", 8081, "listening port")
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
