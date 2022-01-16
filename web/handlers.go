package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gust1n/go-render/render"
	"github.com/sad-pixel/mailbin/repository"
)

type WebSettings struct {
	Port string
}

type Handlers struct {
	templates  map[string]*template.Template
	repository *repository.EmailRepository
}

func NewHandlers(repo *repository.EmailRepository) (*Handlers, error) {
	h := &Handlers{repository: repo}

	templates, err := render.Load("templates")
	h.templates = templates

	if err != nil {
		return nil, fmt.Errorf("could not load templates: %w", err)
	}

	return h, nil
}

func (h *Handlers) GetIndex(w http.ResponseWriter, r *http.Request) {
	messages := ToDisplayEmails(h.repository.GetAll())
	err := h.RenderTemplate(w, "index.html", map[string]interface{}{
		"messages": messages,
	})
	if err != nil {
		log.Println(err)
	}
}
