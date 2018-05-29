package main

import (
  "log"
  "net/http"
  "github.com/DavidToca/simplechat/templatehandler"
)


func main() {
  handler := &templatehandler.TemplateHandler{Filename: "chat.html"}
  http.Handle("/", handler)
  // start web server
  if err := http.ListenAndServe(":8000", nil); err != nil{
    log.Fatal("ListenAndServe:", err)
  }


}
