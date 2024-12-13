package handlers

import (
	"encoding/json"
	"net/http"

	"errors"

	"github.com/OctaneAL/Shortly/internal/db"
	"github.com/OctaneAL/Shortly/internal/util"
	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func Decode(w http.ResponseWriter, r *http.Request) {
	type request struct {
		URL string `json:"url"`
	}

	req := request{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if req.URL == "" {
		ape.RenderErr(w, problems.BadRequest(errors.New("missing 'url' field"))...)
		return
	}

	shortened := util.HashAndConvert(req.URL)

	// Error
	// undefined: db.SaveURL
	// I need to somehow retrieve connection created
	// in service/main.go newService() function
	err := db.SaveURL(r.Context(), req.URL, shortened)
	if err != nil {
		ape.RenderErr(w, []*jsonapi.ErrorObject{problems.InternalError()}...)
		return
	}

	ape.Render(w, map[string]string{
		"original_url": req.URL,
		"short_code":   shortened,
	})
}
