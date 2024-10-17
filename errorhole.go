package errorhole

import (
	"errors"
	"log"
	"runtime"
	"strconv"
	"strings"
)

var DefaultHandler = func(x error) {
	log.Println(x)
}

func Catch(handlers ...(func(x error))) {
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
