package faas

import (
	"net/http"

	"github.com/cybora/shipping_go/handlers/rest"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	rest.TranslateHandler(w, r)
}
