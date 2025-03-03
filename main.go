package main

import (
	_ "errors" // Unused import, wird von SonarCloud als Problem erkannt
	"fmt"      // Unused import (wenn wir es nicht im Code nutzen)
	"net/http"
	"time"
)

const serverPort = ":8888" // Hardcoded port - könnte als Problem erkannt werden

func handler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to HTML
	w.Header().Set("Content-Type", "text/html")

	t := time.Now()
	today := t.Format("2006-01-02") // YYYY-MM-DD format
	fmt.Fprintf(w, "Hello, it is "+fmt.Sprint(t.Hour())+":"+fmt.Sprint(getMinute(t.Minute(), t.Second()))+"<br />")
	fmt.Fprintf(w, "Today is: <b>%s</b>", today)
}

func getMinute(minute int, second int) int {
	return minute + second/30 // Magic number 30 könnte als Issue erkannt werden
}

// http://localhost:8888
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(serverPort, nil) // Ignored error handling
}
