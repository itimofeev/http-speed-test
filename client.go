package speedt

import (
	"fmt"
	"github.com/itimofeev/http-speed-test/humanize"
	"io"
	"log"
	"net/http"
	"time"
)

type countingWriter struct {
	writeBytes uint64
}

func (r *countingWriter) Write(p []byte) (int, error) {
	r.writeBytes += uint64(len(p))
	return len(p), nil
}

func RunClient() {
	serverAddress := "localhost:13579"
	size := 10 * 1024 * 1024 * 1024
	start := time.Now()

	resp, err := http.Get(fmt.Sprintf("http://%s/download?size=%d", serverAddress, size))
	if err != nil {
		log.Fatal(err)
	}

	writer := &countingWriter{}
	_, err = io.Copy(writer, resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	duration := time.Now().Sub(start)
	log.Printf("downloaded %s for %s, speed %s", humanize.IBytes(writer.writeBytes), duration, FormatSpeed(writer.writeBytes, duration))
}

func FormatSpeed(bytes uint64, duration time.Duration) string {
	speed := float64(bytes) / duration.Seconds()
	if speed < 1024 {
		return fmt.Sprintf("%f b/s", speed)
	}
	if speed < 1024*1024 {
		return fmt.Sprintf("%f Kib/s", speed/1024)
	}
	if speed < 1024*1024*1024 {
		return fmt.Sprintf("%f Mib/s", speed/1024/1024)
	}
	return fmt.Sprintf("%f Gib/s", speed/1024/1024/1024)
}
