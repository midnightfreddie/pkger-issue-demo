package main

import (
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/markbates/pkger"
	"github.com/midnightfreddie/pkger-issue-demo/kitties"
)

func main() {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	fmt.Println("http://127.0.0.1:" + strconv.Itoa(listener.Addr().(*net.TCPAddr).Port))

	staticFiles := http.FileServer(pkger.Dir("/cmd/server/html"))

	http.Handle("/", staticFiles)
	http.Handle("/kitties", kitties.KittiesHandler())
	panic(http.Serve(listener, nil))
}
