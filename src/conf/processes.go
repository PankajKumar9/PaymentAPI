package conf

import (
	"errors"
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
)

func Firstlogic(s string, c *routing.Context) error {
	fmt.Println(s)
	x := string(c.Request.Header.Peek("Transaction"))
	fmt.Println(x)
	return errors.New("hello")
}
