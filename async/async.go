package async

import (
	"fmt"
	"net/http"
	"time"
)

// 异步
type AsyncHandler struct {
}

var num int = 1

func (a AsyncHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println("async")
	fmt.Println("num begin value:", num)

	go a.sleep()

	w.Write([]byte("hello, i am async."))
}

// 5s后打印
func (a AsyncHandler) sleep() {
	time.Sleep(5 * time.Second)

	fmt.Println("I slept for 5 seconds")

	// *被调用者*通过状态、通知来通知调用者，或通过回调函数处理这个调用
	num++
	fmt.Println("num end value:", num)
}
