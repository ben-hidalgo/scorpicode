package main

import (
	"backend/internal/hats/hatdao"
	"backend/internal/hats/orderdao"
	"backend/internal/soxie/config"
	"backend/internal/soxie/soxierabbit"
	"backend/pkg/authnz"
	"backend/pkg/rabbit"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var (
	// Time allowed to write the file to the client.
	writeWait = time.Duration(config.WriteWaitSeconds) * time.Second

	// Time allowed to read the next pong message from the client.
	pongWait = time.Duration(config.PongWaitSeconds) * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// TODO: validate against list of injected origins
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

func main() {
	logrus.Infof("soxie.main() starting")

	// connect rabbit
	rabbitConn, err := rabbit.Connect()
	if err != nil {
		logrus.Fatalf("soxie.main() rabbit err=%#v", err)
	}
	defer rabbitConn.Close()

	// start rabbit listeners
	soxierabbit.Listen(rabbitConn, soxierabbit.Channels{
		HatCreatedChannel:   hatCreatedChannel,
		OrderCreatedChannel: orderCreatedChannel,
	})

	http.HandleFunc("/ws", serveWs)
	logrus.Infof("soxie.main() listening on %s", config.ListenAddress)
	if err := http.ListenAndServe(config.ListenAddress, nil); err != nil {
		log.Fatal(err)
	}
}

func reader(ws *websocket.Conn) {
	logrus.Info("soxie.reader()")
	defer ws.Close()
	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error {
		logrus.Info("soxie.PongHandler() pong received")
		ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		// ignore the message
		_, _, err := ws.ReadMessage()
		if err != nil {
			logrus.Infof("soxie.reader() read message err=%#v", err)
			// breaking out of this infinite loop allows this func to return which
			// triggers the defer websocket close
			break
		}
	}
}

var hatCreatedChannel = make(chan hatdao.Hat)
var orderCreatedChannel = make(chan orderdao.Order)

// WsHandleJSON .
func WsHandleJSON(ws *websocket.Conn, q rabbit.Queue, v interface{}) {

	// this map envelope allows the client to use queue name as a message type switch
	var m = map[string]interface{}{
		"queue": q.Name(),
		"data":  v,
	}

	msg, err := json.Marshal(m)
	if err != nil {
		logrus.Errorf("HandleJSON() json marshal err=%#v", err)
		return
	}

	err = ws.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		if err == websocket.ErrCloseSent {
			// TODO: remove socket from array / map
		} else {
			logrus.Errorf("HandleJSON() write message err=%#v", err)
		}

	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	logrus.Info("soxie.serveWs()")

	// TODO: cookie name is hard coded
	bearer, err := authnz.ValidateCookie("id_token", r)
	if err != nil {
		logrus.Errorf("soxie.serveWs() validate request err=%#v", err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorf("soxie.serveWs() handshake err=%#v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// creates a new routine encapsulating a websocket bound to this subject
	go asynchSocketWriter(ws, bearer.GetSubject())

	// reader will block until the socket is closed
	reader(ws)
}

func asynchSocketWriter(ws *websocket.Conn, subject string) {
	logrus.Info("soxie.asynchSocketWriter() subject=%s", subject)

	pingTicker := time.NewTicker(pingPeriod)

	for {
		select {
		case hat := <-hatCreatedChannel:
			logrus.Infof("soxie.asynchSocketWriter() hat.CreatedBy=%s", hat.CreatedBy)
			if hat.CreatedBy == subject {
				WsHandleJSON(ws, rabbit.SoxieHatCreatedQ, hat)
			} else {
				logrus.Errorf("soxie.asynchSocketWriter() hat WTF hat.CreatedBy=%s", hat.CreatedBy)
			}
		case order := <-orderCreatedChannel:
			logrus.Infof("soxie.asynchSocketWriter() order.CreatedBy=%s", order.CreatedBy)
			if order.CreatedBy == subject {
				WsHandleJSON(ws, rabbit.SoxieOrderCreatedQ, order)
			} else {
				logrus.Errorf("soxie.asynchSocketWriter() order WTF order.CreatedBy=%s", order.CreatedBy)
			}
		case <-pingTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				// the error causes a return from this go routine, allowing garbage collection
				return
			}
		}
	}
}
