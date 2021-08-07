package rest

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/borkod/Continuous-Delivery-in-Go-Translation-API/translation"
)

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	language := r.URL.Query().Get("language")
	if language == "" {
		language = "english"
	}

	word := filepath.Base(r.URL.Path)
	translation := translation.Translate(word, language)
	if translation == "" {
		language = ""
		w.WriteHeader(404)
	}
	resp := Resp{
		Language:    language,
		Translation: translation,
	}
	if err := enc.Encode(resp); err != nil {
		panic("unable to encode response")
	}
}
