package nonblock

import (
	"fmt"
	"net/http"
	"time"
)

// 非阻塞
// 调用者隔段时间查询结果
type NonblockHandler struct {
}

var c = make(chan int)

func (h NonblockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println("nonblock")

	// 在不能立刻得到结果之前，该调用不会阻塞当前线程。
	go h.sleep()

	// 在等待结果返回的同时，执行其它任务
	for {
		select {
		case <-c:
			// after sleep.
			w.Write([]byte("hello, i am nonblock."))
			return
		// do something
		default:
			fmt.Println("i am doing something.")
		}
	}
}

func (h NonblockHandler) sleep() {
	time.Sleep(5 * time.Second)

	fmt.Println("I slept for 5 seconds")
	c <- 1
}
