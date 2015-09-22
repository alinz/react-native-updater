package releases

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

func releases(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("requested all releases")))
}
