package business

import (
	"log"
	"net/http"

	"github.com/bs/a-jumbo-backend-test/utils"
)

func NoRoute(w http.ResponseWriter, r *http.Request) {
	// return a status ok
	log.Print("No route")
	w.Write(utils.ProcessResponse(utils.StatusOKCode, utils.StatusOK, []string{"This route has not been implemented"}))
}
