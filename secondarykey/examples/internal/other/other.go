package other

import (
	"net/http"

	. "mypkgs/handler/internal"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//Error
	w.Write(Hello("Other", r.URL))
}
