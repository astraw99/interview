package main

import (
	"log"
	"net/http"
)

func main()  {
	mux := http.NewServeMux()

	rh := http.RedirectHandler("https://www.baidu.com", 302)
	mux.Handle("/foo", rh)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
