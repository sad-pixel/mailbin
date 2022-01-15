package web

import (
	"fmt"
	"io"
)

type TemplateDTO = map[string]interface{}

func (h *Handlers) RenderTemplate(w io.Writer, templateName string, data interface{}) error {
	tpl, ok := h.templates[templateName]
	if !ok {
		return fmt.Errorf("%s is not a valid template", templateName)
	}

	if err := tpl.Execute(w, data); err != nil {
		return fmt.Errorf("could not execute template %s: %s", templateName, err.Error())
	}

	return nil
}
