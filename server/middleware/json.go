package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/zenazn/goji/web"
)

type StructBuilder interface {
	BuildNew() interface{}
}

func JSONDecode(structBuilder StructBuilder) func(c *web.C, h http.Handler) http.Handler {
	return func(c *web.C, h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			to := structBuilder.BuildNew()

			if err := json.NewDecoder(r.Body).Decode(to); err != nil {
				return
			}

			c.Env["parsedBody"] = to
			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}

func JSONEncode(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var response interface{}
		status := http.StatusOK

		if value, ok := c.Env["status"]; ok {
			status = value.(int)
		}

		if value, ok := c.Env["response"]; ok {
			response = value
		}

		writeJSON(w, status, response)

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func writeJSON(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
