package main

import (
  "log"
  "net/http"
  "github.com/DavidToca/simplechat/handlers"
)


func main() {
  r := newRoom()
  handler := &handlers.TemplateHandler{Filename: "chat.html"}
  http.Handle("/", handler)
  http.Handle("/room", r)
  go r.run()
  // start web server
  if err := http.ListenAndServe(":8000", nil); err != nil{
    log.Fatal("ListenAndServe:", err)
  }


}
