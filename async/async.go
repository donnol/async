package async

import (
	"fmt"
	"net/http"
)

// 异步
type AsyncHandler struct {
}

func (b *AsyncHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println("async")

	w.Write([]byte("hello, i am async."))
}
