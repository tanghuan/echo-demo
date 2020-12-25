package utils

import (
	"errors"
	"fmt"
	"strings"
)

// SayHello 输出 Hello XXX.
func SayHello(name string) (string, error) {

	if strings.TrimSpace(name) == "" {
		return "", errors.New("Name can't be empty")
	}

	return fmt.Sprintf("Hello %s", name), nil

}
