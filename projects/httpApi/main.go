package main

import (
	// "fmt"
	"net/http"
	"encoding/json"
)

type Hw struct {
	OsType		string
	HostName	string
}
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	h := Hw{
		OsType: "Windows",
		HostName: "wty-pc",
	}
	output, err := json.MarshalIndent(&h, "", "")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return
}
func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func main(){
	server := http.Server {
		Addr: ":8080",
	}
	http.HandleFunc("/api", handleRequest)
	server.ListenAndServe()
}