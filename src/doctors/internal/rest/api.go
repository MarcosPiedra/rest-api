package rest

import (
	"embed"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type (
	Controller interface {
		Register(router chi.Router)
	}
	DoctorsV1 interface {
		Controller
	}

	Api struct {
		mux       *chi.Mux
		doctorsV1 DoctorsV1
	}
)

//go:embed swagger.json
var swagger embed.FS

// @title           Doctor search engine
// @version         1.0

// @contact.name   Marcos
// @contact.email  piedra.osuna@gmail.com

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func NewApi(
	mux *chi.Mux,
	doctorsV1 DoctorsV1,
) *Api {

	return &Api{
		mux:       mux,
		doctorsV1: doctorsV1,
	}
}

func (api *Api) Init() {
	const specRoot = "/swagger/doctors"

	// API version 1
	api.mux.Route("/doctors/v1", func(r chi.Router) {
		api.doctorsV1.Register(r)
	})

	api.mux.Mount(specRoot, http.StripPrefix(specRoot, http.FileServer(http.FS(swagger))))
}
