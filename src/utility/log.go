package utility

import (
	"fmt"
	"runtime"
	"strings"
	"time"
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

type loghelper = func(a string, b string)
type logfunction = func(a string)

var Log logfunction
var Logx loghelper

func Init() {
	Logx = func(message string, color string) {
		s := "%s " + color

		fmt.Print(s, message)
	}

	Log = func(message string) {
		day := fmt.Sprintf("%v", time.Now().Local().Day())
		month := intmonth(fmt.Sprintf("%v", time.Now().Local().Month()))
		year := fmt.Sprintf("%v", time.Now().Local().Year())[2:]
		hour := fmt.Sprintf("%v", time.Now().Local().Hour())
		min := fmt.Sprintf("%v", time.Now().Local().Minute())
		sec := fmt.Sprintf("%v", time.Now().Local().Second())
		x := day + "-" + month + "-" + year + " " + hour + ":" + min + ":" + sec

		Logx(x, White)

		_, filename, line, _ := runtime.Caller(1)
		files := strings.Split(filename, "/")

		// blacks := color.New(color.BgBlack)
		// boldblacks := blacks.Add(color.Bold)
		//boldblacks.Println(exPath)

		Logx(files[len(files)-1]+"-> "+fmt.Sprintf("%v", line)+" ", Black)

		// Blues := color.New(color.FgHiBlue)
		// boldBlue := Blues.Add(color.Bold)
		// red := color.New(color.FgHiRed)
		// darkRed := red.Add(color.Bold)
		Logx(fmt.Sprintf("%v ", x), Teal)
		Logx(message, Red)
		fmt.Println()

	}
}
func intmonth(month string) string {
	My_map := make(map[string]string)
	My_map[time.January.String()] = "01"
	My_map[time.February.String()] = "02"
	My_map[time.March.String()] = "03"
	My_map[time.April.String()] = "04"
	My_map[time.May.String()] = "05"
	My_map[time.June.String()] = "06"
	My_map[time.July.String()] = "07"
	My_map[time.August.String()] = "08"
	My_map[time.September.String()] = "09"
	My_map[time.October.String()] = "10"
	My_map[time.November.String()] = "11"
	My_map[time.December.String()] = "12"

	return My_map[month]
}

//sadffadsfdsfdsfdssdfaafsfdasfsaffda

func Info(s string) string{
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, s)

}
