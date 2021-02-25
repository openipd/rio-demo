package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Arbitrary sleep so that we can demonstrate autoscaler
	time.Sleep(100 * time.Millisecond)

	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, "Hi friends, I'm running in Kubernetes Pod:", name)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
