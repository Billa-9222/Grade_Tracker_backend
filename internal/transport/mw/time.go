package mw

import (
	"fmt"
	"net/http"
	"time"
)

func Time(new http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		fmt.Println("start", time.Now().Unix())
		new.ServeHTTP(w, r)
		fmt.Println("end", time.Now().Unix())
	})
}
