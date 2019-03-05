package speedt

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func RandString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func getIntValue(r *http.Request, paramName string, defaultValue int) int {
	value := strings.TrimSpace(r.FormValue(paramName))
	if len(value) == 0 {
		return defaultValue
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return v
}
