package main

import (
  "github.com/gorilla/websocket"
)

type client struct {
  socket *websocket.Conn
  send chan []byte
  Room *Room
}

func (c *client) read() {
  defer c.socket.Close()
  for {
    _, msg, err := c.socket.ReadMessage()
    if err != nil{
      return
    }
    c.Room.forward <- msg
  }
}

func (c *client) write() {
  defer c.socket.Close()
  for msg := range c.send {
    err := c.socket.WriteMessage(websocket.TextMessage, msg)
    if err != nil {
      return
    }
  }
}
