package download

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

func downloadBundle(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("requested all releases")))
}
