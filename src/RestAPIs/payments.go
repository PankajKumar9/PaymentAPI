package RestAPIs

import (
	"github.com/PankajKumar9/PaymentAPI/src/conf"
	routing "github.com/qiangxue/fasthttp-routing"
)

func InitApis(r *routing.Router) {

	//Push order by API
	r.Post("/api/movie", func(c *routing.Context) error {
		return conf.Firstlogic(string(c.PostBody()), c)
		//return somefolder.somefunc(string(c.PostBody()), c)
	})
	r.Post("/api/signup", func(c *routing.Context) error {
		return conf.Signup(string(c.PostBody()), c)
		//return somefolder.somefunc(string(c.PostBody()), c)
	})
}
