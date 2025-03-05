package util

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

func FuncCost(funcName string) func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("%s func cost = %v \n", funcName, tc)
	}
}

func FuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	return path.Base(funcName)
}
