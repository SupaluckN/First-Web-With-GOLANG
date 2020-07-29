

package main

import "net/http"

//handlerIndex 
func handlerIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to This Site.</h1>"))
	w.Write([]byte("<p>This website created by GOLANG</p>"))
}	

func main() {
	// Registering a Request Handler
	http.HandleFunc("/", handlerIndex)
	http.ListenAndServe(":8888", nil)
}
