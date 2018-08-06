package http

import (
	"net/http"
	"encoding/json"
	"github.com/ChunmengYang/go-demo/g"
	"time"
	//_ "net/http/pprof" //ip:port//debug/pprof/
	//_ "github.com/mkevac/debugcharts"  //ip:port/debug/charts/
)

func init() {
	configRoutes()
}

func configRoutes() {
	http.HandleFunc("/index", func(w http.ResponseWriter, req *http.Request) {
		if req.ContentLength == 0 {
			http.Error(w, "body is blank", http.StatusBadRequest)
			return
		}

		decoder := json.NewDecoder(req.Body)
		var body string
		err := decoder.Decode(&body)
		if err != nil {
			http.Error(w, "connot decode body", http.StatusBadRequest)
			return
		}

		// TODO

		w.Write([]byte("success"))
	})
}

func Start() {
	if !g.Config().Http.Enabled {
		return
	}

	addr := g.Config().Http.Listen
	if addr == "" {
		return
	}

	server := &http.Server{
		Addr:           addr,
		ReadTimeout:	time.Second * 10,
		WriteTimeout: 	time.Second * 30,
	}

	g.Logger().Println("listening", addr)
	g.Logger().Println(server.ListenAndServe())
}