package godash

import "fmt"

func Expect(condition bool, message ...any) {
	if !condition {
		msg := "Expectation not met"
		if len(message) == 1 {
			msg = message[0].(string)
		}
		if len(message) > 1 {
			msg = fmt.Sprintf(message[0].(string), message[1:]...)
		}
		panic(msg)
	}
}
