package consts

import "fmt"

const (
	reset = "\033[0m"
	red   = "\033[31m"
	green = "\033[32m"
)

func Green(v string) string {
	return fmt.Sprintf("%s%v%s", green, v, reset)
}

func Red(v string) string {
	return fmt.Sprintf("%s%v%s", red, v, reset)
}
