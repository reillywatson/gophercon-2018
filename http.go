package foo

import (
	"fmt"
	"net/http"
)

func FormHandler(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		res.WriteHeader(500)
		return
	}
	name := req.Form.Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(res, "Posted Hello, %s!", name)
}
