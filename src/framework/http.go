package framework

import (
	"fmt"
	"log"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

type HTTPSystem struct {
	Port   string
	Router *routing.Router
}

func (h *HTTPSystem) init() {
	h.Router = routing.New()
	h.routeSysRoute()
}

func (h *HTTPSystem) listen() {
	fmt.Printf("\033[1;34m%s\033[0m", "[HTTP] ")
	log.Println("Listening on port " + h.Port)

	h.Router.Use(func(req *routing.Context) error {
		req.Response.Header.Set("Access-Control-Allow-Origin", "*")
		req.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		req.Response.Header.Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,POST,DELETE,OPTIONS")
		req.Response.Header.Set("Access-Control-Allow-Headers", "*")
		return nil
	})

	if err := fasthttp.ListenAndServe(":"+h.Port, h.Router.HandleRequest); err != nil {
		// Panic Error
		log.Fatalln("Error in ListenAndServe: " + err.Error())
	}
}

func (h *HTTPSystem) routeSysRoute() {
	h.Router.Get("/_/sys", func(c *routing.Context) error {
		fmt.Fprintf(c, "HTTP System functional!")
		return nil
	})

}
