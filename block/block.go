package block

import (
	"fmt"
	"net/http"
)

// 阻塞
type BlockHandler struct {
}

func (b *BlockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("block")

	c := make(chan int)

	<-c
	w.Write([]byte("hello, i am block."))
}
