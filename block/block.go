package block

import (
	"fmt"
	"net/http"
	"time"
)

// 阻塞
type BlockHandler struct {
}

func (b BlockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("block")

	// 调用结果返回之前，当前线程会被挂起。调用线程只有在得到结果之后才会返回。
	b.sleep()

	w.Write([]byte("hello, i am block."))
}

func (b BlockHandler) sleep() {
	time.Sleep(5 * time.Second)

	fmt.Println("I slept for 5 seconds")
}
