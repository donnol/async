package nonblock

import (
	"fmt"
	"net/http"
)

// 非阻塞
type NonblockHandler struct {
}

func (b *NonblockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println("nonblock")

	w.Write([]byte("hello, i am nonblock."))
}
