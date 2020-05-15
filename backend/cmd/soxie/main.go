package main

import (
	"backend/internal/hats/hatdao"
	"backend/internal/hats/orderdao"
	"backend/internal/soxie/config"
	"backend/internal/soxie/soxierabbit"
	"backend/pkg/authnz"
	"backend/pkg/rabbit"
	"encoding/json"
	"html/template"
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

	homeTempl = template.Must(template.New("").Parse(homeHTML))
	upgrader  = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
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

	if config.HomePath != "" {
		http.HandleFunc(config.HomePath, serveHome)
	}

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
		_, pong, err := ws.ReadMessage()
		logrus.Infof("soxie.reader() read message pong=%s", pong)
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
func WsHandleJSON(ws *websocket.Conn, q rabbit.Queue, target string, v interface{}) {

	logrus.Infof("soxie.HandleJSON() target=%s", target)

	var m = map[string]interface{}{
		"queue": q.Name(),
		"data":  v,
	}

	// sockets := subMgr.SubjectSocketMap[target]

	// logrus.Infof("soxie.HandleJSON() sockets=%v", sockets)

	msg, err := json.Marshal(m)
	if err != nil {
		logrus.Errorf("HandleJSON() json marshal err=%#v", err)
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
	_, err := authnz.ValidateCookie("id_token", r)
	if err != nil {
		logrus.Errorf("soxie.serveWs() validate request err=%#v", err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// TODO: validate against list of injected origins
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorf("soxie.serveWs() handshake err=%#v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// creates a new routine encapsulating a websocket bound to this subject
	go asynchSocketWriter(ws)

	// reader will block until the socket is closed
	reader(ws)
}

func asynchSocketWriter(ws *websocket.Conn) {
	logrus.Info("soxie.asynchSocketWriter()")

	pingTicker := time.NewTicker(pingPeriod)

	for {
		select {
		case hat := <-hatCreatedChannel:
			// subMgr.HandleHatCreated(msg)
			WsHandleJSON(ws, rabbit.SoxieHatCreatedQ, hat.CreatedBy, hat)
		case order := <-orderCreatedChannel:
			// subMgr.HandleOrderCreated(msg)
			WsHandleJSON(ws, rabbit.SoxieOrderCreatedQ, order.CreatedBy, order)
		case <-pingTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				// the error causes a return from this go routine
				return
			}
		}
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	logrus.Info("soxie.serveHome() HHHHHH")
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	target := r.FormValue("target")
	logrus.Infof("soxie.serveHome() TTT target=%s", target)

	// TODO: get the "target" from a query param
	var v = struct {
		Host   string
		Target string
	}{
		r.Host,
		target,
	}
	homeTempl.Execute(w, &v)
}

const homeHTML = `<!DOCTYPE html>
<html lang="en">
    <head>
        <title>Soxie</title>
    </head>
	<body>
		<h1>Soxie</h1>
        <pre id="fileData"></pre>
        <script type="text/javascript">
            (function() {
                var data = document.getElementById("fileData");
                var conn = new WebSocket("ws://{{.Host}}/ws?target={{.Target}}");
                conn.onclose = function(evt) {
                    data.textContent = 'Connection closed';
                }
                conn.onmessage = function(evt) {
                    console.log(evt.data);
                    data.textContent = evt.data + '\n' + data.textContent;
                }
            })();
        </script>
    </body>
</html>
`
