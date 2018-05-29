package main

import (
  "log"
  "net/http"
)


func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    w.Write([]byte(`
      <html>
      <body>
        <h1>Chat Test</h1>
      </body>
      </html>
    `))
  })

  // start web server

  if err := http.ListenAndServe(":8000", nil); err != nil{
    log.Fatal("ListenAndServe:", err)
  }


}
