package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)
type movie struct{
ID string `json:"id"`
Isbn string `json:"isbn"`
title string `json:"title"`
director *director `json:"director"`

}
type director{
FirstName string `json:"firstname` 
lastName string `json:"lastname` 
}
var movies []Movie

func main(){
	r:=mux.NewRouter()
	movies=append(movies, Movie{ID:1,Isbn:15,title:"satyajig",Director:&Director{FirstName:"satya",lastName:"belu"}})
    r.HandleFunc("/movies",getMovies).Method("GET")
	r.HandleFunc("/movies/{id}",getMovie).Method("GET")
	r.HandleFunc("/movies",createMovie).Method("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Method("PUT")
	r.HandleFunc("/movies/{id}",deletMovie).Method("DELETE")
	fmt.printf("server started at 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}
func getMovies(w http.ResponseWriter, r *http.Request){
	
}