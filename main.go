package godash

import (
	"bytes"
	"fmt"
	"os"
	"runtime/debug"
)

// var buffer = make([]byte, 1024)

func Expect(condition bool, message ...any) {
	if !condition {
		msg := "Expectation not met"
		if len(message) == 1 {
			msg = message[0].(string)
		}
		if len(message) > 1 {
			msg = fmt.Sprintf(message[0].(string), message[1:]...)
		}
		fmt.Println("[godash] Expect: ", msg)

		var rune_index int
		stack_trace := debug.Stack()

		// TODO: runtime.callers ?
		for range 5 {
			rune_index = bytes.IndexByte(stack_trace, '\n')
			stack_trace = stack_trace[rune_index+1:]
		}

		fmt.Printf("[godash] %s\n", stack_trace)
		os.Exit(1)
	}
}
