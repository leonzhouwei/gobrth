package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/leonzhouwei/llsn/go/common"
	"github.com/leonzhouwei/llsn/go/common/tplutil"
	"github.com/leonzhouwei/llsn/go/conf"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == common.GET {
		err := tplutil.RenderHtml(w, "index")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

}

func main() {
	addr := ":" + strconv.Itoa(conf.Port())
	fmt.Println(addr)

	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(addr, nil))

}
