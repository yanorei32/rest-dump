package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Handler struct {
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	w.WriteHeader(http.StatusOK)

	log.Println(r.Method, r.URL)

	if err != nil {
		log.Println("failed to get body:", err.Error())
		return
	}

	if len(body) == 0 {
		log.Println("empty request")
		return
	}

	f := time.Now().Format("2006-01-02_15-04-05.000") + ".json"

	log.Println("write to", f)

	if err := ioutil.WriteFile(f, body, 0644); err != nil {
		log.Println("failed to save body:", err.Error())
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("./rest-dump [port]")
	}


	s := &http.Server{
		Addr:           os.Args[1],
		Handler:        &Handler{},
		ReadTimeout:    time.Second,
		WriteTimeout:   time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
