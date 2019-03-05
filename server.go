package speedt

import (
	"crypto/rand"
	"fmt"
	"github.com/itimofeev/http-speed-test/humanize"
	"io"
	"log"
	"net/http"
	"time"
)

type SpeedLogFunc func(duration time.Duration, size uint64, speed string)

func newRandomReader(size uint64) *randomReader {
	if size <= 0 {
		log.Panic("size <= 0")
	}
	return &randomReader{
		size: size,
	}
}

type randomReader struct {
	size        uint64
	alreadyRead uint64
}

func (r *randomReader) Read(p []byte) (n int, err error) {
	remaining := r.size - r.alreadyRead
	if remaining == 0 {
		return 0, io.EOF
	}

	bufSize := len(p)
	toReadSize := bufSize
	if remaining < uint64(bufSize) {
		toReadSize = int(remaining)
	}

	limitReader := io.LimitReader(rand.Reader, int64(toReadSize))
	readSize, err := limitReader.Read(p)

	if readSize != toReadSize {
		log.Fatalf("read size %d not equal to toReadSize %d", readSize, toReadSize)
	}
	r.alreadyRead += uint64(readSize)
	return readSize, nil
}

func GetHandlerFunc(logFunc SpeedLogFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		size := uint64(getIntValue(r, "size", 10*1024*1024))

		startTime := time.Now()
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)

		_, err := io.Copy(w, newRandomReader(size))
		if err != nil {
			log.Println(err)
		}

		duration := time.Now().Sub(startTime)
		logFunc(duration, size, FormatSpeed(size, duration))
	}
}

func GetHandler(logFunc SpeedLogFunc) http.Handler {
	handler := http.NewServeMux()
	handler.HandleFunc("/download", GetHandlerFunc(logFunc))
	return handler
}

func RunServer(listenAddr string) {
	if err := http.ListenAndServe(listenAddr, GetHandler(logTime)); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func logTime(duration time.Duration, size uint64, speed string) {
	fmt.Printf("served %s by %s, speed %s\n", humanize.IBytes(size), duration, speed)
}
