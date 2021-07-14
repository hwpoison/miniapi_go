package server 

import (
	"github.com/gorilla/mux"
	"net/http"	
	"fmt"
	"log"
	"strconv"
	"miniApi/pkg"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}


func New() Server {
	log.Println("Listen...")
	a := &api{}
	r := mux.NewRouter()
	r.HandleFunc("/perros/{ID:[a-zA-Z0-9_]+}", a.getDog).Methods(http.MethodGet)
	a.router = r 
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) getDog(w http.ResponseWriter, r *http.Request) {
	if f, ok := w.(http.Flusher); ok { 
		f.Flush() 
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["ID"])	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "<html><title>ERROR</title><h2 style='color:red'>internal error!</h2></html>")
		return 
	}

	n := fakedatabase.GetDogById(id)
	if n == ""{
		fmt.Fprintf(w, "not found!")
	}
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, string(n))
}


