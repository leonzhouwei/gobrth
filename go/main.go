package main

import (
	"io"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/leonzhouwei/llsn/go/conf"
)

func main() {
	addr := ":" + strconv.Itoa(conf.Port())
	fmt.Println(addr)
	
	http.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	})
	log.Fatal(http.ListenAndServe(addr, nil))
	
}
