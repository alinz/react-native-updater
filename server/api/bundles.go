package api

import (
	"net/http"

	"github.com/pressly/cji"
	"github.com/zenazn/goji/web"
)

func bundles(c web.C, w http.ResponseWriter, r *http.Request) {
}

func New() cji.Router {
	r := cji.NewRouter()

	//middleware section

	//routes
	r.Get("/:version", bundles)

	return r
}
