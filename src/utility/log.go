package utility

import (
	"fmt"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"

	Red     = "\033[1;31m%s\033[0m"
	Green   = "\033[1;32m%s\033[0m"
	Black   = "\033[1;30m%s\033[0m"
	Yellow  = "\033[1;33m%s\033[0m"
	Purple  = "\033[1;34m%s\033[0m"
	Magenta = "\033[1;35m%s\033[0m"
	Teal    = "\033[1;36m%s\033[0m"
	White   = "\033[1;37m%s\033[0m"
)

//sadffadsfdsfdsfdssdfaafsfdasfsaffda

func Info(x interface{}) string {
	s := fmt.Sprintf("%v", x)
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, s)

}

func Debug(x interface{}) string {
	s := fmt.Sprintf("%v", x)

	return fmt.Sprintf("\033[%d1;31m%s\033[0m", 34, s)

}
