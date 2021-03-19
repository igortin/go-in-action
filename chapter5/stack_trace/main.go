package main

import (
	"log/syslog"
	"runtime"
)

/*
Write programm stack trace to syslog
Stack trace - sequence calls of functions
*/
func bar() {

	buf := make([]byte, 1024)

	// Write stack to buf
	runtime.Stack(buf, false)

	// syslog Wrirter
	logger, err := syslog.New(syslog.LOG_LOCAL3, "stack_trace")
	if err != nil {
		panic(err)
	}

	// Write to system.log
	logger.Notice(string(buf))
}

func far() {
	bar()
}

func main() {
	far()
}

/*
Output:
Mar 17 10:46:38 Admins-MBP-2 stack_trace[20345]: goroutine 1 [running]:
        main.bar()
                /Users/igortin/go/src/go-in-action/chapter5/stack_trace/main.go:20 +0x6f
        main.far(...)
                /Users/igortin/go/src/go-in-action/chapter5/stack_trace/main.go:13
        main.main()
                /Users/igortin/go/src/go-in-action/chapter5/stack_trace/main.go:9 +0x25
*/
