package main

import (
	"backend/internal/soxie/config"
	"backend/internal/soxie/soxierabbit"
	"backend/pkg/rabbit"
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
		logrus.Fatalf("hats.main() rabbit err=%#v", err)
	}
	defer rabbitConn.Close()

	// start rabbit listeners
	soxierabbit.Listen(rabbitConn, wsChannel)

	http.HandleFunc("/", serveHome)
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

var wsChannel = make(chan string)

func writer(ws *websocket.Conn) {
	logrus.Info("soxie.writer() WWWWWW")
	for {
		select {
		case msg := <-wsChannel:
			if err := ws.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				return
			}
		}
		// TODO: add ping ticker and flush target mapping on no response
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	logrus.Info("soxie.serveWs() SSSSSS")

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorf("soxie.serveWs() handshake err=%#v", err)
		return
	}

	target := r.FormValue("target")
	if target == "" {
		return
	}
	logrus.Infof("soxie.serveWs() target=%s", target)

	go writer(ws)
	reader(ws)
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
        <pre id="fileData"></pre>
        <script type="text/javascript">
            (function() {
                var data = document.getElementById("fileData");
                var conn = new WebSocket("ws://{{.Host}}/ws?target={{.Target}}");
                conn.onclose = function(evt) {
                    data.textContent = 'Connection closed';
                }
                conn.onmessage = function(evt) {
                    console.log('file updated');
                    data.textContent = evt.data;
                }
            })();
        </script>
    </body>
</html>
`
