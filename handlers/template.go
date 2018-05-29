package handlers

import (
  "net/http"
  "text/template"
  "path/filepath"
  "sync"
)

type TemplateHandler struct {
  once sync.Once
  Filename string
  templ *template.Template
}

func (handler *TemplateHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
  handler.once.Do (func(){
      handler.templ = template.Must(
        template.ParseFiles(
            filepath.Join("templates", handler.Filename),
        ),
      )
  })
  handler.templ.Execute(w, nil)
}
