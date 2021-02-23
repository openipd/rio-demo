package main

import (
	"fmt"
        "os"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Arbitrary sleep so that we can demonstrate autoscaler
	time.Sleep(100 * time.Millisecond)

        name, err := os.Hostname()
        if err != nil {
            panic(err)
        }

	fmt.Fprintln(w, "Hi xxx, I'm running in Kubernetes Pod:", name)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
