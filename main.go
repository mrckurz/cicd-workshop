package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to HTML
	w.Header().Set("Content-Type", "text/html")

	t := time.Now()
	today := t.Format("2006-01-02") // YYYY-MM-DD format
	fmt.Fprintf(w, "Hello, it is %d:%d<br />", t.Hour(), getMinute(t.Minute(), t.Second()))
	fmt.Fprintf(w, "Today is: <b>%s</b>", today)
}

func getMinute(minute int, second int) int {
	return minute + second/30
}

// http://localhost:8888
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
