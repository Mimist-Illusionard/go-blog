package log

import (
	"fmt"
)

func Error(message string, err error) {
	fmt.Errorf("ERROR: %s, %s", message, err)
}
