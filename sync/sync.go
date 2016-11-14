package sync

import (
	"fmt"
	"net/http"
)

// 同步
type SyncHandler struct {
}

func (b *SyncHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sync")

	w.Write([]byte("hello, i am sync."))
}
