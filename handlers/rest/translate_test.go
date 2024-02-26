package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cybora/shipping_go/handlers/rest"
)

type stubbedService struct{}

func (s stubbedService) Translate(word string, language string) string {
	if word == "foo" {
		return "bar"
	}
	return ""
}

func TestTranslateAPI(t *testing.T) {
	tt := []struct {
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			Endpoint:            "/foo",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "bar",
		},
		{
			Endpoint:            "/foo?language=german",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "german",
			ExpectedTranslation: "bar",
		}, {
			Endpoint:            "/baz",
			StatusCode:          http.StatusNotFound,
			ExpectedLanguage:    "",
			ExpectedTranslation: "",
		},
	}

	underTest := rest.NewTranslateHandler(&stubbedService{})
	handler := http.HandlerFunc(underTest.TranslateHandler)

	for _, test := range tt {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, test.Endpoint, nil)
		handler.ServeHTTP(rr, req)

		if rr.Code != test.StatusCode {
			t.Errorf("expected status %d but got %d", test.StatusCode, rr.Code)
		}

		var resp rest.Resp

		err := json.Unmarshal(rr.Body.Bytes(), &resp)
		if err != nil {
			t.Logf("err occurred")
		}

		if resp.Language != test.ExpectedLanguage {
			t.Errorf("expected language %s but got %s", test.ExpectedLanguage, resp.Language)
		}

		if resp.Translation != test.ExpectedTranslation {
			t.Errorf("expected translation %s but got %s", test.ExpectedTranslation, resp.Translation)
		}
	}
}
