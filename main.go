package main

import (
	"fmt"
	"log"
	"net/http"

	"jdscript.com/async/async"
	"jdscript.com/async/block"
	"jdscript.com/async/nonblock"
	"jdscript.com/async/sync"
)

func main() {
	port := ":9009"

	fmt.Println("start http server...")
	fmt.Println("listen" + port)

	http.Handle("/block", block.BlockHandler{})
	http.Handle("/nonblock", nonblock.NonblockHandler{})
	http.Handle("/sync", sync.SyncHandler{})
	http.Handle("/async", async.AsyncHandler{})

	log.Fatal(http.ListenAndServe(port, nil))
}
