package RestAPIs

import (
	"errors"

	routing "github.com/qiangxue/fasthttp-routing"
)

func InitApis(r routing.Router) {

	//Push order by API
	r.Post("path string", func(c *routing.Context) error {
		return errors.New("asd")
		//return somefolder.somefunc(string(c.PostBody()), c)
	})
}
