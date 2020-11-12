package r7uter

import (
	"fmt"
	"net/http"
)

func (app *Router) RunCmd(w http.ResponseWriter, r *http.Request) {
	fn := "CreateTest"
	fmt.Println(fn)
}
