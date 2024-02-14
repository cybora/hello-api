package translation_test

import (
	"testing"

	"github.com/cybora/shipping_go/translation"
)

func TestTranslate(t *testing.T) {

	tt := []struct {
		Word        string
		Language    string
		Translation string
	}{
		{
			Word:        "hello",
			Language:    "english",
			Translation: "hello",
		}, {
			Word:        "hello",
			Language:    "german",
			Translation: "hallo",
		}, {
			Word:        "hello",
			Language:    "finnish",
			Translation: "hei",
		}, {
			Word:        "hello",
			Language:    "dutch",
			Translation: "",
		},
	}

	for _, test := range tt {
		res := translation.Translate(test.Word, test.Language)
		if res != test.Translation {
			t.Errorf("Expected %s but got %s", test.Translation, res)
		}
	}
}
