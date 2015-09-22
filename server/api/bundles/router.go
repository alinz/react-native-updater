package bundles

import "github.com/pressly/cji"

func New() cji.Router {
	r := cji.NewRouter()

	//middleware section

	//routes
	r.Get("/:version", bundles)

	return r
}
