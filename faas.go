package faas

import (
	"net/http"

	"github.com/cybora/shipping_go/handlers/rest"
	"github.com/cybora/shipping_go/translation"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	service := translation.NewStaticService()
	handler := rest.NewTranslateHandler(service)
	handler.TranslateHandler(w, r)
}
