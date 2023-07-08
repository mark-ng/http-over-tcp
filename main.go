package main

import (
	"github.com/mark-ng/http-over-tcp/markhttp"
	"log"
)

func main() {
	log.Fatal(markhttp.ListenAndServe(":8080", nil))
}
