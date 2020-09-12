package admin

import (
	"net/http"

	. "mypkgs/handler/internal"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(Hello("Administrator", r.URL))
}
