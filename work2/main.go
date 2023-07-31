/*
* 命令运行 go run main.go -alsologtostderr true
* https://github.com/golang/glog
 */
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/golang/glog"
)

// 状态码设置
const Stats = 200

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", responseHeader)
	mux.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}
}

// health check
func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(Stats)
}

// response
func responseHeader(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]any, 2)
	m["clientIp"] = clientIP(r)
	m["status"] = Stats
	slog, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	flag.Parse()
	defer glog.Flush()
	glog.Infoln(string(slog))

	//response header
	for k, v := range r.Header {
		w.Header().Add(k, fmt.Sprintf("%s", v))
	}
	//add VERSION
	w.Header().Add("VERSION", os.Getenv("VERSION"))

	//response body
	io.WriteString(w, "ok")
}

func clientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
