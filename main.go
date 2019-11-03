package main

import "fmt"
import "net/http"

const CONTENTS = `hello, world
`

type DispViewHandler struct{}

func (h *DispViewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, CONTENTS)
}

func main() {
	var handler DispViewHandler = DispViewHandler{}

	http.ListenAndServe("127.0.0.1:8080", &handler)
}
