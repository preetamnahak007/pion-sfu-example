package signal

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// getHost tries its best to return the request host.
func getHost(r *http.Request) string {
	if r.URL.IsAbs() {
		host := r.Host
		// Slice off any port information.
		if i := strings.Index(host, ":"); i != -1 {
			host = host[:i]
		}
		return host
	}
	return r.URL.Host
}

// HTTPSDPServer starts a HTTP Server that consumes SDPs
func HTTPSDPServer() chan string {
	port := flag.Int("port", 8080, "port of http server")
	dir := flag.String("dir", "static", "base directory of file server")
	flag.Parse()

	sdpChan := make(chan string)
	http.HandleFunc("/sdp", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("/sdp connectedi from %s", r.Host)
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(w, "done")
		sdpChan <- string(body)
	})

	// http server for static files
	fs := http.FileServer(http.Dir(*dir))
	http.Handle("/"+*dir+"/", http.StripPrefix("/"+*dir+"/", fs))

	go func() {
		err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)
		if err != nil {
			panic(err)
		}
	}()

	return sdpChan
}
