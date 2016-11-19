package sync

import (
	"fmt"
	"net/http"
	"time"
)

// 同步
type SyncHandler struct {
}

var c = make(chan int)

func (s SyncHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sync")

	go s.sleep()

	// 调用者主动等调用结果
	<-c

	w.Write([]byte("hello, i am sync."))
}

func (s SyncHandler) sleep() {
	time.Sleep(5 * time.Second)

	fmt.Println("I slept for 5 seconds")
	c <- 1
}
