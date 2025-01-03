package engine

import (
	"encoding/json"
	"net/http"
)

type Server struct {
	Response http.ResponseWriter
	Request  http.Request
}

func (S *Server) hello_world(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(S.Response).Encode("Hello World")

}

func (S *Server) Start() {
	S.handler()
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		json.NewEncoder(S.Response).Encode("Server Starting Error")
	}
}

func (S *Server) handler() {
	http.HandleFunc("/", S.hello_world)

}
