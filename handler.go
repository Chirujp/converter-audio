package main

import (
	"crypto/sha256"
	"fmt"
	"io"

	"github.com/valyala/fasthttp"
)

func Handler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		HomeHandler(ctx)
	case "/convert":
		if ctx.IsPost() {
			ConvertHandler(ctx)
		} else {
			ctx.Error("", fasthttp.StatusMethodNotAllowed)
		}
	default:
		ctx.Error("Not found", fasthttp.StatusNotFound)
	}
}

func HomeHandler(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Churi internal service, please do not use !")
}

func ConvertHandler(ctx *fasthttp.RequestCtx) {
	fileHeader, err := ctx.FormFile("audio")

	if err != nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	file, err := fileHeader.Open()

	if err != nil {
		ctx.Error("", fasthttp.StatusBadRequest)
		return
	}

	defer file.Close()

	Converter(ctx, file)
}