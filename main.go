package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	sleepTime := r.URL.Query().Get("sleep")
	realSleep, _ := strconv.ParseInt(sleepTime, 10, 0)
	time.Sleep(time.Duration(realSleep) * time.Second)
	fmt.Fprintf(w, "Sleeping for %s!", sleepTime)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
