package handlers

import (
	"intelchaos/pkg/config"
	"intelchaos/pkg/models"
	"intelchaos/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo( a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository){
	Repo = r
}

//Home is the home page Handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request)  {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl",&models.TemplateData{})
}

//About is the home page Handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request)  {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again!"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl",&models.TemplateData{
		StringMap: stringMap,
	})
}

