package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nickolaygoloburdin/httpserver/startprogram"
)

type request struct {
	Cmd string
}
type response struct {
	Status string
}

var stat string = "disabled"
var channel chan int

//var cmd *exec.Cmd

func HandlerOn(w http.ResponseWriter, r *http.Request) {
	//enableCors(&w)
	//go RunCommand("/home/copa5/start_copa.sh")
	//profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	fmt.Println("Get request on recieved")
	go startprogram.StartProgram("/home/jetson/copa5/start_copa.sh", channel, &stat)
	res := &response{stat}
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func HandlerOff(w http.ResponseWriter, r *http.Request) {
	//enableCors(&w)
	go startprogram.EndProgram("/home/jetson/copa5/stop_copa.sh", channel, &stat)
	res := &response{stat}
	//profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	fmt.Println("Get request off recieved")
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func HandlerConnect(w http.ResponseWriter, r *http.Request) {
	//enableCors(&w)
	res := &response{stat}
	fmt.Println("Get request connect recieved")
	js, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
}
func main() {
	channel = make(chan int, 1)
	//channel <- 1
	http.HandleFunc("/api/on", HandlerOn)
	http.HandleFunc("/api/off", HandlerOff)
	http.HandleFunc("/api/connect", HandlerConnect) // each request calls handler
	log.Fatal(http.ListenAndServe(":6600", nil))

}
