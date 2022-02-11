package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)


func DefaultHandler(w http.ResponseWriter, req *http.Request) {
	sleepTime := req.URL.Query().Get("sleep")
	st, err := strconv.Atoi(sleepTime)
	if err != nil {
		log.Fatal("Cant convert sleep value")
	}
	time.Sleep(time.Duration(st) * time.Millisecond)

	w.Write([]byte(fmt.Sprintf("Slept %s miliseconds", sleepTime)))
}
