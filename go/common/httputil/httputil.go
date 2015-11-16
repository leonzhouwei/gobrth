package httputil

import (
	"log"
	"net"
	"net/http"

	"github.com/leonzhouwei/llsn/go/common"
)

func SafeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r == nil {
				return
			}

			if err, ok := r.(net.Error); ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				// 或者输出自定义的 50x 错误页面
				// w.WriteHeader(http.StatusInternalServerError)
				// renderHtml(w, "error", e)
				// logging
				log.Println("WARN: panic in %v - %v", fn, err)
			} else {
				log.Println("WARN: panic in %v - %v", fn, r)
			}
		}()
		fn(w, r)
	}
}

func StaticDirHandler(mux *http.ServeMux, prefix string, staticDir string) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		filePath := staticDir + r.URL.Path[len(prefix)-1:]
		if exists := common.FileExist(filePath); !exists {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, filePath)
	})
}
