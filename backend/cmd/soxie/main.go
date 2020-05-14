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
	logrus.Info("soxie.reader() RRRRRRR")
	defer ws.Close()
	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error {
		logrus.Info("soxie.PongHandler() PPPPPP")
		ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		logrus.Info("soxie.reader() read message MMMMMM")
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}

var hatCreatedChannel = make(chan hatdao.Hat)
var orderCreatedChannel = make(chan orderdao.Order)

// SubMgr manages substrictions to target channels
type SubMgr struct {
	// maps a subject (user id) to the array of websockets subscribing to it (array for multiple tabs / browsers)
	SubjectSocketMap map[string][]*websocket.Conn
}

var subMgr = &SubMgr{
	SubjectSocketMap: make(map[string][]*websocket.Conn),
}

// Subscribe add the subscription
func (sm *SubMgr) Subscribe(target string, ws *websocket.Conn) {

	logrus.Infof("soxie.Subscribe() target=%s", target)

	// TODO: ??? prevent double append by Header: 'Sec-WebSocket-Key: zAKVwwGXWAH6qtt5TgzYXA=='
	subMgr.SubjectSocketMap[target] = append(subMgr.SubjectSocketMap[target], ws)
}

// HandleHatCreated write the message to all web sockets subscribed by user id (CreatedBy)
func (sm *SubMgr) HandleHatCreated(hat hatdao.Hat) {
	sm.HandleJSON(rabbit.SoxieHatCreatedQ, hat.CreatedBy, hat)
}

// HandleOrderCreated write the message to all web sockets subscribed by user id (CreatedBy)
func (sm *SubMgr) HandleOrderCreated(order orderdao.Order) {
	sm.HandleJSON(rabbit.SoxieOrderCreatedQ, order.CreatedBy, order)
}

// HandleJSON write message to all web sockets subscribed to the implied target
func (sm *SubMgr) HandleJSON(q rabbit.Queue, target string, v interface{}) {

	logrus.Infof("soxie.HandleJSON() target=%s", target)

	var m = map[string]interface{}{
		"queue": q.Name(),
		"data":  v,
	}

	sockets := subMgr.SubjectSocketMap[target]

	logrus.Infof("soxie.HandleJSON() sockets=%v", sockets)

	for _, ws := range sockets {

		msg, err := json.Marshal(m)
		if err != nil {
			logrus.Errorf("HandleJSON() json marshal err=%#v", err)
		}

		if err := ws.WriteMessage(websocket.TextMessage, msg); err != nil {
			logrus.Errorf("HandleJSON() write message err=%#v", err)
		}
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	logrus.Info("soxie.serveWs() SSSSSS")

	// TODO: cookie name is hard coded
	bearer, err := authnz.ValidateCookie("id_token", r)
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

	subject := bearer.GetSubject()
	logrus.Infof("soxie.serveWs() subject=%s", subject)

	// this is synchronous because the subscription must exist before the channel receives
	subMgr.Subscribe(subject, ws)

	// creates a new routine which will create a socket subject pair and receive messages by channel
	go asynchSocketWriter(ws, subject)

	reader(ws)
}

func asynchSocketWriter(ws *websocket.Conn, target string) {
	logrus.Info("soxie.asynchSocketWriter() WWWWWW")

	for {
		select {
		case msg := <-hatCreatedChannel:
			subMgr.HandleHatCreated(msg)
		case msg := <-orderCreatedChannel:
			subMgr.HandleOrderCreated(msg)
		}
		// TODO: add ping ticker and flush target mapping on no response
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
