

package main

import "net/http"

/* Show Welcome Message on WebPage.
   -- handlerIndex --
*/

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to This Site.</h1>"))
	w.Write([]byte("<p>This website created by GOLANG</p>"))
}	

func main() {
	http.HandleFunc("/", handlerIndex)
	http.ListenAndServe(":8888", nil)
}