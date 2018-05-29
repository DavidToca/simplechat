package main

import (
  "log"
  "net/http"
  "text/template"
  "path/filepath"
  "sync"
)

type templateHandler struct {
  once sync.Once
  filename string
  templ *template.Template
}

func (handler *templateHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
  handler.once.Do (func(){
      handler.templ = template.Must(
        template.ParseFiles(
            filepath.Join("templates", handler.filename),
        ),
      )
  })
  handler.templ.Execute(w, nil)
}

func render(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte(`
    <html>
    <body>
      <h1>Chat Test</h1>
    </body>
    </html>
  `))
}

func main() {
  handler := templateHandler{filename: "chat.html"}
  handler.HandleFunc("/", handler.ServeHTTP)
  // start web server

  if err := http.ListenAndServe(":8000", nil); err != nil{
    log.Fatal("ListenAndServe:", err)
  }


}
