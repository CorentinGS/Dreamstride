package utils

import (
	"log"
	"runtime"
)

func Log(caller string, id string, messages ...string) {
	pc, _, _, _ := runtime.Caller(1)
	callingFunc := runtime.FuncForPC(pc).Name()
	log.Println(callingFunc, "called by", caller, "with ID", id, messages)
}
