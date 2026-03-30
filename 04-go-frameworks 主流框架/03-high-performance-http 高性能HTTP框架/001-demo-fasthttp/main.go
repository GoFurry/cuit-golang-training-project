package main

import (
	"encoding/json"
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	handler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/health":
			writeJSON(ctx, fasthttp.StatusOK, map[string]any{
				"framework": "fasthttp",
				"status":    "ok",
			})
		case "/hello":
			name := string(ctx.QueryArgs().Peek("name"))
			if name == "" {
				name = "Go learner"
			}
			writeJSON(ctx, fasthttp.StatusOK, map[string]string{
				"message": "hello, " + name,
			})
		default:
			writeJSON(ctx, fasthttp.StatusNotFound, map[string]string{
				"error": "route not found",
			})
		}
	}

	log.Println("fasthttp server listening on :8085")
	log.Fatal(fasthttp.ListenAndServe(":8085", handler))
}

func writeJSON(ctx *fasthttp.RequestCtx, status int, payload any) {
	body, err := json.Marshal(payload)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(status)
	ctx.SetContentType("application/json; charset=utf-8")
	ctx.SetBody(body)
}
