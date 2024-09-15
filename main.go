package main

import (
	"fmt"
	"net/http"
	"test/pkg/setting"
	"test/routers"
)

func main() {
	r := routers.InitRoute()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        r,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
