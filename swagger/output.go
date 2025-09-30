package swagger

import "fmt"

func Output(format string, a ...any) {
	fmt.Printf(format, a...)
}
