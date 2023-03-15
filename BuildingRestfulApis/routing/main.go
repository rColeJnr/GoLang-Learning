package main

// basichandler
// func main() {
// 	http.HandleFunc("/hello", MyServer)          // creates /hello route. handleFunc maps an URL to a func
// 	log.Fatal(http.ListenAndServe(":8000", nil)) // start the sever on given port and return error if somethings goes wrong
// }
func main() {
	FileServer()
}

// exeService

// func main() {
// 	router := httprouter.New()
// 	// Mapping to methods is possible with HttpRputer
// 	// GET params: URL Path and HandlerFunc
// 	router.GET("/api/vi/go-version", GoVersion)
// 	router.GET("/api/v1/show-file/:name", GetFileContent)
// 	log.Fatal(http.ListenAndServe(":1205", router))
// }

// custommux
// func main() {
// 	// Any struct that has severHttp func can be a multiplexer
// 	newMux := http.NewServeMux()

// 	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, rand.Float64())
// 	})

// 	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, rand.Intn(100))
// 	})
// 	// mux := &CustomServeMux{}
// 	http.ListenAndServe(":8000", newMux)
// }
