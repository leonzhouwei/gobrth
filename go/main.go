package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/leonzhouwei/llsn/go/common"
	"github.com/leonzhouwei/llsn/go/conf"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == common.GET {
		filePath := filepath.Join(conf.PublicDirHome(), "index.html")
		t, err := template.ParseFiles(filePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
		return
	}
}

func main() {
	addr := ":" + strconv.Itoa(conf.Port())
	fmt.Println(addr)

	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(addr, nil))

}
