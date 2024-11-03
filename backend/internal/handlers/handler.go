package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Img  string `json:"img"`
}

func HandlerFuncWS(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	log.Println("New connection")
	if err != nil {
		log.Println("accept:", err)
		return
	}
	defer c.CloseNow()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var v Item
	err = wsjson.Read(ctx, c, &v)
	if err != nil {
		log.Println("read:", err)
		return
	}
	log.Printf("received: %v", v)

	c.Close(websocket.StatusNormalClosure, "")
}