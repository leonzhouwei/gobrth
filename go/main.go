package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/leonzhouwei/llsn/go/common"
	"github.com/leonzhouwei/llsn/go/common/httputil"
	"github.com/leonzhouwei/llsn/go/common/tplutil"
	"github.com/leonzhouwei/llsn/go/conf"
)

var config conf.Config

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == common.GET {
		err := tplutil.RenderHtml(w, config, "index")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

}

func main() {
	config = conf.NewConfig()

	addr := ":" + strconv.Itoa(config.Port())
	fmt.Println(addr)

	mux := http.NewServeMux()
	httputil.StaticDirHandler(mux, "/public/", "./public")
	mux.HandleFunc("/", httputil.SafeHandler(indexHandler))
	log.Fatal(http.ListenAndServe(addr, mux))

}
