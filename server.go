package main

import (
	"fmt"
	"github.com/itimofeev/http-speed-test/humanize"
	"io"
	"log"
	"net/http"
	"time"
)

type SpeedLogFunc func(duration time.Duration, size uint64)

func GetHandler(logFunc SpeedLogFunc) http.Handler {
	handler := http.NewServeMux()

	handler.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		size := getIntValue(r, "size", 10*1024*1024)
		randString := RandString(size)

		startTime := time.Now()
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)

		_, err := io.WriteString(w, randString)
		if err != nil {
			log.Println(err)
		}

		logFunc(time.Now().Sub(startTime), uint64(size))
	})
	return handler
}

func runServer(listenAddr string) {
	if err := http.ListenAndServe(listenAddr, GetHandler(logTime)); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func logTime(duration time.Duration, size uint64) {
	fmt.Printf("served %s by %s, speed %s\n", humanize.IBytes(size), duration, formatSpeed(size, duration))
}
