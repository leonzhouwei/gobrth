package httputil

import (
	"log"
	"net"
	"net/http"
)

func SafeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if (r == nil) {
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
