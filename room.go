package main

import (
  "log"
  "github.com/gorilla/websocket"
  "net/http"
)

const (
  socketBufferSize = 1024
  messageBufferSize = 256
)

type Room struct {
  forward chan []byte
  join chan *client
  leave chan *client
  clients map[*client]bool
}

// newRoom makes a new Room.
func newRoom() *Room {
  return &Room{
    forward: make(chan []byte),
    join: make(chan *client),
    leave: make(chan *client),
    clients: make(map[*client]bool),
  }
}

func (r *Room) run() {
  for {
    select {
      case client := <-r.join:
        //joining
        r.clients[client] = true
      case client := <-r.leave:
        //leaving
        delete(r.clients, client)
        close(client.send)
      case msg := <-r.forward:
        for client := range r.clients {
          client.send <- msg
        }
    }
  }
}


var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize,
  WriteBufferSize: socketBufferSize}

func (r *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    socket, err := upgrader.Upgrade(w, req, nil)
    if err != nil {
      log.Fatal("ServeHTTP:", err)
      return
    }
    client := &client{
      socket: socket,
      send: make(chan []byte, messageBufferSize),
      Room: r,
    }
    r.join <- client
    defer func() { r.leave <- client } ()
    go client.write()
    client.read()
}
