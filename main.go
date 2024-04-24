package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	PORT := os.Getenv("PORT")
	println("$PORT=", PORT)

	m := http.NewServeMux()

	addr := fmt.Sprintf(":%v", PORT)

	m.HandleFunc("/", handlePage)

	srv := http.Server{
		Handler:      m,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("server started on port", addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	const page = `
	<html>
<head></head>
<body>
	<p> Hi Docker, I pushed a new version! </p>
</body>
</html>
	`

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(page))
}
