package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
)

// Proxy is the reverse proxy instance
type Proxy struct {
	HostPrefix string
	FromPath   string
	ToPath     string
	url        *url.URL
	proxy      *httputil.ReverseProxy
}

// GetHandle .
func (p *Proxy) GetHandle() func(w http.ResponseWriter, r *http.Request) {
	url, _ := url.Parse(p.HostPrefix)

	p.url = url
	p.proxy = httputil.NewSingleHostReverseProxy(url)

	return p.handle
}

// func (p *Proxy) String() string {
// 	return fmt.Sprintf("%s%s => %s", p.HostPrefix, p.FromPath, p.ToPath)
// }

// Handle routes the request to the proxy and cracks JWT
func (p *Proxy) handle(w http.ResponseWriter, r *http.Request) {

	logrus.Infof("handle() %s %s", r.Method, r.URL.Path)

	// all in one place
	if r.Method == "OPTIONS" {
		enableCors(&w)
		return
	}

	enableCors(&w)

	r.Header.Add("X-Forwarded-Host", r.Host)

	if strings.HasPrefix(r.URL.Path, p.FromPath) {
		r.URL.Path = strings.ReplaceAll(r.URL.Path, p.FromPath, p.ToPath)
	}

	logrus.Debugf("handle() serving %s", r.URL.Path)
	p.proxy.ServeHTTP(w, r)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")

}
