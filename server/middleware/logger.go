package middleware

import (
	"net/http"

	"github.com/alinz/react-native-updater/server/lib/logme"
)

//LogHTTP log every requests coming to server. kind of like access.log
func LogHTTP(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		logme.Info(r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
