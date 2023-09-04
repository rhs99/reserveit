package handlers

import (
	"net/http"

	"github.com/rhs99/reserveit/pkg/config"
	"github.com/rhs99/reserveit/pkg/models"
	"github.com/rhs99/reserveit/pkg/render"
)


type Repository struct{
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig)*Repository{
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository){
	Repo = r
}

func (m *Repository)Home(w http.ResponseWriter, r *http.Request){
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository)About(w http.ResponseWriter, r *http.Request){
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: stringMap})
}
