package main

import "net/http"

func main() {
	err := http.ListenAndServe(":9000", http.FileServer(http.Dir("wasm/assets")))
	if err != nil {
		panic(err)
	}
}
