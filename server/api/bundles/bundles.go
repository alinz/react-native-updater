package bundles

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

func bundles(c web.C, w http.ResponseWriter, r *http.Request) {
	version := c.URLParams["version"]
	w.Write([]byte(fmt.Sprintf("requested %s !", version)))
}
