package main

import (
	"log"
	"time"
	"net/http"
	"io/ioutil"
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
	s := &http.Server {
		Addr: ":8080",
		Handler: &Handler{},
		ReadTimeout: time.Second,
		WriteTimeout: time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}

