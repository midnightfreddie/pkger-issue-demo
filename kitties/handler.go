package kitties

import (
	"fmt"
	"math/rand"
	"net/http"
)

var randomKitties = []string{
	"random kitty pics",
	"random pics kitty",
	"kitty random pics",
	"kitty pics random",
	"pics random kitty",
	"pics kitty random",
}

// KittiesHandler returns random kitty pics
func KittiesHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, randomKitties[rand.Intn(len(randomKitties))])
	})
}
