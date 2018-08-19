package main

import (
	"fmt"
	"net/http"
"time"
)

func init() {
	fmt.Println("File Service inited!")
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("filesDir"))
	pics := http.FileServer(http.Dir("picturesDir"))

	mux.Handle("/files/", http.StripPrefix("/files/", files))
	mux.Handle("/pics/", http.StripPrefix("/pics/", pics))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "67.218.142.250:8989",
		Handler: mux,
	}

	server.ListenAndServe()
}

var count = 0

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format(time.UnixDate),"from",r.RemoteAddr)
	count += 1
	fmt.Println("page view:", count)
}
