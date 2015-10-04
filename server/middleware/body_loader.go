package middleware

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/zenazn/goji/web"
)

func BodyLoader(maxSize int64) func(c *web.C, h http.Handler) http.Handler {
	return func(c *web.C, h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			body, err := ioutil.ReadAll(io.LimitReader(r.Body, maxSize))
			defer r.Body.Close()

			if err != nil {
				writeJSON(w, 222, err)
				return
			}

			c.Env["rawBody"] = body
			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
