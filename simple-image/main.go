package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

var (
	port = flag.Int("port", 13579, "server port")
)

func main() {
	flag.Parse()
	// 初始化 handler
	handler := NewServerHandler()

	// 初始化 server mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.BaseInfo)
	mux.HandleFunc("/cnt", handler.GetCnt)

	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), mux)
	if err != nil {
		log.Fatal(err)
	}
}

type serverHandler struct {
	cnt      int64
	baseInfo string
}

func NewServerHandler() *serverHandler {
	return &serverHandler{
		baseInfo: fmt.Sprintf(`
		<div>
		<a href="http://127.0.0.1:%d">base info</a>
		</div>
		<div>
		<a href="http://127.0.0.1:%d/cnt">count info</a>
		</div>
		`, *port, *port),
	}
}

func (s *serverHandler) addCnt() {
	log.Println("add cnt", s.cnt)
	atomic.AddInt64(&s.cnt, 1)
}

func (s *serverHandler) BaseInfo(w http.ResponseWriter, r *http.Request) {
	defer s.addCnt()
	w.Write([]byte(s.baseInfo))
}

func (s *serverHandler) GetCnt(w http.ResponseWriter, r *http.Request) {
	defer s.addCnt()
	w.Write([]byte(fmt.Sprintf(`
	<div>
	cur cnt:%d
	</div>
	`, s.cnt)))
}
