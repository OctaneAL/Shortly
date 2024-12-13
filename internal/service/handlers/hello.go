package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	ape.Render(w, map[string]string{"message": "Hello, World!"})
}
