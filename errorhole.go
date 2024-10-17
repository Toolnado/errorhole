package errorhole

import (
	"errors"
	"log"
	"runtime"
	"strconv"
	"strings"
)

type Handler func(x error)

var DefaultHandler = func(x error) {
	log.Println(x)
}

// Expects an error. Called with defer
func Catch(handlers ...Handler) {
	err := recover()
	if err != nil {
		if len(handlers) > 0 {
			for _, f := range handlers {
				f(err.(error))
			}
		} else {
			DefaultHandler(err.(error))
		}
	}
}

// Checks if the error is nil.
func Nil(err error) {
	if err != nil {
		pc, file, line, ok := runtime.Caller(1)
		build := strings.Builder{}
		if ok {
			f := runtime.FuncForPC(pc).Name()
			build.WriteString(file)
			build.WriteString(":")
			build.WriteString(f)
			build.WriteString(":")
			n := strconv.Itoa(line)
			build.WriteString(n)
			build.WriteString(": ")
		}
		build.WriteString(err.Error())
		panic(errors.New(build.String()))
	}
}

// Checks if the error is one of
func Other(err error, miss ...error) {
	if err != nil {
		missIt := false
		for _, w := range miss {
			if errors.Is(err, w) {
				missIt = true
			}
		}

		if !missIt {
			Nil(err)
		}
	}
}
