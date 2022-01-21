package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
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
	baseInfo string
	cntProxy CntRepo
}

func NewServerHandler() *serverHandler {
	return &serverHandler{
		cntProxy: NewCntRepo(),
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

func (s *serverHandler) BaseInfo(w http.ResponseWriter, r *http.Request) {
	defer func() {
		_ = s.cntProxy.Incr(1)
	}()
	w.Write([]byte(s.baseInfo))
}

func (s *serverHandler) GetCnt(w http.ResponseWriter, r *http.Request) {
	defer func() {
		_ = s.cntProxy.Incr(1)
	}()
	cnt,err := s.cntProxy.Get()
	if err != nil {
		log.Printf("get cnt fail,err:%+v",err)
	}
	w.Write([]byte(fmt.Sprintf(`
	<div>
	cur cnt:%d
	</div>
	`, cnt)))
}
