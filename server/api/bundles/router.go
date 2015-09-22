package bundles

import "github.com/pressly/cji"

//New creates all /bundles routes
func New() cji.Router {
	r := cji.NewRouter()

	//middleware section

	//routes
	r.Get("/:version", bundles)

	return r
}
