package main

import (
	"fmt"
	"github.com/go-greeting/pkg/setting"
	r "github.com/go-greeting/router"
	"net/http"
)

func main() {
	router := r.InitRouter()

	s := &http.Server{
		Addr:              fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:           router,
		ReadHeaderTimeout: setting.ReadTimeOut,
		WriteTimeout:      setting.WriteTimeOut,
		MaxHeaderBytes:    1 << 20,
	}

	s.ListenAndServe()
}
