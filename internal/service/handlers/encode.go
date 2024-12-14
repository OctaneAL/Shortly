package handlers

import (
	"encoding/json"
	"net/http"

	"errors"

	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func Encode(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Shortened string `json:"shortened_url"`
	}

	req := request{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if req.Shortened == "" {
		ape.RenderErr(w, problems.BadRequest(errors.New("missing 'shortened_url' field"))...)
		return
	}

	database := DB(r.Context())

	originalURL, err := database.GetOriginalURL(r.Context(), req.Shortened)

	if err != nil {
		ape.RenderErr(w, []*jsonapi.ErrorObject{problems.InternalError()}...)
		return
	}

	ape.Render(w, map[string]string{
		"original_url": originalURL.OriginalURL,
	})
}
