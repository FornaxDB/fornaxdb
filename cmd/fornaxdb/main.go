package main

import (
	"github.com/FornaxDB/fornaxdb/logger"
	"github.com/FornaxDB/fornaxdb/errors"
)


func main() {
	l := logger.New()
	l.Trace("Hello, World!", map[string]interface{}{"foo": "bar"})
	err := x()
	if err != nil {
		l.Fatal(err.Error(), map[string]interface{}{})
	}
}

func x() error {
	return errors.SchemaAlreadyExists.New("bla")
}
