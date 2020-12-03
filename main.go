package main

import (
	"flag"
	"log"

	"github.com/valyala/fasthttp"
)

var (
	add = flag.String("add", "localhost", "proxy address")
)

type myLog struct {
	add string
}

func (l myLog) Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l myLog) myhandler(ctx *fasthttp.RequestCtx) {
	backend := fasthttp.HostClient{
		Addr: l.add,
	}
	backend.Do(&ctx.Request, &ctx.Response)
}

func main() {
	flag.Parse()
	log := myLog{
		add: *add,
	}
	s := fasthttp.Server{
		Handler:      log.myhandler,
		LogAllErrors: true,
		Logger:       log,
	}
	s.ListenAndServe(":8000")
}
