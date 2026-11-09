package main

import (
	"net/http"
)

type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the server"))
}
func main() {
	api := &api{addr: ":8080"}
	mux:=http.NewServeMux()
	srv:=&http.Server{
		Addr: api.addr,
		Handler:mux,
	}
	mux.HandleFunc("GET /users",api.getUsersHandler)
	mux.HandleFunc("POST /user",api.createUsersHandler)

	
     err:=srv.ListenAndServe()
	 if err!=nil{
		panic(err)
	 }
}
