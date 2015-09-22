package api

import (
	"net/http"

	"github.com/pressly/cji"
	"github.com/zenazn/goji/web"
)

func releases(c web.C, w http.ResponseWriter, r *http.Request) {
}

func New() cji.Router {
	r := cji.NewRouter()

	return r
}
