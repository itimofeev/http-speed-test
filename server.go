package main

import (
	"fmt"
	"github.com/itimofeev/http-speed-test/humanize"
	"io"
	"log"
	"net/http"
	"time"
)

func runServer() {
	handler := http.NewServeMux()

	handler.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer logTime(startTime)

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)

		size := getIntValue(r, "size", 10*1024*1024)
		randString := RandString(size)

		_, _ = io.WriteString(w, randString)
	})

	if err := http.ListenAndServe(":13579", handler); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func logTime(startTime time.Time) {
	fmt.Printf("diff: %s, %s\n", time.Now().Sub(startTime), humanize.Bytes(1024*1024*10))
}
