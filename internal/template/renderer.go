package template

import (
	"fmt"
	"html/template"
	"io"
	"path/filepath"
)

type Renderer struct {
	templates map[string]*template.Template
}

func NewRenderer() (*Renderer, error) {
	r := &Renderer{
		templates: make(map[string]*template.Template),
	}

	pages := "web/templates/pages/*.html"
	layout := "web/templates/base/*.html"
	partials := "web/templates/partials/*.html"

	layoutFiles, err := filepath.Glob(layout)
	if err != nil {
		return nil, err
	}
	partialFiles, err := filepath.Glob(partials)
	if err != nil {
		return nil, err
	}
	pageFiles, err := filepath.Glob(pages)
	if err != nil {
		return nil, err
	}

	commonFiles := append(layoutFiles, partialFiles...)

	for _, page := range pageFiles {
		name := filepath.Base(page)

		allFiles := append(commonFiles, page)

		ts, err := template.ParseFiles(allFiles...)
		if err != nil {
			return nil, err
		}
		r.templates[name] = ts
	}

	if len(r.templates) == 0 {
		return nil, fmt.Errorf("no templates found in web/templates/pages/")
	}

	return r, nil
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}) error {
	tmpl, ok := r.templates[name]
	if !ok {
		return fmt.Errorf("template %s not found", name)
	}
	return tmpl.ExecuteTemplate(w, "layout", data)
}
