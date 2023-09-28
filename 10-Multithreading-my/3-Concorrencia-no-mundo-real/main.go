package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

func main() {
	//m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//m.Lock()
		//number++
		// atomic implement lock and unlock internally
		atomic.AddUint64(&number, 1)
		//m.Unlock()

		time.Sleep(200 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante número %d\n", number)))
	})
	http.ListenAndServe(":3000", nil)
}
