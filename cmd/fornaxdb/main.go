package main

import (
	"github.com/FornaxDB/fornaxdb/logger"
)


func main() {
	l := logger.New()
	l.Trace("Hello, World!", map[string]interface{}{"foo": "bar"})
}
