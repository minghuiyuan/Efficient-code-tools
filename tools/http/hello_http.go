package main

import (
"fmt"
"net/http"

"time"
)

// go run hello_http.go
// goto web broswer to : 127.0.0.1:8088

func main() {

	//127.0.0.1:8088/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	//127.0.0.1:8088/time
	http.HandleFunc("/time/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
		w.Write([]byte(timeStr))
	})

	http.ListenAndServe(":8088", nil)   //hander is a callback function.
}