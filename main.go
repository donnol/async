package main

import (
	"fmt"
	"net/http"

	"jdscript.com/async/block"
)

func main() {
	port := "9009"
	fmt.Println("start http server...")
	fmt.Println("listen:" + port)
	http.ListenAndServe(":"+port, &block.BlockHandler{})
}
