package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)
//func main() {
//	http.HandleFunc("/", handler) // each request calls handler
//	log.Fatal(http.ListenAndServe("localhost:8000", nil))
//}
//
//// handler echoes the Path component of the request URL r.
//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
//}

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
//func handler(w http.ResponseWriter, r *http.Request) {
//	mu.Lock()
//	count++
//	mu.Unlock()
//	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
//}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	} }

//A handler pattern that ends with a slash matches any URL that has the pattern as a prefix.
//Behind the scenes, the server runs the handler for each incoming request in a separate goroutine
//so that it can serve multiple requests simultaneously.
//However, if two concurrent requests try to update count at the same time,
//it might not be incremented consistently;
//the program would have a serious bug called a race condition (§9.1).
//To avoid this problem, we must ensure that at most one goroutine accesses the variable at a time,
//which is the purpose of the mu.Lock() and mu.Unlock() calls that bracket each access of count