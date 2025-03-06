package main

import (
	_ "errors"
	"fmt"
	"net/http"
	"time"
)

const serverPort = ":8888" // Hardcoded port

// Handler for HTTP Requests
func handler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to HTML
	w.Header().Set("Content-Type", "text/html")

	t := time.Now()
	today := t.Format("2006-01-02") // YYYY-MM-DD format
	fmt.Fprintf(w, "Hello, it is "+fmt.Sprint(t.Hour())+":"+fmt.Sprint(getMinute(t.Minute(), t.Second()))+"<br />")
	fmt.Fprintf(w, "Today is: <b>%s</b>", today)
}

func getMinute(minute int, second int) int {
	return minute + second/30
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(serverPort, nil)
}
