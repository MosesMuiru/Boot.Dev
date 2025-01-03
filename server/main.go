package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlePage)

	const addr = ":8080"

	server := http.Server{
		Handler: mux,
		Addr: addr,
	}

	fmt.Println("server started on: ", addr)

	err := server.ListenAndServe()
	log.Fatal(err)

}

func handlePage(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	const page = `<html>
<head></head>
<body>
	<p> Hello from Docker! I'm a Go server. </p>
</body>
</html>
`

	w.Write([]byte(page))
}
