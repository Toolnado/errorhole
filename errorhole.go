package errorhole

import (
	"fmt"
	"log"
	"runtime"
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
		pc, file, line, _ := runtime.Caller(1)
		f := runtime.FuncForPC(pc).Name()
		panic(fmt.Errorf("%s:%s:%d: error: %s", file, f, line, err))
	}
}
