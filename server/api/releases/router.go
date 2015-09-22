package releases

import "github.com/pressly/cji"

//New creates all /releases routes
func New() cji.Router {
	r := cji.NewRouter()

	//middleware section

	//routes
	r.Get("/", releases)

	return r
}
