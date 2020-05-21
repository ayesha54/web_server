package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}



func incrementCounter(w http.ResponseWriter, r *http.Request) {
	
	mutex.Lock()
	if counter == 100{
		counter += 0
	} else{
		counter +=1
	}
	log.SetOutput(os.Stdout)
	log.Print(strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {
	
	
	http.HandleFunc("/increment", incrementCounter)
	
	log.Fatal(http.ListenAndServe(":8081", nil))
	
}

