package business

import (
	"net/http"

	"golang-api-testing/utils"
)

func NoRoute(w http.ResponseWriter, r *http.Request) {
	// return a status ok
	w.Write(utils.ProcessBadResponse(utils.NotFoundCode, utils.NotFound, []string{"This route has not been implemented"}))
}
