package main

import (
	"fmt"
	"mime/multipart"
	"os/exec"
	"bytes"

	"github.com/valyala/fasthttp"
)

func Converter(ctx *fasthttp.RequestCtx, file multipart.File) {
	fmt.Println("Converter")

	cmd := exec.Command(
		"ffmpeg",
		"-i", "pipe:",
		"-c:a", "libvorbis",
		"-b:a", "32k",
		"-bufsize", "32k",
		"-ar", "22050",
		"-fflags", "+bitexact",
		"-flags:v", "+bitexact",
		"-flags:a", "+bitexact",
		"-analyzeduration", "0",
		"-loglevel", "0",
		"-f", "ogg",
		"pipe:")

	cmd.Stdin = file

	var out bytes.Buffer

	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		panic(err)
	}

	ctx.SetContentType("audio/ogg")
	ctx.SetBody(out.Bytes())
}